package service

import (
	"context"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"math/rand"
	"net/http"
	"sync"

	"things-service/entity"
	"things-service/model"
	"time"
)

var runningSimDeviceMap = sync.Map{}

func init() {
	common.RegisterDeleteBeforeHook("Sim_device", deleteSimDeviceHook)
	common.RegisterBatchDeleteBeforeHook("Sim_device", batchDeleteSimDeviceHook)
	/*
		go func() {
			time.Sleep(time.Second * 5) //启动阶段，调不到服务。
			restartSimDevices()
		}()
	*/
}

func restartSimDevices() {
	simDeviceInfos, err := common.DbQuery[model.Sim_device](context.Background(), common.GetDaprClient(), model.Sim_deviceTableInfo.Name, "")
	if err != nil {
		common.Logger.Error("restartSimDevices: query sim_device error", err)
		return
	}
	for _, simDeviceInfo := range simDeviceInfos {
		err = startOneSimDeviceByDbSimDevice(context.Background(), "system", &simDeviceInfo)
		if err != nil {
			err = errors.Wrap(err, "restartSimDevices: start sim_device error")
			common.Logger.Error(err)
		}
	}
}

func deleteSimDeviceHook(r *http.Request, in any) (out any, err error) {
	id := cast.ToString(in)
	stopChan, exists := runningSimDeviceMap.Load(id)
	if exists {
		stopChan.(chan bool) <- true
	}
	runningSimDeviceMap.Delete(id)
	return
}
func batchDeleteSimDeviceHook(r *http.Request, in any) (out any, err error) {
	var ids = in.([]string)
	for _, id := range ids {
		stopChan, exists := runningSimDeviceMap.Load(id)
		if exists {
			stopChan.(chan bool) <- true
		}
		runningSimDeviceMap.Delete(id)
	}
	return
}

func ChangeSimDeviceEnabled(ctx context.Context, sub, deviceId string, enabled int) (err error) {
	simDeviceInfo, err := common.DbGetOne[model.Sim_device](ctx, common.GetDaprClient(), model.Sim_deviceTableInfo.Name, model.Sim_device_FIELD_NAME_id+"="+deviceId)
	if err != nil {
		err = errors.Wrap(err, "ChangeSimDeviceEnabled: query sim_device error")
		return err
	}
	if simDeviceInfo == nil {
		simDeviceInfo, err = createDeviceSimToDb(ctx, sub, deviceId)
		if err != nil {
			return errors.Wrap(err, "start one device sim error")
		}
	}
	if enabled == 0 {
		stopChan, exists := runningSimDeviceMap.Load(deviceId)
		if exists {
			stopChan.(chan bool) <- true
		}
		runningSimDeviceMap.Delete(deviceId)
		common.DbDelete(ctx, common.GetDaprClient(), model.Sim_deviceTableInfo.Name, model.Sim_device_FIELD_NAME_id, deviceId)
	} else {
		err = startOneSimDeviceByDbSimDevice(ctx, sub, simDeviceInfo)
		if err != nil {
			err = errors.Wrap(err, "ChangeSimDeviceEnabled: start sim_device error")
			return err
		}
	}
	return
	//simDeviceInfo.Enabled = int32(enabled)

	//return common.DbUpsert[model.Sim_device](ctx, common.GetDaprClient(), *simDeviceInfo, model.Sim_deviceTableInfo.Name, model.Sim_device_FIELD_NAME_id)
}

func startOneSimDeviceByDbSimDevice(ctx context.Context, sub string, simDeviceInfo *model.Sim_device) (err error) {
	ruleJson := simDeviceInfo.RuleData
	if ruleJson == "" {
		err = errors.New("ChangeSimDeviceEnabled: rule data is empty")
		return err
	}
	var rule entity.SimRule
	err = json.Unmarshal([]byte(ruleJson), &rule)
	if err != nil {
		err = errors.Wrap(err, "ChangeSimDeviceEnabled: unmarshal rule error")
		return err
	}
	deviceInfo, err := common.DbGetOne[model.Device_info](ctx, common.GetDaprClient(), model.Device_infoTableInfo.Name, model.Device_info_FIELD_NAME_id+"="+simDeviceInfo.ID)
	if err != nil {
		err = errors.Wrap(err, "ChangeSimDeviceEnabled: query device_info error")
		return err
	}
	if deviceInfo == nil {
		err = errors.New("ChangeSimDeviceEnabled: device_info not found")
		return err
	}
	err = startOneDeviceSim(ctx, sub, rule, *deviceInfo)
	return err
}

