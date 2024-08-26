package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"things-service/service"
)

func InitCustomDeviceDataRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/device-history-data", DeviceHistoryDataHandler)

}

// @Summary 查询设备历史数据
// @Description 查询设备历史数据
// @Tags DeviceData
// @Param ids query string true "ids,逗号分隔"
// @param start_time query string true "开始时间,2006-01-02 15:04:05"
// @param end_time query string true "结束时间,2006-01-02 15:04:05"
// @Produce  json
// @Success 200 {object} common.Response{data=[]entity.DeviceHistoryData} "实例的数组"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /device-history-data [get]
func DeviceHistoryDataHandler(w http.ResponseWriter, r *http.Request) {
	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")
	ids := r.URL.Query().Get("ids")
	if ids == "" || startTime == "" || endTime == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("参数错误"))
		return
	}

	datas, err := service.GetHistoryData(r.Context(), ids, startTime, endTime)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("获取数据错误").AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(datas))
}
