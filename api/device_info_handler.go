package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"net/http"
	"strconv"

	"things-service/entity"
	"things-service/service"
)

var (
	_ = null.String{}
)

func InitDevice_infoRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/device-info/page", Device_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/device-info", Device_infoListHandler)
	r.Get(common.BASE_CONTEXT+"/device-info/search-device", tagSearchDeviceHandler)
	r.Get(common.BASE_CONTEXT+"/device-info/fuzzy-search-device", fuzzySearchDeviceHandler)
}

// @Summary 根据标签查询设备
// @Description 根据标签查询设备, 如果要根据产品信息，则product_id 或 product_name 要设置一个
// @Tags Device_info
// @Param tags query string false "tags标签，key:value,key:value形式"
// @Param _page query int false "page"
// @Param _page_size query int false "pageSize"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param identifier query string false "identifier"
// @Param name query string false "name"
// @Param product_id query string false "product_id"
// @Param product_name query string false "product_name"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]entity.DeviceInfo}} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-info/search-device [get]
func tagSearchDeviceHandler(w http.ResponseWriter, r *http.Request) {
	tagsStr := r.URL.Query().Get("tags")
	pageS := r.URL.Query().Get("_page")
	pageSizeS := r.URL.Query().Get("_page_size")
	page := 1
	if pageS != "" {
		page, _ = strconv.Atoi(pageS)
	}
	if page < 1 {
		common.Logger.Error("page < 1, use 1")
		page = 1
	}
	pageSize := 20
	if pageSizeS != "" {
		pageSize, _ = strconv.Atoi(pageSizeS)
	}
	if pageSize < 1 {
		common.Logger.Error("pageSize < 1, use 20")
		pageSize = 20
	}
	order := r.URL.Query().Get("_order")
	name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")
	identifier := r.URL.Query().Get("identifier")
	productId := r.URL.Query().Get("product_id")
	productName := r.URL.Query().Get("product_name")
	datas, err := service.QueryDeviceByTagsAndProductIdAndNameIdentifier(r.Context(), page, pageSize, tagsStr, order, id, name, identifier, productId, productName)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(datas))
}

// @Summary 模糊搜索查询设备
// @Description 根据标签、输入字符串模糊搜索查询设备(从标签的value 和名字中搜索), 如果要根据产品信息，则product_id 或 product_name 要设置一个
// @Tags Device_info
// @Param tags query string false "tags标签，key:value,key:value形式"
// @Param q query string false "模糊搜索字符串"
// @Param _page query int false "page"
// @Param _page_size query int false "pageSize"
// @Param _order query string false "order"
// @Param product_id query string false "product_id"
// @Param product_name query string false "product_name"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]entity.DeviceInfo}} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-info/fuzzy-search-device [get]
func fuzzySearchDeviceHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	pageS := r.URL.Query().Get("_page")
	pageSizeS := r.URL.Query().Get("_page_size")
	page := 1
	if pageS != "" {
		page, _ = strconv.Atoi(pageS)
	}
	if page < 1 {
		common.Logger.Error("page < 1, use 1")
		page = 1
	}
	pageSize := 20
	if pageSizeS != "" {
		pageSize, _ = strconv.Atoi(pageSizeS)
	}
	if pageSize < 1 {
		common.Logger.Error("pageSize < 1, use 20")
		pageSize = 20
	}
	order := r.URL.Query().Get("_order")
	productId := r.URL.Query().Get("product_id")
	productName := r.URL.Query().Get("product_name")
	tagsStr := r.URL.Query().Get("tags")
	datas, err := service.QueryDeviceByFuzzyText(r.Context(), page, pageSize, tagsStr, q, order, productId, productName)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(datas))
}

// @Summary 分页查询所有实例
// @Description 分页查询所有实例, 可设置page, page_size, order, 以及查询条件等，例如 status=1, name=$like.%CAM% 等
// @Tags Device_info
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
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
// @Param enabled query string false "enabled"
// @Param identifier query string false "identifier"
// @Param product_name query string false "product_name"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]entity.DeviceInfo}} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-info/page [get]
func Device_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("参数错误"))
		return
	}
	if page != "" && pageSize != "" {
		common.CommonPageQuery[entity.DeviceInfo](w, r, common.GetDaprClient(), "v_device_with_tag", "id")
	}

}

// @Summary 查询所有实例
// @Description 查询所有实例, 可设置order, 以及查询条件等，例如 status=1, name=$like.%CAM% 等
// @Tags Device_info
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
// @Param enabled query string false "enabled"
// @Param identifier query string false "identifier"
// @Param product_name query string false "product_name"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Device_info} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-info [get]
func Device_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[entity.DeviceInfo](w, r, common.GetDaprClient(), "v_device_with_tag", "id")
}
