package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitSim_deviceRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/sim-device/page", Sim_devicePageListHandler)
	r.Get(common.BASE_CONTEXT+"/sim-device", Sim_deviceListHandler)
	r.Post(common.BASE_CONTEXT+"/sim-device", UpsertSim_deviceHandler)
	r.Delete(common.BASE_CONTEXT+"/sim-device/{id}", DeleteSim_deviceHandler)
	r.Post(common.BASE_CONTEXT+"/sim-device/batch-delete", batchDeleteSim_deviceHandler)
	r.Post(common.BASE_CONTEXT+"/sim-device/batch-upsert", batchUpsertSim_deviceHandler)
	r.Get(common.BASE_CONTEXT+"/sim-device/groupby", Sim_deviceGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Sim_device
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /sim-device/groupby [get]
func Sim_deviceGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_sim_device")
}

// @Summary batch update
// @Description batch update
// @Tags Sim_device
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /sim-device/batch-upsert [post]
func batchUpsertSim_deviceHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Sim_deviceTableInfo.Name, model.Sim_device_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Sim_device
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
// @Param parent_id query string false "parent_id"
// @Param group_id query string false "group_id"
// @Param product_id query string false "product_id"
// @Param protocol_config query string false "protocol_config"
// @Param identifier query string false "identifier"
// @Param enabled query string false "enabled"
// @Param rule_data query string false "rule_data"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Sim_device}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /sim-device/page [get]
func Sim_devicePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Sim_device](w, r, common.GetDaprClient(), "o_sim_device", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Sim_device
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param parent_id query string false "parent_id"
// @Param group_id query string false "group_id"
// @Param product_id query string false "product_id"
// @Param protocol_config query string false "protocol_config"
// @Param identifier query string false "identifier"
// @Param enabled query string false "enabled"
// @Param rule_data query string false "rule_data"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Sim_device} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /sim-device [get]
func Sim_deviceListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Sim_device](w, r, common.GetDaprClient(), "o_sim_device", "id")
}

// @Summary save
// @Description save
// @Tags Sim_device
// @Accept       json
// @Param item body model.Sim_device true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Sim_device} "object"
// @Failure 500 {object} common.Response ""
// @Router /sim-device [post]
func UpsertSim_deviceHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Sim_device
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Sim_device")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Sim_device)
	}

	err = common.DbUpsert[model.Sim_device](r.Context(), common.GetDaprClient(), val, model.Sim_deviceTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Sim_device
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Sim_device} "object"
// @Failure 500 {object} common.Response ""
// @Router /sim-device/{id} [delete]
func DeleteSim_deviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Sim_device")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_sim_device", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Sim_device
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /sim-device/batch-delete [post]
func batchDeleteSim_deviceHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Sim_device")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_sim_device", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
