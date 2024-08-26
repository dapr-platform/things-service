package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"github.com/spf13/cast"
	"net/http"

	"things-service/entity"
	"things-service/model"
)

var (
	_ = null.String{}
)

func InitUser_device_infoRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/user-device-info/page", User_device_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/user-device-info", User_device_infoListHandler)
	r.Get(common.BASE_CONTEXT+"/user-device-info/status-statistics", User_device_status_statisticsHandler)
}

// @Summary 用户设备状态统计
// @Description 用户设备状态统计
// @Tags User_device_info
// @Produce  json
// @Success 200 {object} common.Response{data=entity.DeviceStatusStatistics} "统计"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-device-info/status-statistics [get]
func User_device_status_statisticsHandler(w http.ResponseWriter, r *http.Request) {
	sub, _ := common.ExtractUserSub(r)
	qstr := "_select=status,sum:status&_groupby=status&user_id=" + sub
	common.Logger.Debugf("sub=%s", sub)
	datas, err := common.DbQuery[map[string]any](r.Context(), common.GetDaprClient(), model.User_device_infoTableInfo.Name, qstr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.Logger.Debugf("datas=%+v", datas)
	f := 0
	a := 0
	l := 0
	n := 0
	for _, d := range datas {
		switch cast.ToInt(d["status"]) {
		case 0:
			l = cast.ToInt(d["sum"])
		case 1:
			n = cast.ToInt(d["sum"])
		case 2:
			a = cast.ToInt(d["sum"])
		case 3:
			f = cast.ToInt(d["sum"])
		}
	}
	var demo = entity.DeviceStatusStatistics{
		Fault:  f,
		Alert:  a,
		Loss:   l,
		Normal: n,
	}
	common.HttpResult(w, common.OK.WithData(demo).AppendMsg("real data"))
}

// @Summary 分页查询所有实例
// @Description 分页查询所有实例, 可设置_page(从1 开始), _page_size, _order, 以及查询条件等，例如 status=1, name=$like.%CAM% 等
// @Tags User_device_info
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param user_id query string false "user_id"
// @Param device_id query string false "device_id"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Param index query string false "index"
// @Param locked query string false "locked"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param enabled query string false "enabled"
// @Param identifier query string false "identifier"
// @Param status query string false "status"
// @Param json_data query string false "json_data"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]entity.UserDeviceInfo}} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-device-info/page [get]
func User_device_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("参数错误"))
		return
	}
	sub, _ := common.ExtractUserSub(r)
	v := r.URL.Query()
	v.Add("user_id", sub)
	r.URL.RawQuery = v.Encode()

	common.CommonPageQuery[entity.UserDeviceInfo](w, r, common.GetDaprClient(), "v_user_device_info", "id")

}

// @Summary 查询所有实例
// @Description 查询所有实例, 可设置order, 以及查询条件等，例如 status=1, name=$like.%CAM% 等
// @Tags User_device_info
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param user_id query string false "user_id"
// @Param device_id query string false "device_id"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Param index query string false "index"
// @Param locked query string false "locked"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param enabled query string false "enabled"
// @Param identifier query string false "identifier"
// @Param status query string false "status"
// @Param json_data query string false "json_data"
// @Produce  json
// @Success 200 {object} common.Response{data=[]entity.UserDeviceInfo} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-device-info [get]
func User_device_infoListHandler(w http.ResponseWriter, r *http.Request) {
	sub, _ := common.ExtractUserSub(r)
	v := r.URL.Query()
	v.Add("user_id", sub)
	r.URL.RawQuery = v.Encode()
	common.CommonQuery[entity.UserDeviceInfo](w, r, common.GetDaprClient(), "v_user_device_info", "id")
}
