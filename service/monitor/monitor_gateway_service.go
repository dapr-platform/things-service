package monitor

import (
	"context"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"strings"
	
	"things-service/entity"
	"things-service/model"
	"time"
)

//monitor service， 提供设备监控数据查询接口, 通过 victoriametrics 的http api接口，定期查询 cpu_usage_active等指标，作为host存活的依据，并
//将指标缓存到内存中。(使用expiredmap,超时15秒）

var cacheGateways = make(map[string]*entity.MonitorListGatewayItem)

func refreshGateways() {
	for {
		select {
		case <-time.After(time.Second * 15):
			// 15秒刷新一次
			err := refreshGatewayMetrics()
			if err != nil {
				common.Logger.Error("refreshGatewayMetrics err", err)
			}
		}

	}
}
func refreshGatewayMetrics() (err error) {
	qstr := "type=2"
	gateways, err := common.DbQuery[entity.DeviceInfo](context.Background(), common.GetDaprClient(), model.Device_with_tagTableInfo.Name, qstr)
	if err != nil {
		common.Logger.Error("refreshGatewayMetrics err", err)
	}
	for _, v := range gateways {
		key := v.Identifier
		if _, ok := cacheGateways[key]; !ok {
			cacheGateways[key] = &entity.MonitorListGatewayItem{}
		}
		cacheGateways[key].Status = v.Status
		cacheGateways[key].Identifier = v.Identifier
		cacheGateways[key].Name = v.Name
		aiCount := 0
		diCount := 0
		aoCount := 0
		doCount := 0
		for _, t := range v.Tags {
			tarr := strings.Split(t, ":")
			if len(tarr) == 2 {
				if tarr[0] == "AI数量" {
					if tarr[1] != "" {
						aiCount = cast.ToInt(tarr[1])
					}
				} else if tarr[0] == "DI数量" {
					if tarr[1] != "" {
						diCount = cast.ToInt(tarr[1])
					}
				} else if tarr[0] == "AO数量" {
					if tarr[1] != "" {
						aoCount = cast.ToInt(tarr[1])
					}
				} else if tarr[0] == "DO数量" {
					if tarr[1] != "" {
						doCount = cast.ToInt(tarr[1])
					}
				}
			}

		}
		cacheGateways[key].AIPortCount = aiCount
		cacheGateways[key].DIPortCount = diCount
		cacheGateways[key].AOPortCount = aoCount
		cacheGateways[key].DOPortCount = doCount

	}

	return
}

func GetMonitorListGateway() (result []*entity.MonitorListGatewayItem) {
	for _, v := range cacheGateways {
		result = append(result, v)
	}
	return
}
func GetMonitorGatewayTotal() (result entity.MonitorTotalItem) {
	result.Name = "网关"
	for _, v := range cacheGateways {
		result.Total++
		if v.Status == 1 {
			result.Normal++
		} else {
			result.Error++
		}
	}
	return
}

func GetMonitorGatewayDetail(ctx context.Context, identifier string) (result *entity.MonitorGatewayDetail, err error) {
	qstr := "identifier=" + identifier
	gateway, err := common.DbGetOne[entity.DeviceInfo](ctx, common.GetDaprClient(), model.Device_with_tagTableInfo.Name, qstr)
	if err != nil {
		err = errors.Wrap(err, "query device info error")
		return
	}
	result = &entity.MonitorGatewayDetail{
		Name:       gateway.Name,
		Identifier: identifier,
		Tags:       gateway.Tags,
		Points:     make([]entity.MonitorGatewayDetailPoint, 0),
	}
	qstr = "gateway_identifier=" + identifier
	datas, err := common.DbQuery[entity.PointInfo](ctx, common.GetDaprClient(), model.Point_infoTableInfo.Name, qstr)
	if err != nil {
		err = errors.Wrap(err, "query point info error")
	}
	for _, d := range datas {
		tags := d.Tags
		accessType := ""
		Address485 := ""
		AddressIO := ""
		for _, t := range tags {
			tarr := strings.Split(t, ":")
			if len(tarr) == 2 {
				if tarr[0] == "接入类型" {
					accessType = tarr[1]
				} else if tarr[0] == "485地址" {
					Address485 = tarr[1]
				} else if tarr[0] == "端子" {
					AddressIO = tarr[1]
				}
			}
		}
		status := 0
		if d.Ts.Valid {
			if d.Ts.Time.Unix() > time.Now().Unix()-POINT_DATA_OFFLINE_CHECK_SECONDS {
				status = 1
				result.Status = 1
			}
		}

		result.Points = append(result.Points, entity.MonitorGatewayDetailPoint{
			Name:             d.Name,
			DeviceIdentifier: d.DeviceIdentifier,
			AccessType:       accessType,
			Address485:       Address485,
			AddressIO:        AddressIO,
			Status:           status,
			UpdateTime:       d.Ts.String(),
			Value:            cast.ToString(d.Value),
		})

	}
	return
}
