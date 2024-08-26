package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"things-service/entity"
	"things-service/service"
)

func InitCustomTagRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/tag/search-point-names", tagSearchPointNamesHandler)
	r.Get(common.BASE_CONTEXT+"/tag/search-device-by-values", valueSearchDeviceHandler)
	r.Get(common.BASE_CONTEXT+"/tag/all-key-with-value", tagAllKeyWithValueHandler)
	r.Get(common.BASE_CONTEXT+"/tag/all-value", tagAllValueHandler)
	r.Post(common.BASE_CONTEXT+"/tag/batch-save", tagBatchSaveHandler)
	r.Get(common.BASE_CONTEXT+"/tag/batch-get", tagBatchGetHandler)
}

// @Summary 批量获取一个对象的tag
// @Description 批量获取一个对象的tag
// @Tags Tag
// @Param rel_id query string true "对象id"
// @Produce  json
// @Success 200 {object} common.Response{data=entity.BatchTagInfo} "成功"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /tag/batch-get [get]
func tagBatchGetHandler(w http.ResponseWriter, r *http.Request) {
	relId := r.URL.Query().Get("rel_id")
	if relId == "" {
		common.HttpResult(w, common.ErrParam)
		return
	}

	info, err := service.BatchGetTag(r.Context(), relId)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(info))
}

// @Summary 批量保存一个对象的tag
// @Description 批量保存一个对象的tag
// @Tags Tag
// @Param req body entity.BatchTagInfo true "批量保存tag"
// @Produce  json
// @Success 200 {object} common.Response "成功"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /tag/batch-save [post]
func tagBatchSaveHandler(w http.ResponseWriter, r *http.Request) {
	var req entity.BatchTagInfo
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	err = service.BatchSaveTag(r.Context(), req, true)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)
}

// @Summary 根据标签查询Point 名称
// @Description 根据标签查询Point名称
// @Tags Tag
// @Param tags query string false "tags标签，key:value,key:value形式"
// @Produce  json
// @Success 200 {object} common.Response{data=[]string} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /tag/search-point-names [get]
func tagSearchPointNamesHandler(w http.ResponseWriter, r *http.Request) {
	tagsStr := r.URL.Query().Get("tags")

	datas, err := service.QueryPointNamesByTags(r.Context(), tagsStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(datas))
}

// @Summary 根据标签值 模糊搜索 设备数据
// @Description 根据标签值 模糊搜索 设备数据
// @Tags Tag
// @Param values query string false "tags value,空格分隔，模糊搜索"
// @Param _page query int false "page"
// @Param _page_size query int false "pageSize"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]entity.DeviceCurrentData}} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /tag/search-device-by-values [get]
func valueSearchDeviceHandler(w http.ResponseWriter, r *http.Request) {
	valuesStr := r.URL.Query().Get("values")
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

	datas, err := service.QueryDeviceByTagValueLike(r.Context(), page, pageSize, valuesStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(datas))
}

// @Summary 查询所有distinct标签key 和 values
// @Description 查询所有distinct标签key,不同的类型会有不同的标签，需要传设备类型。
// @Tags Tag
// @Produce  json
// @Param type query int false "设备类型,1:device,2:gateway,3:point,4:product 默认1"
// @Param product_id query string false "产品id，默认空，则是全部设备，否则为该产品下的所有设备"
// @Success 200 {object} common.Response{data=[]string} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /tag/all-key-with-value [get]
func tagAllKeyWithValueHandler(w http.ResponseWriter, r *http.Request) {
	relType := r.URL.Query().Get("type")
	if relType == "" {
		relType = "1"
	}
	productId := r.URL.Query().Get("product_id")
	datas, err := service.QueryAllTagKeyAndValue(r.Context(), productId, relType)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(datas))
}

// @Summary 根据标签key查询所有distinct value
// @Description 根据标签key查询所有distinct value
// @Tags Tag
// @Param tag query string true "key"
// @Produce  json
// @Success 200 {object} common.Response{data=[]string} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /tag/all-value [get]
func tagAllValueHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	if tag == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("tag is blank"))
		return
	}

	datas, err := service.QueryAllTagValue(r.Context(), tag)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(datas))
}
