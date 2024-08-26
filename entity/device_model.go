package entity

import "github.com/dapr-platform/common"

const MODEL_IDENTIFIER_KEY = "identifier"
const MODEL_NAME_KEY = "name"

type DeviceModelRaw struct {
	Schema     string           `json:"schema"`
	Profile    map[string]any   `json:"profile"`
	Properties []map[string]any `json:"properties"`
	Events     []map[string]any `json:"events"`
	Services   []map[string]any `json:"services"`
}

type ProductModel struct {
	Profile      map[string]any `json:"profile"`
	DeviceModels []DeviceModel  `json:"deviceModels"`
	ParseScript  ParseScript    `json:"parseScript"`
	Tags         []PTag         `json:"tags"`
	Ui           map[string]any `json:"ui"`
}
type PTag struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ParseScript struct {
	Language string `json:"language"`
	Content  string `json:"content"`
}
type DeviceMirror struct {
	State                 DeviceMirrorState     `json:"state"`
	Metadata              DeviceMirrorMetadata  `json:"metadata"`
	Timestamp             int64                 `json:"timestamp"`
	TimestampStr          string                `json:"timestamp_str"`
	LatestRawMsg          string                `json:"latest_raw_msg"`
	Alerts                []map[string]any      `json:"alerts"`
	RecentEvents          []DeviceMirrorEvent   `json:"recent_events"`
	CurrentKpiDatas       []CurrentKpiData      `json:"current_kpi_datas"`
	Recent_InvokeServices []DeviceMirrorService `json:"recent_invoke_services"`
	Version               int64                 `json:"version"`
}
type DeviceMirrorEvent struct {
	Timestamp    int64  `json:"timestamp"`
	TimestampStr string `json:"timestamp_str"`
	Identifier   string `json:"identifier"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Desc         string `json:"desc"`
	Method       string `json:"method"`
	Required     bool   `json:"required"`
	OutputData   any    `json:"output_data"`
}
type DeviceMirrorService struct {
	Timestamp    int64  `json:"timestamp"`
	TimestampStr string `json:"timestamp_str"`
	Identifier   string `json:"identifier"`
	Name         string `json:"name"`
	CallType     string `json:"call_type"`
	Method       string `json:"method"`
	Desc         string `json:"desc"`
	Required     bool   `json:"required"`
	InputData    any    `json:"input_data"`
	OutputData   any    `json:"output_data"`
}
type CurrentKpiData struct {
	Name  string           `json:"name"`
	Id    string           `json:"id"`
	Ts    common.LocalTime `json:"ts"`
	Value any              `json:"value"`
	Unit  string           `json:"unit"`
}

type DeviceMirrorState struct {
	Reported map[string]interface{} `json:"reported"`
	Desired  map[string]interface{} `json:"desired"`
}

type DeviceMirrorMetadata struct {
	Reported map[string]interface{} `json:"reported"`
	Desired  map[string]interface{} `json:"desired"`
}

type DeviceModel struct {
	Schema     string                `json:"schema"`
	Profile    map[string]any        `json:"profile"`
	Properties []DeviceModelProperty `json:"properties"`
	Events     []DeviceModelEvent    `json:"events"`
	Services   []DeviceModelService  `json:"services"`
}

type DeviceModelProperty struct {
	Identifier string         `json:"identifier"`
	Name       string         `json:"name"`
	DataType   DataType       `json:"dataType"`
	Required   bool           `json:"required"`
	AccessMode string         `json:"accessMode"` //r:rw
	Optional   bool           `json:"optional,omitempty"`
	Ui         map[string]any `json:"ui,omitempty"`
}
type DeviceModelPropertyWithData struct {
	Property DeviceModelProperty `json:"property"`
	Data     interface{}         `json:"data"`
}

type DeviceModelEvent struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"` //info/error/alert
	Required   bool   `json:"required"`
	Method     string `json:"method"`
	Desc       string `json:"desc"`
	OutputData []Args `json:"outputData"`
}
type DeviceModelService struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Required   bool   `json:"required"`
	CallType   string `json:"callType"` //async:sync
	Method     string `json:"method"`
	Desc       string `json:"desc"`
	InputData  []Args `json:"inputData"`
	OutputData []Args `json:"outputData"`
}
type Args struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	DataType   DataType `json:"dataType"`
}
type DataType struct {
	Type  string `json:"type"`
	Specs any    `json:"specs"`
}
