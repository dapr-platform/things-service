package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"net/http"
	"things-service/service"
)

var (
	_ = null.String{}
)

func InitCustomDevice_modelRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/device-model/rebuild-metadata", rebuildDevice_modelMetadataHandler)
}

// @Summary 根据所有的物模型，重建meta
// @Description 根据所有的物模型，重建meta, 维护使用
// @Tags Device_model
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-model/rebuild-metadata [post]
func rebuildDevice_modelMetadataHandler(w http.ResponseWriter, r *http.Request) {
	err := service.RebuildDeviceModelMetaData(r.Context())
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
