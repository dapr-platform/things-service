package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitDeviceRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/device/page", DevicePageListHandler)
	r.Get(common.BASE_CONTEXT+"/device", DeviceListHandler)
	r.Post(common.BASE_CONTEXT+"/device", UpsertDeviceHandler)
	r.Delete(common.BASE_CONTEXT+"/device/{id}", DeleteDeviceHandler)
	r.Post(common.BASE_CONTEXT+"/device/batch-delete", batchDeleteDeviceHandler)
	r.Post(common.BASE_CONTEXT+"/device/batch-upsert", batchUpsertDeviceHandler)
	r.Get(common.BASE_CONTEXT+"/device/groupby", DeviceGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Device
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device/groupby [get]
func DeviceGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_device")
}

// @Summary batch update
// @Description batch update
// @Tags Device
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /device/batch-upsert [post]
func batchUpsertDeviceHandler(w http.ResponseWriter, r *http.Request) {

	var entities []map[string]any
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.DeviceTableInfo.Name, model.Device_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Device
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Param parent_id query string false "parent_id"
// @Param group_id query string false "group_id"
// @Param product_id query string false "product_id"
// @Param protocol_config query string false "protocol_config"
// @Param identifier query string false "identifier"
// @Param enabled query string false "enabled"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Device}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device/page [get]
func DevicePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Device](w, r, common.GetDaprClient(), "o_device", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Device
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Param parent_id query string false "parent_id"
// @Param group_id query string false "group_id"
// @Param product_id query string false "product_id"
// @Param protocol_config query string false "protocol_config"
// @Param identifier query string false "identifier"
// @Param enabled query string false "enabled"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Device} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device [get]
func DeviceListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Device](w, r, common.GetDaprClient(), "o_device", "id")
}

// @Summary save
// @Description save
// @Tags Device
// @Accept       json
// @Param item body model.Device true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Device} "object"
// @Failure 500 {object} common.Response ""
// @Router /device [post]
func UpsertDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Device
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Device")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Device)
	}

	err = common.DbUpsert[model.Device](r.Context(), common.GetDaprClient(), val, model.DeviceTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Device
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Device} "object"
// @Failure 500 {object} common.Response ""
// @Router /device/{id} [delete]
func DeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Device")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_device", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Device
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /device/batch-delete [post]
func batchDeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Device")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_device", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
