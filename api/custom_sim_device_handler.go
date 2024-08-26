package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"things-service/service"
)

func InitCustomSim_deviceRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/sim-device/change-enabled", changeEnabledHandler)
}

// @Summary 设置设备enabled
// @Description 设置设备enabled
// @Tags Sim_device
// @Accept  json
// @Param id query string true "设备id"
// @Param enabled query string true "1:enabled,0:disabled"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /sim-device/change-enabled [get]
func changeEnabledHandler(w http.ResponseWriter, r *http.Request) {
	sub, _ := common.ExtractUserSub(r)
	id := r.URL.Query().Get("id")
	enabled := r.URL.Query().Get("enabled")
	if id == "" || enabled == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	enabledInt, err := strconv.Atoi(enabled)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	err = service.ChangeSimDeviceEnabled(r.Context(), sub, id, enabledInt)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
