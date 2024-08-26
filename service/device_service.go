package service

import (
	"context"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"

	"things-service/entity"
	"things-service/eventpub"
	"things-service/model"
	"things-service/parsescript"
	"time"
)

var STATUS_ALERT = 2
var STATUS_ONLINE = 1
var STATUS_OFFLINE = 0

func init() {
	common.RegisterUpsertBeforeHook("Device", ProcessUpsertDevice)
	common.RegisterUpsertBeforeHook("User_device", ProcessUpsertUserDevice)
}
func ProcessUpsertUserDevice(r *http.Request, in any) (out any, err error) {
	sub, _ := common.ExtractUserSub(r)

	inObj := in.(model.User_device)
	if inObj.UserID == "" {
		inObj.UserID = sub
	}

	if inObj.ID == "" {
		inObj.ID = common.NanoId()
		inObj.CreatedTime = common.LocalTime(time.Now())
	}

	inObj.UpdatedTime = common.LocalTime(time.Now())

	out = inObj
	return
}

func ProcessUpsertDevice(r *http.Request, in any) (out any, err error) {
	sub, _ := common.ExtractUserSub(r)

	inObj := in.(model.Device)
	if inObj.ID == "" {
		inObj.ID = common.NanoId()
	}
	if inObj.CreatedBy == "" {
		inObj.CreatedBy = sub
	}
	zeroTime, _ := time.Parse("2006-01-02 15:04:05", "0001-01-01 00:00:00")
	if inObj.CreatedTime == common.LocalTime(zeroTime) {
		inObj.CreatedTime = common.LocalTime(time.Now())
	}
	inObj.UpdatedBy = sub
	inObj.UpdatedTime = common.LocalTime(time.Now())
	if inObj.Enabled == 0 { //如果disable 一个设备，那么将状态设为离线
		inObj.Status = 0
	}

	out = inObj
	common.PublishDbUpsertMessage(r.Context(), common.GetDaprClient(), model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id, "", false, inObj)
	return
}

