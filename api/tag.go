package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitTagRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/tag/page", TagPageListHandler)
	r.Get(common.BASE_CONTEXT+"/tag", TagListHandler)
	r.Post(common.BASE_CONTEXT+"/tag", UpsertTagHandler)
	r.Delete(common.BASE_CONTEXT+"/tag/{id}", DeleteTagHandler)
	r.Post(common.BASE_CONTEXT+"/tag/batch-delete", batchDeleteTagHandler)
	r.Post(common.BASE_CONTEXT+"/tag/batch-upsert", batchUpsertTagHandler)
	r.Get(common.BASE_CONTEXT+"/tag/groupby", TagGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Tag
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /tag/groupby [get]
func TagGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_tag")
}

// @Summary batch update
// @Description batch update
// @Tags Tag
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /tag/batch-upsert [post]
func batchUpsertTagHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.TagTableInfo.Name, model.Tag_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Tag
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param rel_id query string false "rel_id"
// @Param key query string false "key"
// @Param value query string false "value"
// @Param editable query string false "editable"
// @Param rel_type query string false "rel_type"
// @Param user_id query string false "user_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Tag}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /tag/page [get]
func TagPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Tag](w, r, common.GetDaprClient(), "o_tag", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Tag
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param rel_id query string false "rel_id"
// @Param key query string false "key"
// @Param value query string false "value"
// @Param editable query string false "editable"
// @Param rel_type query string false "rel_type"
// @Param user_id query string false "user_id"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Tag} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /tag [get]
func TagListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Tag](w, r, common.GetDaprClient(), "o_tag", "id")
}

// @Summary save
// @Description save
// @Tags Tag
// @Accept       json
// @Param item body model.Tag true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Tag} "object"
// @Failure 500 {object} common.Response ""
// @Router /tag [post]
func UpsertTagHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Tag
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Tag")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Tag)
	}

	err = common.DbUpsert[model.Tag](r.Context(), common.GetDaprClient(), val, model.TagTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Tag
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Tag} "object"
// @Failure 500 {object} common.Response ""
// @Router /tag/{id} [delete]
func DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Tag")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_tag", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Tag
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /tag/batch-delete [post]
func batchDeleteTagHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Tag")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_tag", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
