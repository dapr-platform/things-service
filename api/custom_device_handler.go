package api

import (
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"github.com/pkg/errors"
	"net/http"
	"things-service/entity"
	"things-service/service"
)

var (
	_ = null.String{}
)

func InitCustomDeviceRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/device/service-invoke", serviceInvokeHandler)
	r.Post(common.BASE_CONTEXT+"/device/property-set", propertySetHandler)
	r.Post(common.BASE_CONTEXT+"/device/sim-device-mirror", simDeviceMirrorsHandler)
}

// @Summary 模拟发送device_mirror
// @Description 模拟发送device_mirror,调试使用
// @Tags Device
// @Accept  json
// @Param req body entity.SimDeviceMirrorReq true "req"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device/sim-device-mirror [post]
func simDeviceMirrorsHandler(w http.ResponseWriter, r *http.Request) {
	//sub, _ := common.ExtractUserSub(r)
	var req entity.SimDeviceMirrorReq
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("unmarshal req error"))
		return
	}
	var dm entity.DeviceMirror
	err = json.Unmarshal([]byte(req.JsonData), &dm)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("unmarshal json data error"))
		return
	}
	device, err := service.GetDeviceWithTagByIdentifier(r.Context(), req.DeviceIdentifier)
	if err != nil {
		err = errors.Wrap(err, "GetDeviceWithTagByIdentifier")
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	err = service.PersistDeviceMirror(r.Context(), req.DeviceIdentifier, &dm)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	err = service.SendDeviceMirrorToMessageChannel(r.Context(), req.DeviceIdentifier, device, &dm, "")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
	}
	common.HttpSuccess(w, common.OK)
}

// @Summary 属性设置
// @Description 属性设置, 异步调用，返回成功表示命令调用成功。
// @Tags Device
// @Accept  json
// @Param req body entity.PropertySetReq true "req"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device/property-set [post]
func propertySetHandler(w http.ResponseWriter, r *http.Request) {
	//sub, _ := common.ExtractUserSub(r)
	var req entity.PropertySetReq
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	err = service.ProcessPropertySet(r.Context(), req)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)
}

// @Summary 服务调用
// @Description 服务调用
// @Tags Device
// @Accept  json
// @Param req body entity.ServiceInvokeReq true "req"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device/service-invoke [post]
func serviceInvokeHandler(w http.ResponseWriter, r *http.Request) {
	//sub, _ := common.ExtractUserSub(r)
	var req entity.ServiceInvokeReq
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	//TODO
	common.HttpSuccess(w, common.OK)
}

func cleanModelMap[T any](m map[string]any) map[string]any {
	buf, _ := json.Marshal(m)
	var t T
	json.Unmarshal(buf, &t)
	var newMap map[string]any
	newBuf, _ := json.Marshal(t)
	json.Unmarshal(newBuf, &newMap)
	for k, _ := range newMap {
		_, ok := m[k]
		if !ok {
			delete(newMap, k)
		}
	}
	return newMap
}
