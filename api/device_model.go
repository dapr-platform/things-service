package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitDevice_modelRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/device-model/page", Device_modelPageListHandler)
	r.Get(common.BASE_CONTEXT+"/device-model", Device_modelListHandler)
	r.Post(common.BASE_CONTEXT+"/device-model", UpsertDevice_modelHandler)
	r.Delete(common.BASE_CONTEXT+"/device-model/{id}", DeleteDevice_modelHandler)
	r.Post(common.BASE_CONTEXT+"/device-model/batch-delete", batchDeleteDevice_modelHandler)
	r.Post(common.BASE_CONTEXT+"/device-model/batch-upsert", batchUpsertDevice_modelHandler)
	r.Get(common.BASE_CONTEXT+"/device-model/groupby", Device_modelGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Device_model
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device-model/groupby [get]
func Device_modelGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_device_model")
}

// @Summary batch update
// @Description batch update
// @Tags Device_model
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /device-model/batch-upsert [post]
func batchUpsertDevice_modelHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Device_modelTableInfo.Name, model.Device_model_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Device_model
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param descript query string false "descript"
// @Param categories query string false "categories"
// @Param json_data query string false "json_data"
// @Param service_script query string false "service_script"
// @Param property_script query string false "property_script"
// @Param event_script query string false "event_script"
// @Param cover_file query string false "cover_file"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Device_model}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device-model/page [get]
func Device_modelPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Device_model](w, r, common.GetDaprClient(), "o_device_model", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Device_model
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param descript query string false "descript"
// @Param categories query string false "categories"
// @Param json_data query string false "json_data"
// @Param service_script query string false "service_script"
// @Param property_script query string false "property_script"
// @Param event_script query string false "event_script"
// @Param cover_file query string false "cover_file"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Device_model} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device-model [get]
func Device_modelListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Device_model](w, r, common.GetDaprClient(), "o_device_model", "id")
}

// @Summary save
// @Description save
// @Tags Device_model
// @Accept       json
// @Param item body model.Device_model true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Device_model} "object"
// @Failure 500 {object} common.Response ""
// @Router /device-model [post]
func UpsertDevice_modelHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Device_model
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Device_model")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Device_model)
	}

	err = common.DbUpsert[model.Device_model](r.Context(), common.GetDaprClient(), val, model.Device_modelTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Device_model
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Device_model} "object"
// @Failure 500 {object} common.Response ""
// @Router /device-model/{id} [delete]
func DeleteDevice_modelHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Device_model")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_device_model", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Device_model
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /device-model/batch-delete [post]
func batchDeleteDevice_modelHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Device_model")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_device_model", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
