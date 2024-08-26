package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitAlarm_ruleRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/alarm-rule/page", Alarm_rulePageListHandler)
	r.Get(common.BASE_CONTEXT+"/alarm-rule", Alarm_ruleListHandler)
	r.Post(common.BASE_CONTEXT+"/alarm-rule", UpsertAlarm_ruleHandler)
	r.Delete(common.BASE_CONTEXT+"/alarm-rule/{id}", DeleteAlarm_ruleHandler)
	r.Post(common.BASE_CONTEXT+"/alarm-rule/batch-delete", batchDeleteAlarm_ruleHandler)
	r.Post(common.BASE_CONTEXT+"/alarm-rule/batch-upsert", batchUpsertAlarm_ruleHandler)
	r.Get(common.BASE_CONTEXT+"/alarm-rule/groupby", Alarm_ruleGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Alarm_rule
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /alarm-rule/groupby [get]
func Alarm_ruleGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_alarm_rule")
}

// @Summary batch update
// @Description batch update
// @Tags Alarm_rule
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /alarm-rule/batch-upsert [post]
func batchUpsertAlarm_ruleHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Alarm_ruleTableInfo.Name, model.Alarm_rule_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Alarm_rule
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param create_id query string false "create_id"
// @Param update_id query string false "update_id"
// @Param name query string false "name"
// @Param status query string false "status"
// @Param content query string false "content"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Alarm_rule}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /alarm-rule/page [get]
func Alarm_rulePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Alarm_rule](w, r, common.GetDaprClient(), "o_alarm_rule", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Alarm_rule
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param create_id query string false "create_id"
// @Param update_id query string false "update_id"
// @Param name query string false "name"
// @Param status query string false "status"
// @Param content query string false "content"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Alarm_rule} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /alarm-rule [get]
func Alarm_ruleListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Alarm_rule](w, r, common.GetDaprClient(), "o_alarm_rule", "id")
}

// @Summary save
// @Description save
// @Tags Alarm_rule
// @Accept       json
// @Param item body model.Alarm_rule true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Alarm_rule} "object"
// @Failure 500 {object} common.Response ""
// @Router /alarm-rule [post]
func UpsertAlarm_ruleHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Alarm_rule
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Alarm_rule")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Alarm_rule)
	}

	err = common.DbUpsert[model.Alarm_rule](r.Context(), common.GetDaprClient(), val, model.Alarm_ruleTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Alarm_rule
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Alarm_rule} "object"
// @Failure 500 {object} common.Response ""
// @Router /alarm-rule/{id} [delete]
func DeleteAlarm_ruleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Alarm_rule")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_alarm_rule", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Alarm_rule
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /alarm-rule/batch-delete [post]
func batchDeleteAlarm_ruleHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Alarm_rule")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_alarm_rule", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
