package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"things-service/model"
)

func InitDevice_mirrorRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/device-mirror/page", Device_mirrorPageListHandler)
	r.Get(common.BASE_CONTEXT+"/device-mirror", Device_mirrorListHandler)
	r.Post(common.BASE_CONTEXT+"/device-mirror", UpsertDevice_mirrorHandler)
	r.Delete(common.BASE_CONTEXT+"/device-mirror/{id}", DeleteDevice_mirrorHandler)
	r.Post(common.BASE_CONTEXT+"/device-mirror/batch-delete", batchDeleteDevice_mirrorHandler)
	r.Post(common.BASE_CONTEXT+"/device-mirror/batch-upsert", batchUpsertDevice_mirrorHandler)
	r.Get(common.BASE_CONTEXT+"/device-mirror/groupby", Device_mirrorGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Device_mirror
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device-mirror/groupby [get]
func Device_mirrorGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_device_mirror")
}

// @Summary batch update
// @Description batch update
// @Tags Device_mirror
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /device-mirror/batch-upsert [post]
func batchUpsertDevice_mirrorHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Device_mirrorTableInfo.Name, model.Device_mirror_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Device_mirror
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param json_data query string false "json_data"
// @Param updated_time query string false "updated_time"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Device_mirror}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device-mirror/page [get]
func Device_mirrorPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}
	common.CommonPageQuery[model.Device_mirror](w, r, common.GetDaprClient(), "o_device_mirror", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Device_mirror
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param json_data query string false "json_data"
// @Param updated_time query string false "updated_time"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Device_mirror} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /device-mirror [get]
func Device_mirrorListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Device_mirror](w, r, common.GetDaprClient(), "o_device_mirror", "id")
}

// @Summary save
// @Description save
// @Tags Device_mirror
// @Accept       json
// @Param item body model.Device_mirror true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Device_mirror} "object"
// @Failure 500 {object} common.Response ""
// @Router /device-mirror [post]
func UpsertDevice_mirrorHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Device_mirror
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Device_mirror")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Device_mirror)
	}

	err = common.DbUpsert[model.Device_mirror](r.Context(), common.GetDaprClient(), val, model.Device_mirrorTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Device_mirror
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Device_mirror} "object"
// @Failure 500 {object} common.Response ""
// @Router /device-mirror/{id} [delete]
func DeleteDevice_mirrorHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Device_mirror")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_device_mirror", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Device_mirror
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /device-mirror/batch-delete [post]
func batchDeleteDevice_mirrorHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Device_mirror")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_device_mirror", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