func ProcessPropertySet(ctx context.Context, req entity.PropertySetReq) (err error) {
	productModel, status, err := GetDeviceProductModel(ctx, req.DeviceIdentifier)
	if err != nil {
		common.Logger.Error("MqttProcessDataDevice GetDeviceProductModel " + err.Error() + " identifier=" + req.DeviceIdentifier)
		err = errors.Wrap(err, "MqttProcessDataDevice GetDeviceProductModel")
		return
	}
	if productModel == nil {
		err = errors.New("productModel is nil")
		return
	}
	if status == STATUS_OFFLINE {
		err = errors.New("device " + req.DeviceIdentifier + " is offline")
		return
	}

	propertyName := "" //底层采集点位根据中文名称匹配
	for _, deviceModel := range productModel.DeviceModels {
		for _, property := range deviceModel.Properties {
			if property.Identifier == req.PropertyIdentifier {
				propertyName = property.Name
				goto FINDPROPNAME
			}

		}
	}
FINDPROPNAME:
	if propertyName == "" {
		err = errors.New(req.DeviceIdentifier + " property not found")
		return
	}

	sendPropMap := make(map[string]interface{}, 0)
	sendPropDesiredMap := make(map[string]interface{}, 0)
	parseScript := productModel.ParseScript
	if parseScript.Content != "" {
		if parseScript.Language == "golang" {
			processor, err := parsescript.GetGolangScriptProcessor(parseScript.Content)
			if err != nil {
				err = errors.Wrap(err, "GetGolangScriptProcessor")
				return err
			}
			deviceMirror, err := GetDeviceMirror(ctx, req.DeviceIdentifier)
			if err != nil {
				err = errors.Wrap(err, "GetDeviceMirror")
				return err
			}
			if deviceMirror != nil {
				sendPropMap = processor.ProcessTranslatePropertySet(deviceMirror.State.Reported, propertyName, req.Value)
				sendPropDesiredMap = processor.ProcessTranslatePropertySetDesired(deviceMirror.State.Reported, propertyName, req.Value)
			} else {
				sendPropMap = processor.ProcessTranslatePropertySet(make(map[string]any, 0), propertyName, req.Value)
				sendPropDesiredMap = processor.ProcessTranslatePropertySetDesired(make(map[string]any, 0), propertyName, req.Value)
			}

		} else {
			common.Logger.Debugf("%s process parsescript not golang", req.DeviceIdentifier)
			err = errors.New("parse script language not support")
			return
		}
	} else {
		sendPropMap[propertyName] = req.Value
		sendPropDesiredMap[propertyName] = req.Value
	}

	for k, v := range sendPropMap { //k 是中文
		propSetMsg := make(map[string]interface{}, 0)
		propSetMsg["device_identifier"] = req.DeviceIdentifier
		propSetMsg["property_name"] = k
		propSetMsg["property_value"] = v
		err = eventpub.PublishInternalMessage(ctx, common.PUBSUB_NAME, common.PROPERTY_SET_TOPIC, propSetMsg)
		if err != nil {
			common.Logger.Error("MqttProcessDataDevice PublishInternalMessage " + err.Error() + " identifier=" + req.DeviceIdentifier)
			err = errors.Wrap(err, "MqttProcessDataDevice PublishInternalMessage")
			return
		}

	}
	lck := GetDeviceMirrorLock(req.DeviceIdentifier)
	lck.Lock()
	defer lck.Unlock()
	for k, v := range sendPropDesiredMap {
		for _, deviceModel := range productModel.DeviceModels {
			for _, property := range deviceModel.Properties {
				if property.Name == k {
					err = MergeDeviceMirrorDesired(ctx, req.DeviceIdentifier, property.Identifier, v)
					if err != nil {
						err = errors.Wrap(err, "MergeDeviceMirrorDesired")
						return
					}
					goto ENDFOR
				}

			}
		}
	ENDFOR:
	}
	deviceMirror, _ := GetDeviceMirror(ctx, req.DeviceIdentifier)
	//立刻发一次到web
	err = SendDeviceMirrorToMessageChannel(ctx, req.DeviceIdentifier, nil, deviceMirror, "")
	if err != nil {
		common.Logger.Error("MqttProcessDataDevice SendDeviceMirrorToMessageChannel " + err.Error() + " identifier=" + req.DeviceIdentifier)
	}
	return
}

