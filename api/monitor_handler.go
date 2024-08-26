package api

import (
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"

	"things-service/entity"
	"things-service/service/monitor"
)

func InitMonitorHandler(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/monitor/total", MonitorTotalHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/list/hosts", MonitorListHostHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/list/dbs", MonitorListDbHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/list/services", MonitorListServiceHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/list/gateways", MonitorListGatewayHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/list/interfaces", MonitorListInterfaceHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/log", MonitorLogHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/log/detail", MonitorLogDetailHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/log/label/values", MonitorLogLabelValuesHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/gateway/detail", MonitorGatewayDetailHandler)
	r.Get(common.BASE_CONTEXT+"/monitor/service/metrics_panel", MonitorServiceMetricsHandler)
}

// @Summary 监控总览
// @Description 监控总览
// @Tags Monitor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response{data=entity.MonitorTotal} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/total [get]
func MonitorTotalHandler(w http.ResponseWriter, r *http.Request) {
	data := entity.MonitorTotal{
		Total: []entity.MonitorTotalItem{
			monitor.GetMonitorHostTotal(),
			monitor.GetMonitorDbTotal(),
			monitor.GetMonitorServiceTotal(),
			monitor.GetMonitorGatewayTotal(),
			monitor.GetMonitorInterfacesTotal(),
		},
	}
	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控主机列表
// @Description 监控主机列表
// @Tags Monitor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response{data=[]entity.MonitorListHostItem} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/list/hosts [get]
func MonitorListHostHandler(w http.ResponseWriter, r *http.Request) {
	data := monitor.GetMonitorListHost()

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控数据库列表
// @Description 监控数据库列表
// @Tags Monitor
// @Accept  json
// @Produce  json
// @Param name query string false "过滤名称, 不填则为全部"
// @Success 200 {object} common.Response{data=[]entity.MonitorListDbItem} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/list/dbs [get]
func MonitorListDbHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data := monitor.GetMonitorListDb()
	if name != "" {
		filterData := make([]*entity.MonitorListDbItem, 0)
		for _, item := range data {
			if item.Name == name {
				filterData = append(filterData, item)
			}
		}
		data = filterData
	}

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控服务列表
// @Description 监控服务列表
// @Tags Monitor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response{data=[]entity.MonitorListServiceItem} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/list/services [get]
func MonitorListServiceHandler(w http.ResponseWriter, r *http.Request) {
	data := monitor.GetMonitorListService()

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控采集网关列表
// @Description 监控采集网关列表
// @Tags Monitor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response{data=[]entity.MonitorListGatewayItem} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/list/gateways [get]
func MonitorListGatewayHandler(w http.ResponseWriter, r *http.Request) {
	data := monitor.GetMonitorListGateway()

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控接口列表
// @Description 监控接口列表
// @Tags Monitor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response{data=[]entity.MonitorListInterfaceItem} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/list/interfaces [get]
func MonitorListInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	data := monitor.GetMonitorListInterfaces()

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控接口-log
// @Description 监控接口-log
// @Tags Monitor
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=entity.MonitorLog} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/log [get]
func MonitorLogHandler(w http.ResponseWriter, r *http.Request) {

	levels := []string{"debug", "info", "warn", "error"}
	data, err := monitor.GetLokiLogStatistics(r.Context())
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	for _, l := range levels {
		_, exist := data[l]
		if !exist {
			data[l] = 0
		}
	}

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控接口-log-detail
// @Description 监控接口-log-detail
// @Tags Monitor
// @Accept json
// @Produce json
// @Param name query string false "名称"
// @Param level query string false "级别， debug,info,error,warn,不填则为全部"
// @Param filter query string false "过滤,可使用正则表达式"
// @Param limit query int false "限制条数，默认100"
// @Param pre_hours query int false "多少小时前，默认24"
// @Success 200 {object} common.Response{data=entity.MonitorLog} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/log/detail [get]
func MonitorLogDetailHandler(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get("level")
	levelQ := ".+"
	if level != "" {
		levelQ = level
	}
	filter := r.URL.Query().Get("filter")
	limit := r.URL.Query().Get("limit")
	limitI := 100
	if limit != "" {
		limitI, _ = strconv.Atoi(limit)
	}
	preHoursI := 24

	if preHoursIStr := r.URL.Query().Get("pre_hours"); preHoursIStr != "" {
		preHoursI, _ = strconv.Atoi(preHoursIStr)
	}
	name := r.URL.Query().Get("name")

	data, err := monitor.GetLokiLogDetail(r.Context(), name, levelQ, filter, limitI, preHoursI)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控接口-log-label values
// @Description 监控接口-label values
// @Tags Monitor
// @Accept json
// @Produce json
// @Param label query string true "label"
// @Success 200 {object} common.Response{data=[]string} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/log/label/values [get]
func MonitorLogLabelValuesHandler(w http.ResponseWriter, r *http.Request) {
	label := r.URL.Query().Get("label")
	if label == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("label is blank"))
		return
	}
	data, err := monitor.GetLokiLabelVales(r.Context(), label)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控接口-gateway-detail
// @Description 监控接口-gateway-detail
// @Tags Monitor
// @Accept json
// @Produce json
// @Param identifier query string true "identifier"
// @Success 200 {object} common.Response{data=entity.MonitorGatewayDetail} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/gateway/detail [get]
func MonitorGatewayDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("identifier")
	if id == "" {
		common.HttpResult(w, common.ErrService.AppendMsg("identifier is blank"))
		return
	}
	data, err := monitor.GetMonitorGatewayDetail(r.Context(), id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK.WithData(data))
}

// @Summary 监控接口-service-metrics
// @Description 监控接口-获取服务监控指标集
// @Tags Monitor
// @Accept json
// @Produce json
// @Param host query string true "host"
// @Param name query string true "name"
// @Success 200 {object} common.Response{data=[]common.MetricPanel} "成功code和成功信息"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /monitor/service/metrics_panel [get]
func MonitorServiceMetricsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		common.HttpResult(w, common.ErrService.AppendMsg("name is blank"))
		return
	}
	host := r.URL.Query().Get("host")
	data, err := common.GetDaprClient().InvokeMethod(r.Context(), name, "/metrics_panel", "GET")
	result := make([]common.MetricPanel, 0)
	if err != nil {
		//common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		common.Logger.Error(err.Error())
		result = common.DefaultmetricsPanel
		//return
	} else {
		err = json.Unmarshal(data, &result)
		if err != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
			return
		}
	}
	cResult := make([]common.MetricPanel, 0)
	for _, p := range result {
		p.Query = strings.ReplaceAll(p.Query, "${HOST}", host)
		p.Query = strings.ReplaceAll(p.Query, "${NAME}", name)
		cResult = append(cResult, p)
	}

	common.HttpResult(w, common.OK.WithData(cResult))
}
