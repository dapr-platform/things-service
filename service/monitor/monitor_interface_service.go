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

var cacheInterfaces = make(map[string]*entity.MonitorListInterfaceItem)

func refreshInterfaces() {
	for {
		select {
		case <-time.After(time.Second * 15):
			// 15秒刷新一次
			err := refreshInterfacesMetrics()
			if err != nil {
				common.Logger.Error("refreshInterfaces err", err)
			}
		}

	}
}
func refreshInterfacesMetrics() (err error) {
	query := "iot_service_live"
	err = fillOneInterfaceMetric(query, "status")
	if err != nil {
		common.Logger.Error("refreshInterfaces "+query+" err", err)
	}
	query = "iot_service_alarm_total"
	err = fillOneInterfaceMetric(query, "alarm_total")
	if err != nil {
		common.Logger.Error("refreshInterfaces "+query+" err", err)
	}

	return
}

func fillOneInterfaceMetric(query string, fieldName string) (err error) {
	vector, err := fetchOneMetricVal(context.Background(), query, time.Now())
	if err != nil {
		common.Logger.Error("refreshInterfaces "+query+" err", err)
		return
	}
	for _, v := range *vector {
		//common.Logger.Info("refreshInterfaces "+query+" ", v.Metric, v.Value)
		instance := cast.ToString(string(v.Metric["instance"]))
		instance = instance[:strings.Index(instance, ":")]
		ident := cast.ToString((string)(v.Metric["ident"]))
		key := ident + "_" + instance
		service_type := cast.ToString((string)(v.Metric["service_type"]))
		thirdSystem := cast.ToString((string)(v.Metric["third_system"]))
		protocol := cast.ToString((string)(v.Metric["protocol"]))
		if service_type != "third" {
			continue
		}

		item, exist := cacheInterfaces[key]
		if !exist {
			item = &entity.MonitorListInterfaceItem{
				Name:             key,
				ThirdPartySystem: thirdSystem,
				Protocol:         protocol,
				LatestDataTime:   v.Timestamp.Time().Format("2006-01-02 15:04:05"),
			}

		}

		switch fieldName {
		case "status":
			item.Status = int(v.Value)
		case "alarm_total":
			item.AlertCount = int(v.Value)
		default:
			common.Logger.Error("refreshInterfaces unknown field ", fieldName)
			err = errors.New("unknown field ")
			return
		}

		cacheInterfaces[key] = item
	}
	return
}

func GetMonitorListInterfaces() (result []*entity.MonitorListInterfaceItem) {
	for _, v := range cacheInterfaces {
		result = append(result, v)
	}
	return
}
func GetMonitorInterfacesTotal() (result entity.MonitorTotalItem) {
	result.Name = "第三方接口"
	for _, v := range cacheInterfaces {
		result.Total++
		if v.Status == 1 {
			result.Normal++
		} else {
			result.Error++
		}
	}
	return
}
