package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitModel_metaRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/model-meta/page", Model_metaPageListHandler)
	r.Get(common.BASE_CONTEXT+"/model-meta", Model_metaListHandler)
	r.Post(common.BASE_CONTEXT+"/model-meta", UpsertModel_metaHandler)
	r.Delete(common.BASE_CONTEXT+"/model-meta/{id}", DeleteModel_metaHandler)
	r.Post(common.BASE_CONTEXT+"/model-meta/batch-delete", batchDeleteModel_metaHandler)
	r.Post(common.BASE_CONTEXT+"/model-meta/batch-upsert", batchUpsertModel_metaHandler)
	r.Get(common.BASE_CONTEXT+"/model-meta/groupby", Model_metaGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Model_meta
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /model-meta/groupby [get]
func Model_metaGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_model_meta")
}

// @Summary batch update
// @Description batch update
// @Tags Model_meta
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /model-meta/batch-upsert [post]
func batchUpsertModel_metaHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Model_metaTableInfo.Name, model.Model_meta_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Model_meta
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param identifier query string false "identifier"
// @Param json_data query string false "json_data"
// @Param meta_type query string false "meta_type"
// @Param category query string false "category"
// @Param type query string false "type"
// @Param tag query string false "tag"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Model_meta}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /model-meta/page [get]
func Model_metaPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Model_meta](w, r, common.GetDaprClient(), "o_model_meta", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Model_meta
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param identifier query string false "identifier"
// @Param json_data query string false "json_data"
// @Param meta_type query string false "meta_type"
// @Param category query string false "category"
// @Param type query string false "type"
// @Param tag query string false "tag"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Model_meta} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /model-meta [get]
func Model_metaListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Model_meta](w, r, common.GetDaprClient(), "o_model_meta", "id")
}

// @Summary save
// @Description save
// @Tags Model_meta
// @Accept       json
// @Param item body model.Model_meta true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Model_meta} "object"
// @Failure 500 {object} common.Response ""
// @Router /model-meta [post]
func UpsertModel_metaHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Model_meta
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Model_meta")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Model_meta)
	}

	err = common.DbUpsert[model.Model_meta](r.Context(), common.GetDaprClient(), val, model.Model_metaTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Model_meta
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Model_meta} "object"
// @Failure 500 {object} common.Response ""
// @Router /model-meta/{id} [delete]
func DeleteModel_metaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Model_meta")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_model_meta", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Model_meta
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /model-meta/batch-delete [post]
func batchDeleteModel_metaHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Model_meta")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_model_meta", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
