package service

import (
	"context"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"
	"strings"

	"things-service/entity"
	"things-service/model"
	"time"
)

func init() {
	common.RegisterUpsertBeforeHook("Device_model", ProcessUpsertDeviceModel)
	common.RegisterUpsertBeforeHook("Product", ProcessUpsertProduct)
}

func GetDeviceProductModel(ctx context.Context, identifier string) (pmodel *entity.ProductModel, status int, err error) {
	dbEntity, err := common.DbGetOne[model.Device_identifier_product_json](ctx, common.GetDaprClient(), model.Device_identifier_product_jsonTableInfo.Name, model.Device_identifier_product_json_FIELD_NAME_identifier+"="+identifier)
	if err != nil {
		err = errors.Wrap(err, "GetDeviceProductModel "+identifier+" error")
		return
	}
	if dbEntity == nil {
		err = errors.New("GetDeviceProductModel " + identifier + " not found")
		return
	}
	status = int(dbEntity.Status)
	pmodel = &entity.ProductModel{}
	err = json.Unmarshal([]byte(dbEntity.JSONData), pmodel)

	if err != nil {
		err = errors.Wrap(err, "GetDeviceProductModel Unmarshal "+dbEntity.Identifier)
		common.Logger.Debug(dbEntity.JSONData)
	}
	return
}

func GetDeviceProductModelJsonString(ctx context.Context, identifier string) (jsonStr string, err error) {
	dbEntity, err := common.DbGetOne[model.Device_identifier_product_json](ctx, common.GetDaprClient(), model.Device_identifier_product_jsonTableInfo.Name, model.Device_identifier_product_json_FIELD_NAME_identifier+"="+identifier)
	if err != nil {
		err = errors.Wrap(err, "GetDeviceProductModel "+identifier+" error")
		return
	}
	if dbEntity == nil {
		err = errors.New("GetDeviceProductModel " + identifier + " not found")
		return
	}
	jsonStr = dbEntity.JSONData
	return
}

func ProcessUpsertProduct(r *http.Request, in any) (out any, err error) {
	sub, _ := common.ExtractUserSub(r)

	inModel := in.(model.Product)
	if inModel.ID == "" {
		inModel.ID = common.NanoId()
	}
	if inModel.CreatedBy == "" {
		inModel.CreatedBy = sub
	}
	zeroTime, _ := time.Parse("2006-01-02 15:04:05", "0001-01-01 00:00:00")
	if inModel.CreatedTime == common.LocalTime(zeroTime) {
		inModel.CreatedTime = common.LocalTime(time.Now())
	}
	inModel.UpdatedBy = sub
	inModel.UpdatedTime = common.LocalTime(time.Now())
	out = inModel

	jsonStr := inModel.JSONData
	pmodel := &entity.ProductModel{}
	err = json.Unmarshal([]byte(jsonStr), pmodel)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal product json error")
		return
	}
	//TODO 如果设备很多，改成异步的方式
	err = processProductTags(r.Context(), inModel.ID, pmodel.Tags)
	return
}

func processProductTags(ctx context.Context, productId string, tags []entity.PTag) (err error) {

	qstr := "_select=id,type&"
	qstr += model.Device_info_FIELD_NAME_product_id + "=" + productId
	devices, err := common.DbQuery[map[string]any](ctx, common.GetDaprClient(), model.Device_infoTableInfo.Name, qstr)
	if err != nil {
		err = errors.Wrap(err, "get product "+productId+" devices error")
		return
	}
	dbTags := make([]entity.Tag, 0)
	for _, t := range tags {
		dbTags = append(dbTags, entity.Tag{
			Key:      t.Label,
			Value:    t.Value,
			Editable: 1,
		})
	}
	for _, d := range devices {
		req := entity.BatchTagInfo{
			RelId:   cast.ToString(d["id"]),
			RelType: cast.ToInt(d["type"]),
			Tags:    dbTags,
		}
		err = BatchSaveTag(ctx, req, false)
		if err != nil {
			err = errors.Wrap(err, " BatchSaveTag ")
			return
		}
	}

	return
}

