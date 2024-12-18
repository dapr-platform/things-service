package parse

import (
	"github.com/spf13/cast"
	"math"
)

func uint16ArrToFloat32(data []any) (value float32) {
	if len(data) != 2 {
		return
	}
	uuval := make([]uint16, 2)
	uuval[0] = cast.ToUint16(data[0])
	uuval[1] = cast.ToUint16(data[1])
	value = math.Float32frombits(uint32(uuval[0])<<16 + uint32(uuval[1]))
	return
}

// 读转换
func TranslateProperty(data map[string]any) (value map[string]any) {
	value = make(map[string]any)
	for k, v := range data {
		value[k] = v
	}

	return
}

// 写转换
func TranslatePropertySet(data map[string]any, propName string, val any) (value map[string]any) {
	value = make(map[string]any)
	value[propName] = val
	return
}

// 写转换desired
func TranslatePropertySetDesired(data map[string]any, propName string, val any) (value map[string]any) {
	value = make(map[string]any)
	value[propName] = val
	return
}

func CheckAlert(data map[string]any) (alerts []map[string]any) {
	alerts = make([]map[string]any, 0)
	if cast.ToInt(data["水浸状态"]) == 1 {
		event := make(map[string]any)
		event["level"] = 1
		event["alert_property"] = "水浸状态"
		event["alert_value"] = 1
		alerts = append(alerts, event)
	}
	return

}
