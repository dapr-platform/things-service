package monitor

import (
	"context"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"github.com/spf13/cast"
	"math"
	"strings"
	"things-service/monitor_client"
)

//monitor service， 提供设备监控数据查询接口, 通过 victoriametrics 的http api接口，定期查询 cpu_usage_active等指标，作为host存活的依据，并
//将指标缓存到内存中。(使用expiredmap,超时15秒）

func GetLokiLabelVales(ctx context.Context, label string) (ret []string, err error) {
	return monitor_client.LokiLabelValues(ctx, label)
}

func GetLokiLogDetail(ctx context.Context, name, levelQ, filter string, limit int, preHours int) (ret []string, err error) {
	qstr := "{level=~\"" + levelQ + "\""
	if name != "" {
		if strings.Index(name, "_") > 0 {
			arr := strings.Split(name, "_")
			name = arr[1]
			qstr += ",hostname=\"" + arr[0] + "\""
		}
		qstr += ",compose_service=\"" + name + "\""
	}
	qstr += "}"
	if filter != "" {
		qstr += "|~\"" + filter + "\""
	}

	result, err := monitor_client.LokiStreamQuery(ctx, qstr, limit, preHours)
	if err != nil {
		err = errors.Wrap(err, "GetLokiLogDetail")
		return
	}
	ret = make([]string, 0)
	for _, s := range result.Result {
		for _, v := range s.Values {
			ret = append(ret, cast.ToString(v[1]))
		}
	}
	return
}

func GetLokiLogStatistics(ctx context.Context) (ret map[string]int, err error) {
	qstr := "sum(count_over_time({level=~\".+\"}[24h])) by (level)"
	result, err := monitor_client.LokiQuery(ctx, qstr, 1)
	if err != nil {
		err = errors.Wrap(err, "GetLokiLogStatistics")
		return
	}
	ret = make(map[string]int)
	//common.Logger.Debug("result type =" + result.Result.Type().String())
	switch result.Result.Type().String() {
	case "matrix":
		matrix := result.Result.(model.Matrix)
		for _, sample := range matrix {
			levelStr := sample.Metric["level"]
			levelVal := math.MaxInt
			for _, value := range sample.Values {
				//value[0] is timestamp
				//value[1] is value
				ivalue := cast.ToInt(value.Value.String())
				if ivalue > levelVal {
					levelVal = ivalue
				}
			}
			ret[cast.ToString(levelStr)] = levelVal
		}
	case "vector":
		vector := result.Result.(model.Vector)
		//common.Logger.Debug("vector = ", vector)
		for _, sample := range vector {
			levelStr := string(sample.Metric["level"])
			levelVal := cast.ToInt(float64(sample.Value))
			ret[levelStr] = levelVal
		}
	}
	return
}
