package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dapr-platform/common"
	"github.com/mozillazg/go-pinyin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"strings"

	"things-service/model"
	"time"
)

type DeviceExcelFileProcessor struct {
	Tag int64
}

func NewDeviceExcelFileProcessor() *DeviceExcelFileProcessor {
	return &DeviceExcelFileProcessor{
		Tag: time.Now().Unix(),
	}
}

func (p *DeviceExcelFileProcessor) ProcessModbusExcel(ctx context.Context, filePath string) (err error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		common.Logger.Error(err.Error())
		return err
	}
	for _, name := range f.GetSheetMap() {
		rows, err := f.GetRows(name)
		if err != nil {
			err = errors.Wrap(err, name+" get rows error")
			common.Logger.Error(err.Error())
			return err
		}
		switch name {
		case "access_protocol":
			err = p.processModbusAccessProtocol(ctx, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusGateway error")
				return err
			}
		}

	}
	devices := make([]model.Device, 0)
	for _, name := range f.GetSheetMap() {
		rows, err := f.GetRows(name)
		tbName := name
		if strings.Index(tbName, "#") > 0 {
			tbName = strings.Split(tbName, "#")[0]
		}
		if err != nil {
			err = errors.Wrap(err, name+" get rows error")
			common.Logger.Error(err.Error())
			return err
		}
		switch tbName {
		case "gateway":
			err = p.processModbusGateway(ctx, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusGateway error")
				return err
			}
		case "device":
			devices, err = p.processModbusDevice(ctx, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusDevice error")
				return err
			}

		}

	}

	for _, name := range f.GetSheetMap() {
		rows, err := f.GetRows(name)
		if err != nil {
			err = errors.Wrap(err, name+" get rows error")
			common.Logger.Error(err.Error())
			return err
		}
		tbName := name
		if strings.Index(tbName, "#") > 0 {
			tbName = strings.Split(tbName, "#")[0]
		}
		switch tbName {
		case "point_485":
			err = p.processModbusPoint485(ctx, devices, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusPoint485 error")
				return err
			}
		case "point_io":
			err = p.processModbusPointIO(ctx, devices, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusPoint485 error")
				return err
			}
		}

	}
	return
}

func (p *DeviceExcelFileProcessor) ProcessMqttExcel(ctx context.Context, filePath string) (err error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		common.Logger.Error(err.Error())
		return err
	}
	for _, name := range f.GetSheetMap() {
		rows, err := f.GetRows(name)
		if err != nil {
			err = errors.Wrap(err, name+" get rows error")
			common.Logger.Error(err.Error())
			return err
		}
		switch name {
		case "access_protocol":
			err = p.processModbusAccessProtocol(ctx, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusGateway error")
				return err
			}
		}

	}
	devices := make([]model.Device, 0)
	for _, name := range f.GetSheetMap() {
		rows, err := f.GetRows(name)
		if err != nil {
			err = errors.Wrap(err, name+" get rows error")
			common.Logger.Error(err.Error())
			return err
		}
		switch name {
		case "gateway":
			err = p.processModbusGateway(ctx, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusGateway error")
				return err
			}
		case "device":
			devices, err = p.processModbusDevice(ctx, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusDevice error")
				return err
			}

		}

	}

	for _, name := range f.GetSheetMap() {
		rows, err := f.GetRows(name)
		if err != nil {
			err = errors.Wrap(err, name+" get rows error")
			common.Logger.Error(err.Error())
			return err
		}
		switch name {
		case "point":
			err = p.processModbusPointIO(ctx, devices, rows)
			if err != nil {
				err = errors.Wrap(err, name+" processModbusPoint error")
				return err
			}

		}

	}
	return
}

