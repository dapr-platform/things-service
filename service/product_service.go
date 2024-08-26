package service

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"

	"things-service/model"
	"time"
)

func init() {
}

func ProductAddDevice(ctx context.Context, sub, identifier, name, productId string) (device *model.Device, err error) {
	device, err = getDeviceByIdentifier(ctx, sub, identifier, name, productId)
	if err != nil {
		err = errors.Wrap(err, "get device error")
		return
	}

	err = common.DbUpsert[model.Device](ctx, common.GetDaprClient(), *device, model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id)
	if err != nil {
		err = errors.Wrap(err, "upsert device error")
	}
	return
}
func getDeviceByIdentifier(ctx context.Context, sub, identifier, name, productId string) (device *model.Device, err error) {
	device, err = common.DbGetOne[model.Device](ctx, common.GetDaprClient(), model.DeviceTableInfo.Name, model.Device_FIELD_NAME_identifier+"="+identifier)
	if err != nil {
		err = errors.Wrap(err, "get device error")
		return
	}
	if device != nil {
		device.ProductID = productId
		device.UpdatedBy = sub
		device.Name = name
		device.UpdatedTime = common.LocalTime(time.Now())
	} else {
		device = &model.Device{
			ID:          common.GetMD5Hash(identifier),
			Identifier:  identifier,
			Name:        name,
			ProductID:   productId,
			CreatedTime: common.LocalTime(time.Now()),
			UpdatedTime: common.LocalTime(time.Now()),
			CreatedBy:   sub,
			UpdatedBy:   sub,
		}
	}
	return
}
func ProductBatchAddDevice(ctx context.Context, sub string, entities []model.Device) (err error) {
	productId := entities[0].ProductID
	if productId == "" {
		err = errors.New("product id is null")
		return
	}
	product, err := common.DbGetOne[model.Product](ctx, common.GetDaprClient(), model.ProductTableInfo.Name, model.Product_FIELD_NAME_id+"="+productId)
	if err != nil {
		err = errors.Wrap(err, "get product error")
		return
	}
	if product == nil {
		err = errors.New("product not found")
		return
	}

	dm := make([]model.Device, 0)
	for _, d := range entities {
		if d.Identifier == "" {
			err = errors.New("identifier is null")
			return
		}
		if d.ID == "" {
			d.ID = common.GetMD5Hash(d.Identifier)
			d.CreatedBy = sub
			d.CreatedTime = common.LocalTime(time.Now())
			d.Type = product.Type
		}
		if d.Name == "" {
			d.Name = d.Identifier
		}
		d.UpdatedBy = sub
		d.UpdatedTime = common.LocalTime(time.Now())
		dm = append(dm, d)
	}

	err = common.DbBatchUpsert[model.Device](ctx, common.GetDaprClient(), dm, model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id)
	if err != nil {
		err = errors.Wrap(err, "upsert device error")
		return
	}
	return
}
func ProductImportExcel(ctx context.Context, productId, sub, filePath string) (err error) {
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
		devices := make([]model.Device, 0)
		for i, row := range rows {
			if i == 0 { //标题行
				continue
			}
			if len(row) == 0 {
				continue
			}
			identifier := row[0]
			device, err := getDeviceByIdentifier(ctx, sub, identifier, identifier, productId)
			if err != nil {
				common.Logger.Error(err.Error())
				err = errors.Wrap(err, "get device error")
				return err
			}
			if device == nil {
				device = &model.Device{
					ID:          common.GetMD5Hash(identifier),
					Identifier:  identifier,
					Name:        identifier,
					ProductID:   productId,
					CreatedTime: common.LocalTime(time.Now()),
					UpdatedTime: common.LocalTime(time.Now()),
				}
			}
			devices = append(devices, *device)
		}
		err = common.DbBatchUpsert[model.Device](ctx, common.GetDaprClient(), devices, model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id)
		if err != nil {
			common.Logger.Error(err.Error())
			err = errors.Wrap(err, "upsert device error")
			return err
		}
	}

	return nil

}
