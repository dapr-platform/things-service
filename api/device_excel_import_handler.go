package api

import (
	"context"
	"encoding/base64"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"strings"
	"things-service/entity"
	"things-service/service"
)

func InitDeviceExcelImportRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/device-import/excel-import-binary", importExcelBinaryHandler)
	r.Post(common.BASE_CONTEXT+"/device-import/excel-modbus-import", importExcelModbusHandler)

}

// @Summary 导入excel
// @Description 导入excel,文件字节base64方式
// @Tags 设备导入
// @Param req body entity.FileUploadReq true "文件上传请求"
// @Produce  application/json
// @Success 200 {object} common.Response "{"status":0,"data":{},"msg":"success"}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-import/excel-import-binary [post]
func importExcelBinaryHandler(w http.ResponseWriter, r *http.Request) {
	var req entity.FileUploadReq
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.Logger.Error("读取文件失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg("read file binary error"))
		return
	}
	newName := common.NanoId() + "_" + req.FileName
	common.Logger.Debug("newName=", newName)
	f, err := os.OpenFile(newName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		common.Logger.Error("打开文件失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg("open file error"))

		return
	}
	fileBytes, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		common.Logger.Error("转换文件base64失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg("translate file binary error"))
		return
	}
	_, err = f.Write(fileBytes)
	if err != nil {
		common.Logger.Error("写入文件字节失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg("write file binary error"))
		return
	}
	processor := service.NewDeviceExcelFileProcessor()
	err = processor.ProcessModbusExcel(r.Context(), newName)
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

// @Summary 导入excel
// @Description 导入excel,modbus格式
// @Tags 设备导入
// @Param file formData file true "文件"
// @Produce  application/json
// @Success 200 {object} common.Response "{"status":0,"data":{},"msg":"success"}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-import/excel-modbus-import [post]
func importExcelModbusHandler(w http.ResponseWriter, r *http.Request) {
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

	go func() {
		processor := service.NewDeviceExcelFileProcessor()
		err = processor.ProcessModbusExcel(context.Background(), newName)
		if err != nil {
			common.Logger.Error("处理文件失败", err)
		}
		f.Close()
		os.Remove(newName)
	}()

	common.HttpResult(w, common.OK.AppendMsg("异步处理"))
}