func createDeviceSimToDb(ctx context.Context, sub string, id string) (simDevice *model.Sim_device, err error) {
	deviceInfo, err := common.DbGetOne[model.Device_info](ctx, common.GetDaprClient(), model.Device_infoTableInfo.Name, model.Device_info_FIELD_NAME_id+"="+id)
	if err != nil {
		err = errors.Wrap(err, "createDeviceSimToDb: query device_info error")
		return
	}
	if deviceInfo == nil {
		err = errors.New("can't find deviceInfo by id " + id)
		return
	}
	rule := entity.SimRule{}
	ruleJson, err := json.Marshal(rule)
	if err != nil {
		common.Logger.Error("marshal rule error:" + err.Error())
		return
	}
	simDevice = &model.Sim_device{
		ID:             deviceInfo.ID,
		Name:           deviceInfo.Name,
		CreatedBy:      sub,
		UpdatedBy:      sub,
		CreatedTime:    common.LocalTime(time.Now()),
		UpdatedTime:    common.LocalTime(time.Now()),
		Type:           deviceInfo.Type,
		ParentID:       deviceInfo.ParentID,
		GroupID:        deviceInfo.GroupID,
		ProductID:      deviceInfo.ProductID,
		ProtocolConfig: deviceInfo.ProtocolConfig,
		Identifier:     deviceInfo.Identifier,
		Enabled:        1,
		RuleData:       string(ruleJson),
	}
	err = common.DbUpsert[model.Sim_device](ctx, common.GetDaprClient(), *simDevice, model.Sim_deviceTableInfo.Name, model.Sim_device_FIELD_NAME_id)
	if err != nil {
		common.Logger.Error("dbUpsert sim_device error:device id=" + deviceInfo.ID + " err=" + err.Error())
		return
	}
	return
}
func startOneDeviceSim(ctx context.Context, sub string, rule entity.SimRule, deviceInfo model.Device_info) (err error) {

	sleepSeconds := rule.SimIntervalSeconds
	if sleepSeconds == 0 {
		sleepSeconds = 10
	}
	_, exist := runningSimDeviceMap.Load(deviceInfo.ID)
	if exist {
		return
	}
	stop := make(chan bool)
	runningSimDeviceMap.Store(deviceInfo.ID, stop)
	go func() {

		pmodel, _, err := GetDeviceProductModel(context.Background(), deviceInfo.Identifier)

		if err != nil {
			common.Logger.Error("GetDeviceProductModel error:" + err.Error())
			runningSimDeviceMap.Delete(deviceInfo.ID)
			return
		}
		for {
			select {
			case <-stop:
				common.Logger.Info(deviceInfo.ID + " sim stop")
				return
			default:
				oneDeviceSimLoop(context.Background(), rule, pmodel, deviceInfo)
				time.Sleep(time.Duration(sleepSeconds) * time.Second)
			}
		}
	}()
	return
}