func (p *DeviceExcelFileProcessor) processModbusAccessProtocol(ctx context.Context, rows [][]string) (err error) {
	common.Logger.Debug("processModbusAccessProtocol")
	var headers []string
	tagIndex := 2
	acs := make([]model.Access_protocol, 0)

	for i, row := range rows {
		if i == 0 { //标题行
			headers = row
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			common.Logger.Error("empty access protocol name")
			continue
		}

		acId := common.GetMD5Hash(row[0])
		accessProtocol := model.Access_protocol{
			ID:          acId,
			Identifier:  row[0],
			CreatedTime: common.LocalTime(time.Now()),
			UpdatedTime: common.LocalTime(time.Now()),
			Status:      0,
			Enabled:     1,
			Type:        row[1],
		}
		props := make(map[string]string, 0)
		for idx := tagIndex; idx < len(row); idx++ {
			props[headers[idx]] = cast.ToString(row[idx])
		}
		buf, _ := json.Marshal(props)
		accessProtocol.Properties = string(buf)
		acs = append(acs, accessProtocol)

	}
	if len(acs) > 0 {
		err = common.DbBatchUpsert[model.Access_protocol](ctx, common.GetDaprClient(), acs, model.Access_protocolTableInfo.Name, model.Access_protocol_FIELD_NAME_id)
		if err != nil {
			return errors.Wrap(err, "db upsert gateway error")
		}
	}

	return
}

func (p *DeviceExcelFileProcessor) processModbusGateway(ctx context.Context, rows [][]string) (err error) {
	common.Logger.Debug("processModbusGateway")
	var headers []string
	tagIndex := 1
	devices := make([]model.Device, 0)
	tags := make([]model.Tag, 0)

	for i, row := range rows {
		if i == 0 { //标题行
			headers = row
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			common.Logger.Error("empty gateway name")
			continue
		}

		gatewayId := common.GetMD5Hash(row[0])
		gateway, err1 := common.DbGetOne[model.Device](ctx, common.GetDaprClient(), model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id+"="+gatewayId)
		if err1 != nil {
			err = errors.Wrap(err1, gatewayId+" get gateway error")
			return
		}
		if gateway == nil {
			gateway = &model.Device{
				ID:          gatewayId,
				Identifier:  row[0],
				Name:        row[0],
				Type:        2,
				CreatedTime: common.LocalTime(time.Now()),
				UpdatedTime: common.LocalTime(time.Now()),
				ParentID:    "0",
			}
		} else {
			gateway.UpdatedTime = common.LocalTime(time.Now())
		}
		devices = append(devices, *gateway)
		for j := tagIndex; j < len(row); j++ {
			if len(headers) < j {
				continue
			}
			if headers[j] == "" || row[j] == "" {
				common.Logger.Warning(gatewayId + " has empty tag " + headers[j] + " " + row[j])
				continue
			}
			tag := model.Tag{
				ID:       common.GetMD5Hash(gatewayId + "_" + headers[j]),
				RelID:    gatewayId,
				RelType:  2,
				Key:      headers[j],
				Value:    row[j],
				Editable: 0,
			}
			tags = append(tags, tag)
		}

	}
	if len(devices) > 0 {
		err = common.DbBatchUpsert[model.Device](ctx, common.GetDaprClient(), devices, model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id)
		if err != nil {
			return errors.Wrap(err, "db upsert gateway error")
		}
		for _, d := range devices {
			deleteImportTags(ctx, d.ID)
		}
	}

	if len(tags) > 0 {
		err = common.DbBatchUpsert[model.Tag](ctx, common.GetDaprClient(), tags, model.TagTableInfo.Name, model.Tag_FIELD_NAME_id)
		if err != nil {
			return errors.Wrap(err, "db upsert tag error")
		}
	} else {
		common.Logger.Warning("tag is 0")
	}

	return
}

func deleteImportTags(ctx context.Context, id string) {
	err := common.DbDeleteByOps(ctx, common.GetDaprClient(), model.TagTableInfo.Name, []string{model.Tag_FIELD_NAME_rel_id, model.Tag_FIELD_NAME_editable}, []string{"==", "=="}, []any{id, 0})
	if err != nil {
		common.Logger.Error("deleteImportTags", err)
	}

}

