package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"github.com/guregu/null"
	"net/http"
	"strings"

	"things-service/model"
	"time"
)

var (
	_ = null.String{}
)

func InitUser_deviceRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/user-device", UpsertUser_deviceHandler)
	r.Delete(common.BASE_CONTEXT+"/user-device/{id}", DeleteUser_deviceHandler)
	r.Post(common.BASE_CONTEXT+"/user-device/batch-delete", batchDeleteUser_deviceHandler)
	r.Post(common.BASE_CONTEXT+"/user-device/batch-upsert", batchUpsertUser_deviceHandler)
}

// @Summary 批量更新
// @Description 批量更新
// @Tags User_device
// @Accept  json
// @Param entities body []model.User_device true "对象数组"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-device/batch-upsert [post]
func batchUpsertUser_deviceHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.User_device
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam)
		return
	}
	sub, _ := common.ExtractUserSub(r)
	addEntities := make([]model.User_device, 0)
	for _, entity := range entities {
		if entity.UserID == "" {
			entity.UserID = sub
		}
		if entity.ID == "" {
			entity.ID = common.NanoId()
			entity.CreatedTime = common.LocalTime(time.Now())
		}

		entity.UpdatedTime = common.LocalTime(time.Now())
		addEntities = append(addEntities, entity)
	}

	err = common.DbBatchUpsert[model.User_device](r.Context(), common.GetDaprClient(), addEntities, model.User_deviceTableInfo.Name, model.User_device_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK.WithData(addEntities))
}

// @Summary 保存实例
// @Description 保存实例
// @Tags User_device
// @Accept       json
// @Param item body model.User_device true "实例全部信息"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User_device} "实例"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-device [post]
func UpsertUser_deviceHandler(w http.ResponseWriter, r *http.Request) {
	var val model.User_device
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	sub, _ := common.ExtractUserSub(r)

	if val.UserID == "" {
		val.UserID = sub
	}

	if val.ID == "" {
		val.ID = common.NanoId()
		val.CreatedTime = common.LocalTime(time.Now())
	}

	val.UpdatedTime = common.LocalTime(time.Now())

	err = common.DbUpsert[model.User_device](r.Context(), common.GetDaprClient(), val, model.User_deviceTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary 删除实例
// @Description 删除实例
// @Tags User_device
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User_device} "实例"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-device/{id} [delete]
func DeleteUser_deviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("User_device")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_user_device", "id", "id")
}

// @Summary 批量删除
// @Description 批量删除
// @Tags User_device
// @Accept  json
// @Param ids body []string true "id数组"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user-device/batch-delete [post]
func batchDeleteUser_deviceHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("User_device")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_user_device", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