func oneDeviceSimLoop(ctx context.Context, rule entity.SimRule, productModel *entity.ProductModel, deviceInfo model.Device_info) {
	t := time.Now()
	deviceMirror := &entity.DeviceMirror{
		State:                 entity.DeviceMirrorState{Reported: make(map[string]interface{}, 0), Desired: make(map[string]interface{}, 0)},
		Metadata:              entity.DeviceMirrorMetadata{Reported: make(map[string]interface{}, 0), Desired: make(map[string]interface{}, 0)},
		Timestamp:             t.UnixMilli(),
		Recent_InvokeServices: make([]entity.DeviceMirrorService, 0),
		RecentEvents:          make([]entity.DeviceMirrorEvent, 0),
		TimestampStr:          t.Format("2006-01-02 15:04:05"),
	}
	datas := make([]entity.DeviceModelPropertyWithData, 0)
	for _, deviceModel := range productModel.DeviceModels {
		for _, property := range deviceModel.Properties {
			data := entity.DeviceModelPropertyWithData{
				Property: property,
			}
			specMap, ok := property.DataType.Specs.(map[string]interface{})

			switch property.DataType.Type {
			case "int", "float", "double":
				min := 0
				max := 100
				if ok {
					max = cast.ToInt(specMap["max"])
					if max == 0 {
						common.Logger.Error("max is 0,set max to 100")
						max = 100
					}
					min = cast.ToInt(specMap["min"])
				}
				data.Data = rand.Intn(max-min) + min
			case "enum":
				if ok {
					keys := make([]string, 0)
					for k, _ := range specMap {
						keys = append(keys, k)
					}
					data.Data = keys[rand.Intn(len(keys))]
				}

			case "bool":
				data.Data = rand.Intn(2)
			case "text":
				data.Data = "text"
			case "date":
				data.Data = time.Now().Format("2006-01-02 15:04:05")
			case "struct", "array":
				data.Data = property.DataType.Type
			}
			datas = append(datas, data)
		}

	}
	for _, data := range datas {
		deviceMirror.State.Reported[data.Property.Identifier] = data.Data
	}
	err := PersistDeviceMirror(ctx, deviceInfo.Identifier, deviceMirror)
	if err != nil {
		common.Logger.Error("PersistSimDeviceMirror error:" + err.Error())
		//return
	}
	err = sendSimDeviceMirrorToWeb(ctx, deviceInfo, deviceMirror)
	if err != nil {
		common.Logger.Error("sendSimDeviceMirrorToWeb error:" + err.Error())
		return
	}
	/*
		err = saveSimDeviceData(ctx, deviceInfo.Identifier, t, datas)
		if err != nil {
			common.Logger.Error("saveSimDeviceData error:" + err.Error())
		}

	*/

}
func sendSimDeviceMirrorToWeb(ctx context.Context, deviceInfo model.Device_info, deviceMirror *entity.DeviceMirror) (err error) {
	buf, err := json.Marshal(deviceMirror)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	msg := common.CommonMessage{
		common.COMMON_MESSAGE_KEY_MARK:       "device_mirror_sim",
		common.COMMON_MESSAGE_KEY_CONNECT_ID: deviceInfo.Identifier,
		common.COMMON_MESSAGE_KEY_MESSAGE:    string(buf),
	}
	err = common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, common.WEB_MESSAGE_TOPIC_NAME, msg)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	//for workflow

	mapMsg := make(map[string]any)
	buf1, _ := json.Marshal(deviceInfo)
	_ = json.Unmarshal(buf1, &mapMsg)
	delete(mapMsg, "json_data") //减少数据大小
	for k, v := range deviceMirror.State.Reported {
		mapMsg[k] = v
	}

	err = common.GetDaprClient().PublishEvent(ctx, common.PUBSUB_NAME, "device", mapMsg)
	return
}

func saveSimDeviceData(ctx context.Context, identifier string, ts time.Time, values []entity.DeviceModelPropertyWithData) (err error) {
	dbDatas := make([]model.Sim_device_data, 0)
	for _, data := range values {
		dbData := model.Sim_device_data{
			ID:                 common.GetMD5Hash(identifier + "_" + data.Property.Identifier + "_" + ts.Format("2006-01-02 15:04:05")),
			DeviceIdentifier:   identifier,
			PropertyIdentifier: data.Property.Identifier,
			Ts:                 common.LocalTime(ts),
		}
		switch data.Property.DataType.Type {
		case "int", "enum", "bool":
			dbData.IValue = cast.ToInt32(data.Data)
		case "float", "double":
			dbData.FValue = cast.ToFloat64(data.Data)
		case "text":
			dbData.SValue = cast.ToString(data.Data)
		case "date":
			dbData.TValue = common.LocalTime(cast.ToTime(data.Data))
		case "struct", "array":
			buf, _ := json.Marshal(data.Data)
			dbData.SValue = string(buf)
		}
		dbDatas = append(dbDatas, dbData)
	}
	err = common.DbBatchInsert[model.Sim_device_data](ctx, common.GetDaprClient(), dbDatas, model.Sim_device_dataTableInfo.Name)
	return
}