func (p *DeviceExcelFileProcessor) processModbusPoint485(ctx context.Context, devices []model.Device, rows [][]string) (err error) {
	common.Logger.Debug("processModbusPoint485")
	var headers []string
	tagIndex := 2
	points := make([]model.Point, 0)
	tags := make([]model.Tag, 0)

	for i, row := range rows {
		if i == 0 { //标题行
			headers = row
			continue
		}
		if strings.TrimSpace(row[0]) == "" || strings.TrimSpace(row[1]) == "" {
			common.Logger.Error("row 0 or 1 empty")
			continue
		}
		if len(row) < 2 {
			common.Logger.Error("row length less than 2")
			continue
		}
		pointId := common.GetMD5Hash(row[1] + "_" + row[0])
		deviceId := common.GetMD5Hash(row[1])
		var device *model.Device
		for _, d := range devices {
			if d.ID == deviceId {
				device = &d
				break
			}
		}
		if device == nil {
			common.Logger.Error(row[0] + " device not found")
			continue
		}

		point := model.Point{
			ID:        pointId,
			GatewayID: device.ParentID,
			DeviceID:  deviceId,
			Name:      row[0],
		}
		points = append(points, point)
		for j := tagIndex; j < len(row); j++ {
			if len(headers) < j {
				continue
			}
			if headers[j] == "" || row[j] == "" {
				continue
			}
			tag := model.Tag{
				ID:       common.GetMD5Hash(pointId + "_" + headers[j]),
				RelID:    pointId,
				RelType:  3,
				Key:      headers[j],
				Value:    row[j],
				Editable: 0,
			}
			tags = append(tags, tag)
		}

	}
	if len(points) > 0 {
		err = common.DbBatchUpsert[model.Point](ctx, common.GetDaprClient(), points, model.PointTableInfo.Name, model.Point_FIELD_NAME_id)
		if err != nil {
			return errors.Wrap(err, "db upsert point modbus error")
		}
		for _, d := range points {
			deleteImportTags(ctx, d.ID)
		}
	}

	if len(tags) > 0 {

		err = common.DbBatchUpsert[model.Tag](ctx, common.GetDaprClient(), tags, model.TagTableInfo.Name, model.Tag_FIELD_NAME_id)
		if err != nil {
			return errors.Wrap(err, "db upsert tag error")
		}
	}

	return
}

func (p *DeviceExcelFileProcessor) processModbusPointIO(ctx context.Context, devices []model.Device, rows [][]string) (err error) {
	common.Logger.Debug("processModbusPointIO")
	var headers []string
	tagIndex := 2
	points := make([]model.Point, 0)
	tags := make([]model.Tag, 0)

	for i, row := range rows {
		if i == 0 { //标题行
			headers = row
			continue
		}
		if strings.TrimSpace(row[0]) == "" || strings.TrimSpace(row[1]) == "" {
			common.Logger.Error("row 0 or 1 empty")
			continue
		}
		if len(row) < tagIndex {
			common.Logger.Error("row length less than " + cast.ToString(tagIndex))
			continue
		}
		pointId := common.GetMD5Hash(row[1] + "_" + row[0])
		deviceId := common.GetMD5Hash(row[1])
		var device *model.Device
		for _, d := range devices {
			if d.ID == deviceId {
				device = &d
				break
			}
		}
		if device == nil {
			common.Logger.Error(row[0] + " device not found")
			continue
		}

		point := model.Point{
			ID:        pointId,
			GatewayID: device.ParentID,
			DeviceID:  deviceId,
			Name:      row[0],
		}
		points = append(points, point)
		for j := tagIndex; j < len(row); j++ {
			if len(headers) < j {
				continue
			}
			if headers[j] == "" || row[j] == "" {
				continue
			}
			tag := model.Tag{
				ID:       common.GetMD5Hash(pointId + "_" + headers[j]),
				RelID:    pointId,
				RelType:  3,
				Key:      headers[j],
				Value:    row[j],
				Editable: 0,
			}
			tags = append(tags, tag)
		}

	}
	if len(points) > 0 {
		err = common.DbBatchUpsert[model.Point](ctx, common.GetDaprClient(), points, model.PointTableInfo.Name, model.Point_FIELD_NAME_id)
		if err != nil {
			return errors.Wrap(err, "db upsert point modbus error")
		}
		for _, d := range points {
			deleteImportTags(ctx, d.ID)
		}
	}

	if len(tags) > 0 {

		err = common.DbBatchUpsert[model.Tag](ctx, common.GetDaprClient(), tags, model.TagTableInfo.Name, model.Tag_FIELD_NAME_id)
		if err != nil {
			return errors.Wrap(err, "db upsert tag error")
		}
	}

	return
}

