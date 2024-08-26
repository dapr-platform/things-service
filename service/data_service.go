package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"strings"

	"things-service/entity"
)

/*
func init() {
	p, _ := ants.NewPool(500)
	go service.MqttListen(config.MQTT_BROKER, config.MQTT_USER, config.MQTT_PASSWORD, config.MQTT_CLIENT_ID, func(c mqtt.Client, m mqtt.Message) {
		_ = p.Submit(func() { //deviceProcess
			MqttProcessDataDevice(m.Topic(), m.Payload())
		})
	})
}

*/

// TODO use redis cache
func MqttProcessDataDevice(topic string, payload []byte) {
	var deviceMsg entity.DeviceInfoMsg
	err := json.Unmarshal(payload, &deviceMsg)
	if err != nil {
		common.Logger.Error("MqttProcessDataDevice Unmarshal " + err.Error())
		return
	}
	err = ProcessDeviceMsg(context.Background(), deviceMsg)
	if err != nil {
		common.Logger.Error("MqttProcessDataDevice ProcessDeviceMsg " + err.Error())
	}
	return
}

func GetHistoryData(ctx context.Context, ids, startTime, endTime string) (data []entity.DeviceHistoryData, err error) {
	selectStrFmt := `d.device_id as id,d.device_name as name,
        json_agg(
            json_build_object(
                'id', d.point_id,
                'name', d.point_name,
                'datas', (SELECT 
                                json_agg(
                                    json_build_object(
                                        'ts', f.ts,
                                        'value', f.value
                                    ) ORDER BY f.ts
                                ) 
                            FROM 
                                f_point_data f 
                            WHERE 
                                f.id = d.point_id 
                                AND f.ts >= '%s' 
                                AND f.ts <= '%s'
            )
        )
			) as points
`
	selectStr := fmt.Sprintf(selectStrFmt, startTime, endTime)
	idSlice := strings.Split(ids, ",")
	for i, id := range idSlice {
		idSlice[i] = fmt.Sprintf("'%s'", id)
	}
	newIds := strings.Join(idSlice, ",")
	whereStr := `1=1 GROUP BY 
    d.device_id,
    d.device_name`
	fromStrFmt := `(
        SELECT 
            d.id AS device_id,
            d.name AS device_name,
            p.id AS point_id,
            p.name AS point_name
        FROM 
            o_device d
            INNER JOIN o_point p ON d.id = p.device_id
        WHERE 
            d.id in (%s)
    ) AS d
`
	fromStr := fmt.Sprintf(fromStrFmt, newIds)
	data, err = common.CustomSql[entity.DeviceHistoryData](ctx, common.GetDaprClient(), selectStr, fromStr, whereStr)
	if err != nil {
		err = errors.Wrap(err, "customSql query history error . select "+selectStr+" from "+fromStr+" where "+whereStr)
	}
	return
}