func ProcessUpsertDeviceModel(r *http.Request, in any) (out any, err error) {
	sub, _ := common.ExtractUserSub(r)

	inModel := in.(model.Device_model)
	if inModel.ID == "" {
		inModel.ID = common.NanoId()
	}
	if inModel.CreatedBy == "" {
		inModel.CreatedBy = sub
	}
	zeroTime, _ := time.Parse("2006-01-02 15:04:05", "0001-01-01 00:00:00")
	if inModel.CreatedTime == common.LocalTime(zeroTime) {
		inModel.CreatedTime = common.LocalTime(time.Now())
	}
	inModel.UpdatedBy = sub
	inModel.UpdatedTime = common.LocalTime(time.Now())
	categoryType := inModel.Categories
	if categoryType == "" {
		err = errors.New("Categories is blank")
		return
	}
	jsonData := inModel.JSONData
	err = processDeviceModelJsonData(jsonData, categoryType, time.Now().Unix())
	if err != nil {
		err = errors.Wrap(err, "processDeviceModelJsonData")
		return
	}

	out = inModel
	return
}
func RebuildDeviceModelMetaData(ctx context.Context) (err error) {
	tag := time.Now().Unix()
	deviceModels, err := common.DbQuery[model.Device_model](ctx, common.GetDaprClient(), model.Device_modelTableInfo.Name, "")
	if err != nil {
		err = errors.Wrap(err, "RebuildDeviceModelMetaData")
		return
	}
	for _, m := range deviceModels {
		categoryType := m.Categories
		if categoryType == "" {
			common.Logger.Error("categoryType is blank ")
			continue
		}
		jsonData := m.JSONData
		err = processDeviceModelJsonData(jsonData, categoryType, tag)
		if err != nil {
			common.Logger.Error(err.Error())
			continue
		}
	}
	err = common.DbDeleteByOps(ctx, common.GetDaprClient(), model.Model_metaTableInfo.Name, []string{model.Model_meta_FIELD_NAME_tag}, []string{"<"}, []any{tag})
	return
}
func processDeviceModelJsonData(jsonData, categoryType string, tag int64) (err error) {
	var entityModel entity.DeviceModelRaw
	err = json.Unmarshal([]byte(jsonData), &entityModel)
	if err != nil {
		return
	}

	categoryArr := strings.Split(categoryType, "/")
	category := categoryArr[len(categoryArr)-1]

	for _, p := range entityModel.Properties {
		err = processOneMeta(category, "properties", p, tag)
		if err != nil {
			return
		}
	}
	for _, p := range entityModel.Services {
		err = processOneMeta(category, "services", p, tag)
		if err != nil {
			return
		}
	}
	for _, p := range entityModel.Events {
		err = processOneMeta(category, "events", p, tag)
		if err != nil {
			return
		}
	}
	return
}

func processOneMeta(category string, metaType string, p map[string]any, tag int64) (err error) {
	identifier := p[entity.MODEL_IDENTIFIER_KEY]
	name := p[entity.MODEL_NAME_KEY]
	pbuf, _ := json.Marshal(p)
	if identifier == "" || name == "" {
		err = errors.New("identifier or name is blank " + string(pbuf))
		return
	}
	mAttr := model.Model_meta{
		ID:         common.GetMD5Hash(category + "_" + identifier.(string)),
		Name:       name.(string),
		Identifier: identifier.(string),
		MetaType:   metaType,
		Category:   category,
		JSONData:   string(pbuf),
		Tag:        tag,
	}
	err = common.DbUpsert[model.Model_meta](context.Background(), common.GetDaprClient(), mAttr, model.Model_metaTableInfo.Name, model.Model_meta_FIELD_NAME_id)
	if err != nil {
		err = errors.Wrap(err, "dbupsert "+metaType+" error")
		return
	}
	return
}
