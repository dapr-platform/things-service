package monitor

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"things-service/entity"
	"time"
)

//monitor service， 提供设备监控数据查询接口, 通过 victoriametrics 的http api接口，定期查询 cpu_usage_active等指标，作为host存活的依据，并
//将指标缓存到内存中。(使用expiredmap,超时15秒）

var cacheDbs = make(map[string]*entity.MonitorListDbItem)
var ignoreDbNames = []string{"postgres_global", "template0", "template1"}

func refreshDB() {
	for {
		select {
		case <-time.After(time.Second * 15):
			// 15秒刷新一次
			err := refreshDbMetrics()
			if err != nil {
				common.Logger.Error("refreshDbMetrics err", err)
			}
		}

	}
}
func refreshDbMetrics() (err error) {
	query := "postgresql_numbackends"
	err = fillOneDbMetric(query, "connection_count")
	if err != nil {
		common.Logger.Error("refreshDbMetrics "+query+" err", err)
	}
	query = "postgresql_deadlocks"
	err = fillOneDbMetric(query, "deadlock_count")
	if err != nil {
		common.Logger.Error("refreshDbMetrics "+query+" err", err)
	}
	query = "postgresql_conflicts"
	err = fillOneDbMetric(query, "conflict_count")
	if err != nil {
		common.Logger.Error("refreshDbMetrics "+query+" err", err)
	}
	query = "postgresql_blks_hit/(postgresql_blks_hit+postgresql_blks_read)"
	err = fillOneDbMetric(query, "cache_hit_rate")
	if err != nil {
		common.Logger.Error("refreshDbMetrics "+query+" err", err)
	}

	return
}

func fillOneDbMetric(query string, fieldName string) (err error) {
	vector, err := fetchOneMetricVal(context.Background(), query, time.Now())
	if err != nil {
		common.Logger.Error("refreshDbMetrics "+query+" err", err)
		return
	}
	for _, v := range *vector {
		//common.Logger.Info("refreshDbMetrics "+query+" ", v.Metric, v.Value)
		//host := cast.ToString(string(v.Metric["ident"]))
		db := cast.ToString(string(v.Metric["db"]))
		server := cast.ToString(string(v.Metric["server"]))
		key := server + "_" + db
		for _, vv := range ignoreDbNames {
			if db == vv {
				continue
			}
		}
		item, exist := cacheDbs[key]
		if !exist {
			item = &entity.MonitorListDbItem{
				Name:   db,
				Server: server,
			}

		}
		switch fieldName {
		case "deadlock_count":
			item.DeadlockCount = int(v.Value)
		case "cache_hit_rate":
			item.CacheHitRate = toFloat64With2bits(float64(v.Value))
		case "conflict_count":
			item.ConflictCount = int(v.Value)
		case "connection_count":
			item.ConnectionCount = int(v.Value)
		default:
			common.Logger.Error("refreshDbMetrics unknown field ", fieldName)
			err = errors.New("unknown field ")
			return
		}

		if v.Timestamp.Unix() < time.Now().Add(-time.Minute*1).Unix() {
			item.Status = 0
		} else {
			item.Status = 1
		}
		cacheDbs[key] = item
	}
	return
}

func GetMonitorListDb() (result []*entity.MonitorListDbItem) {
	for _, v := range cacheDbs {
		result = append(result, v)
	}
	return
}
func GetMonitorDbTotal() (result entity.MonitorTotalItem) {
	result.Name = "数据库"
	for _, v := range cacheDbs {
		result.Total++
		if v.Status == 1 {
			result.Normal++
		} else {
			result.Error++
		}
	}
	return
}
