package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitKpi_infoRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/kpi-info/page", Kpi_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/kpi-info", Kpi_infoListHandler)
	r.Post(common.BASE_CONTEXT+"/kpi-info", UpsertKpi_infoHandler)
	r.Delete(common.BASE_CONTEXT+"/kpi-info/{id}", DeleteKpi_infoHandler)
	r.Post(common.BASE_CONTEXT+"/kpi-info/batch-delete", batchDeleteKpi_infoHandler)
	r.Post(common.BASE_CONTEXT+"/kpi-info/batch-upsert", batchUpsertKpi_infoHandler)
	r.Get(common.BASE_CONTEXT+"/kpi-info/groupby", Kpi_infoGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Kpi_info
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /kpi-info/groupby [get]
func Kpi_infoGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_kpi_info")
}

// @Summary batch update
// @Description batch update
// @Tags Kpi_info
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /kpi-info/batch-upsert [post]
func batchUpsertKpi_infoHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Kpi_infoTableInfo.Name, model.Kpi_info_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Kpi_info
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param label query string false "label"
// @Param description query string false "description"
// @Param unit query string false "unit"
// @Param product_name query string false "product_name"
// @Param interval query string false "interval"
// @Param type query string false "type"
// @Param calc_script query string false "calc_script"
// @Param summary_type query string false "summary_type"
// @Param value_type query string false "value_type"
// @Param org_id query string false "org_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Kpi_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /kpi-info/page [get]
func Kpi_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Kpi_info](w, r, common.GetDaprClient(), "o_kpi_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Kpi_info
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param label query string false "label"
// @Param description query string false "description"
// @Param unit query string false "unit"
// @Param product_name query string false "product_name"
// @Param interval query string false "interval"
// @Param type query string false "type"
// @Param calc_script query string false "calc_script"
// @Param summary_type query string false "summary_type"
// @Param value_type query string false "value_type"
// @Param org_id query string false "org_id"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Kpi_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /kpi-info [get]
func Kpi_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Kpi_info](w, r, common.GetDaprClient(), "o_kpi_info", "id")
}

// @Summary save
// @Description save
// @Tags Kpi_info
// @Accept       json
// @Param item body model.Kpi_info true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Kpi_info} "object"
// @Failure 500 {object} common.Response ""
// @Router /kpi-info [post]
func UpsertKpi_infoHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Kpi_info
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Kpi_info")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Kpi_info)
	}

	err = common.DbUpsert[model.Kpi_info](r.Context(), common.GetDaprClient(), val, model.Kpi_infoTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Kpi_info
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Kpi_info} "object"
// @Failure 500 {object} common.Response ""
// @Router /kpi-info/{id} [delete]
func DeleteKpi_infoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Kpi_info")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_kpi_info", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Kpi_info
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /kpi-info/batch-delete [post]
func batchDeleteKpi_infoHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Kpi_info")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_kpi_info", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