func (p *DeviceExcelFileProcessor) processModbusDevice(ctx context.Context, rows [][]string) (devices []model.Device, err error) {
	var headers []string
	tagIndex := 1
	tags := make([]model.Tag, 0)
	if len(rows) > 1 {
		headers = rows[0]
	}
	productIdx := 0
	for i, h := range headers {
		if h == "产品" { //产品特殊处理
			productIdx = i
			break
		}
	}
	productName2IdMap := make(map[string]string)
	if productIdx > 0 {
		qstr := "_select=id,name"
		products, err := common.DbQuery[map[string]any](ctx, common.GetDaprClient(), model.ProductTableInfo.Name, qstr)
		if err == nil {
			for _, p := range products {
				productName2IdMap[p["name"].(string)] = p["id"].(string)
			}
		}
	}

	for i, row := range rows {
		if i == 0 { //标题行
			headers = row
			continue
		}
		if len(row) < 2 {
			common.Logger.Error("row length less than 2")
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			common.Logger.Error("row 0  is blank")
			continue
		}

		deviceId := common.GetMD5Hash(row[0])
		gatewayId := common.GetMD5Hash(row[1])
		device, err1 := common.DbGetOne[model.Device](ctx, common.GetDaprClient(), model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id+"="+deviceId)
		if err1 != nil {
			err = errors.Wrap(err1, deviceId+" get device error")
			return
		}
		productId := ""
		if productIdx > 0 && row[productIdx] != "" && productName2IdMap[row[productIdx]] != "" {
			productId = productName2IdMap[row[productIdx]]
		}
		if device == nil {

			device = &model.Device{
				ID:          deviceId,
				Identifier:  row[0],
				Name:        row[0],
				Type:        1,
				CreatedTime: common.LocalTime(time.Now()),
				UpdatedTime: common.LocalTime(time.Now()),
				ParentID:    gatewayId,
				ProductID:   productId,
			}
		} else {
			device.ParentID = gatewayId
			if productId != "" {
				device.ProductID = productId
			}
			device.UpdatedTime = common.LocalTime(time.Now())
		}
		err = common.DbDeleteByOps(ctx, common.GetDaprClient(), model.PointTableInfo.Name, []string{model.Point_FIELD_NAME_device_id}, []string{"=="}, []any{deviceId})
		if err != nil {
			common.Logger.Error("delete old point error ")
		}
		devices = append(devices, *device)
		for j := tagIndex; j < len(row); j++ {
			if len(headers) < j {
				continue
			}
			if headers[j] == "" || row[j] == "" {
				continue
			}
			tag := model.Tag{
				ID:       common.GetMD5Hash(deviceId + "_" + headers[j]),
				RelID:    deviceId,
				RelType:  1,
				Key:      headers[j],
				Value:    row[j],
				Editable: 0,
			}
			tags = append(tags, tag)
		}

	}
	if len(devices) > 0 {
		err = common.DbBatchUpsert[model.Device](ctx, common.GetDaprClient(), devices, model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id)
		if err != nil {
			err = errors.Wrap(err, "db upsert devices error")
			return
		}
		for _, d := range devices {
			deleteImportTags(ctx, d.ID)
		}
	}

	if len(tags) > 0 {

		err = common.DbBatchUpsert[model.Tag](ctx, common.GetDaprClient(), tags, model.TagTableInfo.Name, model.Tag_FIELD_NAME_id)
		if err != nil {
			err = errors.Wrap(err, "db upsert tag error")
			return
		}
	}

	return
}

func (p *DeviceExcelFileProcessor) getVendorPinyinKey(vendor string) string {
	pinyins := pinyin.LazyConvert(vendor, nil)
	result := strings.Trim(fmt.Sprint(pinyins), "[]")
	result = strings.Replace(result, " , ", "", -1)
	result = strings.Replace(result, " ", "", -1)
	if result == "" { //英文的
		result = vendor
	}
	return result
}

func (p *DeviceExcelFileProcessor) CalcId(col_value string) (targetVal string) {
	h := md5.New()
	h.Write([]byte(col_value))
	re := h.Sum(nil)
	targetVal = hex.EncodeToString(re)
	return
}