// 处理采集层收到的设备数据
func ProcessDeviceMsg(ctx context.Context, deviceMsg entity.DeviceInfoMsg) (err error) {
	deviceOffline := false
	if len(deviceMsg.Properties) == 0 { //提前判断，后面脚本处理可能添加新的属性。
		deviceOffline = true

	}
	productModel, _, err := GetDeviceProductModel(context.Background(), deviceMsg.Identifier)
	if err != nil {
		common.Logger.Error("MqttProcessDataDevice GetDeviceProductModel " + err.Error() + " identifier=" + deviceMsg.Identifier)
		return
	}
	t := time.UnixMilli(deviceMsg.Ts)
	lck := GetDeviceMirrorLock(deviceMsg.Identifier)
	lck.Lock()
	defer lck.Unlock()

	deviceMirror, err := GetDeviceMirror(ctx, deviceMsg.Identifier)
	if err != nil {
		common.Logger.Error("MqttProcessDataDevice GetDeviceProductModel get mirror error" + err.Error() + " identifier=" + deviceMsg.Identifier)
		return
	}
	if deviceMirror != nil {
		if !deviceOffline {
			deviceMirror.Timestamp = t.UnixMilli()
			deviceMirror.TimestampStr = t.Format("2006-01-02 15:04:05")
		}

	} else {
		deviceMirror = NewDeviceMirror()
	}
	if deviceMirror.Alerts == nil {
		deviceMirror.Alerts = make([]map[string]any, 0)
	}
	existEvent := deviceMirror.Alerts

	events := make([]map[string]any, 0)
	parseScript := productModel.ParseScript
	if parseScript.Content != "" {
		if parseScript.Language == "golang" {
			common.Logger.Debugf("%s process parsescript golang", deviceMsg.Identifier)
			processor, err := parsescript.GetGolangScriptProcessor(parseScript.Content)
			if err != nil {
				err = errors.Wrap(err, "GetGolangScriptProcessor")
				return err
			}
			transMap := processor.ProcessTranslatePropertyValue(deviceMsg.Properties)
			for k, v := range transMap {
				deviceMsg.Properties[k] = v
			}
			events = processor.CheckAlert(deviceMsg.Properties)

		} else {
			common.Logger.Debugf("%s process parsescript not golang", deviceMsg.Identifier)
			err = errors.New("parse script language not support")
			return
		}
	}

	datas := make([]entity.DeviceModelPropertyWithData, 0)
	for _, deviceModel := range productModel.DeviceModels {
		for _, property := range deviceModel.Properties {
			v, exist := deviceMsg.Properties[property.Name]
			if exist {
				deviceMirror.State.Reported[property.Identifier] = v
				desiredV, dexist := deviceMirror.State.Desired[property.Identifier]
				if dexist {
					if cast.ToString(v) == cast.ToString(desiredV) {
						delete(deviceMirror.State.Desired, property.Identifier)
					}
				}

				datas = append(datas, entity.DeviceModelPropertyWithData{
					Property: property,
					Data:     v,
				})

				SetDeviceGaugeValue(cast.ToString(productModel.Profile["name"]), property.Identifier, deviceMsg.Identifier, v)
			} else {
				common.Logger.Error(property.Name + " can't find in device message")
				continue
			}

		}
	}
	device, err := GetDeviceWithTagByIdentifier(ctx, deviceMsg.Identifier)
	if err != nil {
		err = errors.Wrap(err, "GetDeviceWithTagByIdentifier")
		return
	}
	if !deviceOffline {
		err = SendDeviceMirrorToMessageChannel(ctx, deviceMsg.Identifier, device, deviceMirror, "")
		if err != nil {
			common.Logger.Error("SendDeviceMirrorToMessageChannel " + err.Error())
		}
	}

	err = SaveDeviceData(ctx, deviceMsg.Identifier, t, datas)
	if err != nil {
		common.Logger.Error("SaveDeviceData " + err.Error())
	}
	for _, oldEvent := range existEvent { //发送清除告警
		find := false
		for _, event := range events {
			if event["alert_property"] == oldEvent["alert_property"] {
				find = true
				break
			}
		}
		if !find {
			alertProperty := cast.ToString(oldEvent["alert_property"])
			alertValue := cast.ToString(oldEvent["alert_value"])
			eventpub.ConstructAndSendEvent(ctx, common.EventTypeDevice, deviceMsg.Identifier+" "+alertProperty, deviceMsg.Identifier+" "+alertProperty+" 状态为"+alertValue, common.EventStatusClosed, common.EventLevelMajor, time.Now(), device.Id, device.Name, "")
		}
	}
	for _, event := range events { //发送活动告警
		alertProperty := cast.ToString(event["alert_property"])
		alertValue := cast.ToString(event["alert_value"])
		eventpub.ConstructAndSendEvent(ctx, common.EventTypeDevice, deviceMsg.Identifier+" "+alertProperty, deviceMsg.Identifier+" "+alertProperty+" 状态为"+alertValue, common.EventStatusActive, common.EventLevelMajor, time.Now(), device.Id, device.Name, "")
	}
	deviceMirror.Alerts = events
	if len(events) > 0 {
		err = SetDeviceStatus(ctx, common.GetMD5Hash(deviceMsg.Identifier), STATUS_ALERT)
	} else {
		if deviceOffline {
			err = SetDeviceStatus(ctx, common.GetMD5Hash(deviceMsg.Identifier), STATUS_OFFLINE)
			eventpub.ConstructAndSendEvent(ctx, common.EventTypeDevice, device.Identifier+" 离线", deviceMsg.Identifier+" 离线", common.EventStatusActive, common.EventLevelCritical, time.Now(), device.Id, device.Name, "")
		} else {
			if device.Status == STATUS_OFFLINE {
				eventpub.ConstructAndSendEvent(ctx, common.EventTypeDevice, device.Identifier+" 离线", deviceMsg.Identifier+" 离线", common.EventStatusClosed, common.EventLevelCritical, time.Now(), device.Id, device.Name, "")
			}
			err = SetDeviceStatus(ctx, common.GetMD5Hash(deviceMsg.Identifier), STATUS_ONLINE)

		}
	}

	if err != nil {
		common.Logger.Error("SetDeviceStatus " + err.Error())
	}
	if !deviceOffline {
		err = PersistDeviceMirror(ctx, device.Identifier, deviceMirror)
		if err != nil {
			common.Logger.Error("PersistDeviceMirror " + err.Error())
		}
	}

	return
}

