package entity

import "github.com/dapr-platform/common"

type PointInfo struct {
	ID                string                   `json:"id"`                 //id
	GatewayID         string                   `json:"gateway_id"`         //gateway_id
	GatewayIdentifier string                   `json:"gateway_identifier"` //gateway_identifier
	DeviceID          string                   `json:"device_id"`          //device_id
	DeviceIdentifier  string                   `json:"device_identifier"`  //device_identifier
	Name              string                   `json:"name"`               //name
	Ts                common.LocalNullableTime `json:"ts"`                 //ts
	Value             float64                  `json:"value"`              //value
	Tags              []string                 `json:"tags"`               //tags
}
