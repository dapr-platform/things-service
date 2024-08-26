package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"things-service/entity"
	"things-service/monitor_client"
	"time"
)

func InitMetricsQueryHandler(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/metrics/query", MetricsQueryHandler)
	r.Post(common.BASE_CONTEXT+"/metrics/query-range", MetricsQueryRangeHandler)
	r.Post(common.BASE_CONTEXT+"/metrics/query-range-batch", MetricsQueryRangeBatchHandler)
}

// @Summary 查询batch
// @Description 查询batch
// @Tags Metrics
// @Accept  json
// @Param query body entity.BatchQueryForm true "查询条件"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /metrics/query-range-batch [post]
func MetricsQueryRangeBatchHandler(w http.ResponseWriter, r *http.Request) {
	var query entity.BatchQueryForm
	err := common.ReadRequestBody(r, &query)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	var lst []*monitor_client.QueryResult
	for _, item := range query.Queries {
		ss := time.Now().Add(-time.Hour)
		se := time.Now()
		if item.Start != "" {
			ss, err = time.ParseInLocation("2006-01-02 15:04:05", item.Start, time.Local)
			if err != nil {
				common.HttpResult(w, common.ErrParam.AppendMsg(item.Start+" "+err.Error()))
				return
			}
		}
		if item.End != "" {
			se, err = time.ParseInLocation("2006-01-02 15:04:05", item.End, time.Local)
			if err != nil {
				common.HttpResult(w, common.ErrParam.AppendMsg(item.End+" "+err.Error()))
				return
			}
		}
		if ss.After(se) {
			common.HttpResult(w, common.ErrParam.AppendMsg(item.Start+" "+item.End+" start after end"))
			return
		}
		step := time.Duration(15)
		if item.Step > 0 {
			step = time.Duration(item.Step)
		}
		rg := entity.Range{
			Start: ss,
			End:   se,
			Step:  step,
		}
		resp, err1 := monitor_client.QueryRange(r.Context(), item.Query, rg.Start, rg.End, rg.Step)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()).WithData(lst))
			return
		}
		lst = append(lst, resp)
	}
	common.HttpResult(w, common.OK.WithData(lst))
}

// @Summary 测试查询
// @Description 测试查询
// @Tags Metrics
// @Accept  json
// @Param query body entity.QueryMetrics true "查询条件"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /metrics/query [post]
func MetricsQueryHandler(w http.ResponseWriter, r *http.Request) {
	var query entity.QueryMetrics
	err := common.ReadRequestBody(r, &query)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", query.Time, time.Local)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	data, err := monitor_client.Query(r.Context(), query.Query, t)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	} else {
		common.HttpResult(w, common.OK.WithData(data))
		return
	}
}

// @Summary range查询
// @Description range查询
// @Tags Metrics
// @Accept  json
// @Param query body entity.QueryRangeMetrics true "查询条件"
// @Produce  json
// @Success 200 {object} common.Response "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /metrics/query-range [post]
func MetricsQueryRangeHandler(w http.ResponseWriter, r *http.Request) {
	var query entity.QueryRangeMetrics
	err := common.ReadRequestBody(r, &query)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	}
	ss := time.Now().Add(-time.Hour)
	se := time.Now()
	if query.Start != "" {
		ss, err = time.ParseInLocation("2006-01-02 15:04:05", query.Start, time.Local)
		if err != nil {
			common.HttpResult(w, common.ErrParam.AppendMsg("query.Start "+err.Error()))
			return
		}
	}
	if query.End != "" {
		se, err = time.ParseInLocation("2006-01-02 15:04:05", query.End, time.Local)
		if err != nil {
			common.HttpResult(w, common.ErrParam.AppendMsg("query.End "+err.Error()))
			return
		}
	}
	if ss.After(se) {
		common.HttpResult(w, common.ErrParam.AppendMsg("start after end"))
		return
	}
	step := time.Duration(15)
	if query.Step != 0 {
		step = time.Duration(query.Step)
	}

	data, err := monitor_client.QueryRange(r.Context(), query.Query, ss, se, step)
	if err != nil {
		common.HttpResult(w, common.ErrParam)
		return
	} else {
		common.HttpResult(w, common.OK.WithData(data))
		return
	}
}
