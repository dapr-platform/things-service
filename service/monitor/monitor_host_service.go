package monitor

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/spf13/cast"
	"things-service/entity"
	"time"
)

//monitor service， 提供设备监控数据查询接口, 通过 victoriametrics 的http api接口，定期查询 cpu_usage_active等指标，作为host存活的依据，并
//将指标缓存到内存中。(使用expiredmap,超时15秒）

var cacheHosts = make(map[string]*entity.MonitorListHostItem)

func refreshHost() {
	for {
		select {
		case <-time.After(time.Second * 15):
			// 15秒刷新一次
			err := refreshHostMetrics()
			if err != nil {
				common.Logger.Error("refreshHostMetrics err", err)
			}
		}

	}
}
func refreshHostMetrics() (err error) {
	query := "cpu_usage_active"
	vector, err := fetchOneMetricVal(context.Background(), query, time.Now())
	if err != nil {
		common.Logger.Error("refreshHostMetrics cpu_usage_active err", err)
		return
	}
	for _, v := range *vector {
		//common.Logger.Info("refreshHostMetrics cpu ", v.Metric, v.Value)
		host := cast.ToString(string(v.Metric["ident"]))
		ip := cast.ToString(string(v.Metric["ip"]))
		os := cast.ToString(string(v.Metric["os"]))
		kernel := cast.ToString(string(v.Metric["kernel"]))
		item, exist := cacheHosts[host]
		if !exist {
			item = &entity.MonitorListHostItem{
				Name:     host,
				Ip:       ip,
				Os:       os,
				Kernel:   kernel,
				CpuUsage: toFloat64With2bits(float64(v.Value) / 100),
			}

		} else {
			item.CpuUsage = toFloat64With2bits(float64(v.Value) / 100)
		}
		if v.Timestamp.Unix() < time.Now().Add(-time.Minute*1).Unix() {
			item.Status = 0
		} else {
			item.Status = 1
		}
		cacheHosts[host] = item
	}
	query = "mem_used_percent"
	vector, err = fetchOneMetricVal(context.Background(), query, time.Now())
	if err != nil {
		common.Logger.Error("refreshHostMetrics mem_used_percent err", err)
	}
	for _, v := range *vector {
		//common.Logger.Info("refreshHostMetrics mem_used_percent ", v.Metric, v.Value)
		host := cast.ToString(string(v.Metric["ident"]))
		ip := cast.ToString(string(v.Metric["ip"]))
		item, exist := cacheHosts[host]
		if !exist {
			item = &entity.MonitorListHostItem{
				Name:     host,
				Ip:       ip,
				MemUsage: toFloat64With2bits(float64(v.Value) / 100),
			}
		} else {
			item.MemUsage = toFloat64With2bits(float64(v.Value) / 100)
		}
		if v.Timestamp.Unix() < time.Now().Add(-time.Minute*1).Unix() {
			item.Status = 0
		} else {
			item.Status = 1
		}
		cacheHosts[host] = item
	}

	return
}
func GetMonitorListHost() (result []*entity.MonitorListHostItem) {
	for _, v := range cacheHosts {
		result = append(result, v)
	}
	return
}
func GetMonitorHostTotal() (result entity.MonitorTotalItem) {
	result.Name = "主机"
	for _, v := range cacheHosts {
		result.Total++

		if v.CpuUsage > 0.9 || v.MemUsage > 0.9 {
			result.Error++
		} else {
			result.Normal++
		}
	}
	return
}
