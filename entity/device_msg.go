package entity

type DeviceInfoMsg struct {
	Ts         int64          `json:"ts"`
	Identifier string         `json:"identifier"`
	Properties map[string]any `json:"properties"`
}
