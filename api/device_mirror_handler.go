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

func InitCustomDeviceMirrorRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/device-mirror/{identifier}", getDeviceMirrorHandler)

}

// @Summary 获取实时device-mirror
// @Description  获取实时device-mirror
// @Tags Device_mirror
// @Accept  json
// @Param identifier query string true "identifier"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-mirror/{identifier} [get]
func getDeviceMirrorHandler(w http.ResponseWriter, r *http.Request) {
	var identifier = chi.URLParam(r, "identifier")

	data, err := service.GetDeviceMirror(r.Context(), identifier)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
	}

	common.HttpResult(w, common.OK.WithData(data))
}
