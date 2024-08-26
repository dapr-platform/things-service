package monitor

import (
	"context"
	"fmt"
	"github.com/dapr-platform/common"
	"github.com/prometheus/common/model"
	"strconv"
	"things-service/config"
	"things-service/monitor_client"
	"time"
)

//monitor service， 提供设备监控数据查询接口, 通过 victoriametrics 的http api接口，定期查询 cpu_usage_active等指标，作为host存活的依据，并
//将指标缓存到内存中。(使用expiredmap,超时15秒）

func init() {
	if config.MONITOR_ENABLED {
		go refreshHost()
		go refreshDB()
		go refreshServices()
		go refreshGateways()
		go refreshInterfaces()
	}
}

var POINT_DATA_OFFLINE_CHECK_SECONDS = int64(120)

func toFloat64With2bits(val float64) (v float64) {
	v, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", val), 64)
	return
}

func fetchOneMetricVal(ctx context.Context, query string, queryTime time.Time) (result *model.Vector, err error) {
	data, err := monitor_client.Query(ctx, query, queryTime)
	if err != nil {
		common.Logger.Error("refreshHostMetrics err", err)
		return
	}

	switch data.Result.Type().String() {
	case "vector":
		vector := data.Result.(model.Vector)
		result = &vector
	}
	return
}
