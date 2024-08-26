package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitProductRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/product/page", ProductPageListHandler)
	r.Get(common.BASE_CONTEXT+"/product", ProductListHandler)
	r.Post(common.BASE_CONTEXT+"/product", UpsertProductHandler)
	r.Delete(common.BASE_CONTEXT+"/product/{id}", DeleteProductHandler)
	r.Post(common.BASE_CONTEXT+"/product/batch-delete", batchDeleteProductHandler)
	r.Post(common.BASE_CONTEXT+"/product/batch-upsert", batchUpsertProductHandler)
	r.Get(common.BASE_CONTEXT+"/product/groupby", ProductGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Product
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /product/groupby [get]
func ProductGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_product")
}

// @Summary batch update
// @Description batch update
// @Tags Product
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /product/batch-upsert [post]
func batchUpsertProductHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.ProductTableInfo.Name, model.Product_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Product
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param json_data query string false "json_data"
// @Param vendor query string false "vendor"
// @Param model query string false "model"
// @Param version query string false "version"
// @Param identifier query string false "identifier"
// @Param categories query string false "categories"
// @Param descript query string false "descript"
// @Param type query string false "type"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Product}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /product/page [get]
func ProductPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Product](w, r, common.GetDaprClient(), "o_product", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Product
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param json_data query string false "json_data"
// @Param vendor query string false "vendor"
// @Param model query string false "model"
// @Param version query string false "version"
// @Param identifier query string false "identifier"
// @Param categories query string false "categories"
// @Param descript query string false "descript"
// @Param type query string false "type"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Product} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /product [get]
func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Product](w, r, common.GetDaprClient(), "o_product", "id")
}

// @Summary save
// @Description save
// @Tags Product
// @Accept       json
// @Param item body model.Product true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Product} "object"
// @Failure 500 {object} common.Response ""
// @Router /product [post]
func UpsertProductHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Product
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Product")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Product)
	}

	err = common.DbUpsert[model.Product](r.Context(), common.GetDaprClient(), val, model.ProductTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Product
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Product} "object"
// @Failure 500 {object} common.Response ""
// @Router /product/{id} [delete]
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Product")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_product", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Product
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /product/batch-delete [post]
func batchDeleteProductHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Product")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_product", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
