package api

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"things-service/service"

	"net/http"
	"os"
)

func initDbExcelRouter(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/db-excel-upload", dbExcelUploadHandler)
}

// @Summary 上传数据库格式静态文件
// @Description 上传数据库格式静态文件，注意sheet名称和表名称必须一致,columns必须与数据库字段一致,数据表要有个字符串字段id，计算规则为column前加*字段的md5
// @Tags Ops
// @Produce  application/json
// @Param file formData file true "文件"
// @Param key formData string true "key"
// @Success 200 {object} common.Response "{"status":200,"data":{},"msg":"success"}"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /db-excel-upload [post]
func dbExcelUploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("read file error"))
		return
	}
	defer file.Close()
	log.Println("handler.Filename:", handler.Filename)
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("打开文件失败", err)
		common.HttpResult(w, common.ErrService.AppendMsg("open file error"))

		return
	}
	defer f.Close()
	io.Copy(f, file)
	key := r.FormValue("key")
	go func() {
		processor := service.NewDbExcelFileUploadProcessor()
		err = processor.Process(context.Background(), handler.Filename, key)
		if err != nil {
			log.Println("处理文件失败", err)
		} else {
			//event.ConstructAndSendEvent(context.Background(), common.EventTypePlatform, common.EventSubTypeService, "process excel "+handler.Filename+" error", "", common.EventStatusClosed, common.EventLevelMajor, time.Now(), "nbif-ops", "nbif-ops", "db-excel-upload")

		}
		os.Remove(handler.Filename)
	}()

	common.HttpResult(w, common.OK)
}
