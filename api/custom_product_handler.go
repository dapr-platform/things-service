package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"io"
	"net/http"
	"os"
	"strings"
	"things-service/entity"
	"things-service/model"
	"things-service/service"
)

var (
	_ = null.String{}
)

type ProductAddDeviceReq struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	ProductId  string `json:"product_id"`
}

func InitCustomProductRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/product/add-device", ProductAddDeviceHandler)
	r.Post(common.BASE_CONTEXT+"/product/import-devices", importExcelHandler)
	r.Post(common.BASE_CONTEXT+"/product/batch-add-devices", batchProductUpsertDeviceHandler)
	r.Post(common.BASE_CONTEXT+"/product/batch-set-devices-tag", batchProductSetDevicesTagHandler)
}

// @Summary 添加设备
// @Description 添加设备
// @Tags Product
// @Param req body ProductAddDeviceReq true "req"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Device} "设备"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /product/add-device [post]
func ProductAddDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var req ProductAddDeviceReq
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("read body error"))
		return
	}
	sub, _ := common.ExtractUserSub(r)
	device, err := service.ProductAddDevice(r.Context(), sub, req.Identifier, req.Name, req.ProductId)
	if err != nil {

		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK.WithData(device))
}

// @Summary 批量添加设备的tags
// @Description 批量添加设备的tags
// @Tags Product
// @Accept  json
// @Param reqs body []entity.BatchTagInfo true "对象数组"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /product/batch-set-devices-tag [post]
func batchProductSetDevicesTagHandler(w http.ResponseWriter, r *http.Request) {

	var entities []entity.BatchTagInfo
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("entities is null"))
		return
	}
	for _, e := range entities {
		err = service.BatchSaveTag(r.Context(), e, false)
		if err != nil {
			common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
			return
		}
	}

	common.HttpResult(w, common.OK)
}

// @Summary 批量添加设备
// @Description 批量添加设备
// @Tags Product
// @Accept  json
// @Param entities body []model.Device true "对象数组"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /product/batch-add-devices [post]
func batchProductUpsertDeviceHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Device
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("entities is null"))
		return
	}

	sub, _ := common.ExtractUserSub(r)
	err = service.ProductBatchAddDevice(r.Context(), sub, entities)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK)
}

// @Summary 导入设备
// @Description 导入设备
// @Tags Product
// @Param file formData file true "文件"
// @Param product_id formData string true "产品id"
// @Produce  application/json
// @Success 200 {object} common.Response "{"status":0,"data":{},"msg":"success"}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /product/import-devices [post]
func importExcelHandler(w http.ResponseWriter, r *http.Request) {
	sub, _ := common.ExtractUserSub(r)
	productId := r.FormValue("product_id")
	if productId == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("product_id is null"))
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("read file error"))
		return
	}
	common.Logger.Debug("handler.Filename:", handler.Filename)
	newName := common.NanoId() + handler.Filename[strings.LastIndex(handler.Filename, "."):]
	common.Logger.Debug("newName=", newName)
	f, err := os.OpenFile(newName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		common.Logger.Error("打开文件失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg("open file error"))

		return
	}

	io.Copy(f, file)

	err = service.ProductImportExcel(r.Context(), productId, sub, newName)
	defer func() {
		f.Close()
		os.Remove(newName)
	}()
	if err != nil {
		common.Logger.Error("处理文件失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	} else {
		//event.ConstructAndSendEvent(context.Background(), common.EventTypePlatform, common.EventSubTypeService, "process excel "+handler.Filename+" error", "", common.EventStatusClosed, common.EventLevelMajor, time.Now(), "ops-service", "ops-service", "db-excel-upload")

	}

	common.HttpResult(w, common.OK)
}
