package monitor

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"strings"
	"things-service/entity"
	"time"
)

//monitor service， 提供设备监控数据查询接口, 通过 victoriametrics 的http api接口，定期查询 cpu_usage_active等指标，作为host存活的依据，并
//将指标缓存到内存中。(使用expiredmap,超时15秒）

var cacheServices = make(map[string]*entity.MonitorListServiceItem)

func refreshServices() {
	for {
		select {
		case <-time.After(time.Second * 15):
			// 15秒刷新一次
			err := refreshServiceMetrics()
			if err != nil {
				common.Logger.Error("refreshServiceMetrics err", err)
			}
		}

	}
}
func refreshServiceMetrics() (err error) {
	query := "go_threads"
	err = fillOneServiceMetric(query, "go_threads")
	if err != nil {
		common.Logger.Error("refreshServiceMetrics "+query+" err", err)
	}
	query = "go_memstats_alloc_bytes"
	err = fillOneServiceMetric(query, "memory_usage")
	if err != nil {
		common.Logger.Error("refreshServiceMetrics "+query+" err", err)
	}

	return
}

func fillOneServiceMetric(query string, fieldName string) (err error) {
	vector, err := fetchOneMetricVal(context.Background(), query, time.Now())
	if err != nil {
		common.Logger.Error("refreshServiceMetrics "+query+" err", err)
		return
	}
	for _, v := range *vector {
		//common.Logger.Info("refreshServiceMetrics "+query+" ", v.Metric, v.Value)
		instance := cast.ToString(string(v.Metric["instance"]))
		instance = instance[:strings.Index(instance, ":")]
		ident := cast.ToString((string)(v.Metric["ident"]))
		key := ident + "_" + instance

		item, exist := cacheServices[key]
		if !exist {
			item = &entity.MonitorListServiceItem{
				Name: instance,
			}

		}
		if ident == "" {
			item.Host = "center"
		} else {
			item.Host = ident
		}
		switch fieldName {
		case "go_threads":
			item.ThreadCount = int(v.Value)
		case "memory_usage":
			item.MemoryUsage = int(v.Value)
		default:
			common.Logger.Error("refreshServiceMetrics unknown field ", fieldName)
			err = errors.New("unknown field ")
			return
		}

		if v.Timestamp.Unix() < time.Now().Add(-time.Minute*1).Unix() {
			item.Status = 0
		} else {
			item.Status = 1
		}
		cacheServices[key] = item
	}
	return
}

func GetMonitorListService() (result []*entity.MonitorListServiceItem) {
	for _, v := range cacheServices {
		result = append(result, v)
	}
	return
}
func GetMonitorServiceTotal() (result entity.MonitorTotalItem) {
	result.Name = "服务"
	for _, v := range cacheServices {
		result.Total++
		if v.Status == 1 {
			result.Normal++
		} else {
			result.Error++
		}
	}
	return
}