func SaveDeviceData(ctx context.Context, identifier string, ts time.Time, values []entity.DeviceModelPropertyWithData) (err error) {
	dbDatas := make([]model.Device_data, 0)
	for _, data := range values {

		dbData := model.Device_data{
			ID:                 common.GetMD5Hash(identifier + "_" + data.Property.Name),
			DeviceIdentifier:   identifier,
			PropertyIdentifier: data.Property.Identifier,
			Ts:                 common.LocalTime(ts),
			Unit:               cast.ToString(cast.ToStringMap(data.Property.DataType.Specs)["unit"]),
		}

		switch data.Property.DataType.Type {
		case "int", "enum", "bool":
			dbData.IValue = cast.ToInt32(data.Data)
			dbData.Vtype = 2
		case "float", "double":
			dbData.FValue = cast.ToFloat64(data.Data)
			dbData.Vtype = 1
		case "text":
			dbData.SValue = cast.ToString(data.Data)
			dbData.Vtype = 4
		case "date":
			dbData.TValue = common.LocalTime(cast.ToTime(data.Data))
			dbData.Vtype = 3
		case "struct", "array":
			buf, _ := json.Marshal(data.Data)
			dbData.SValue = string(buf)
			dbData.Vtype = 4
		}
		dbDatas = append(dbDatas, dbData)
	}
	err = common.DbBatchUpsert[model.Device_data](ctx, common.GetDaprClient(), dbDatas, model.Device_dataTableInfo.Name, model.Device_data_FIELD_NAME_id)
	return
}
func GetDeviceWithTagByIdentifier(ctx context.Context, identifier string) (device *entity.DeviceInfo, err error) {
	qstr := "identifier=" + identifier
	device, err = common.DbGetOne[entity.DeviceInfo](ctx, common.GetDaprClient(), model.Device_with_tagTableInfo.Name, qstr)
	return
}
func GetDeviceWithTagById(ctx context.Context, identifier string) (device *entity.DeviceInfo, err error) {
	qstr := "id=" + identifier
	device, err = common.DbGetOne[entity.DeviceInfo](ctx, common.GetDaprClient(), model.Device_with_tagTableInfo.Name, qstr)
	return
}

func SetDeviceStatus(ctx context.Context, deviceId string, status int) (err error) {
	dm := make(map[string]any)
	dm[model.Device_FIELD_NAME_id] = deviceId
	dm[model.Device_FIELD_NAME_status] = status
	err = common.DbUpsert[map[string]any](ctx, common.GetDaprClient(), dm, model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id)
	return
}

func GetEnabledDeviceIdentifiersByProductName(ctx context.Context, productName string) (deviceIds []string, err error) {
	qstr := "_select=identifier&enabled=1&product_name=" + productName
	datas, err := common.DbQuery[model.Device_info](ctx, common.GetDaprClient(), model.Device_infoTableInfo.Name, qstr)
	if err != nil {
		err = errors.Wrap(err, "GetEnabledDeviceIdsByProductName")
		common.Logger.Error(err.Error())
		return
	}
	deviceIds = make([]string, 0)
	for _, data := range datas {
		deviceIds = append(deviceIds, data.Identifier)
	}
	return
}
