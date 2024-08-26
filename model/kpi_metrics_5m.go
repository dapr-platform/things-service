package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"github.com/guregu/null"
	"github.com/satori/go.uuid"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: f_kpi_metrics_5m
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] device_id                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] ts                                             TIMESTAMP            null: false  primary: true   isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] kpi_id                                         VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] value                                          FLOAT4               null: false  primary: false  isArray: false  auto: false  col: FLOAT4          len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "WItXvHPPcPKZqggOfLOObPZBw",    "device_id": "xjXTNxidrvuUZmoNtuUdbBVVb",    "ts": 94,    "kpi_id": "mqMCnItyFEtcCovopAmsHkFjA",    "value": 0.4237050340459372}



*/

var (
	Kpi_metrics_5m_FIELD_NAME_id = "id"

	Kpi_metrics_5m_FIELD_NAME_device_id = "device_id"

	Kpi_metrics_5m_FIELD_NAME_ts = "ts"

	Kpi_metrics_5m_FIELD_NAME_kpi_id = "kpi_id"

	Kpi_metrics_5m_FIELD_NAME_value = "value"
)

// Kpi_metrics_5m struct is a row record of the f_kpi_metrics_5m table in the thingsdb database
type Kpi_metrics_5m struct {
	ID       string           `json:"id"`        //唯一标识
	DeviceID string           `json:"device_id"` //设备id
	Ts       common.LocalTime `json:"ts"`        //时间戳
	KpiID    string           `json:"kpi_id"`    //kpi_id
	Value    float64          `json:"value"`     //值

}

var Kpi_metrics_5mTableInfo = &TableInfo{
	Name: "f_kpi_metrics_5m",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `唯一标识`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "device_id",
			Comment:            `设备id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DeviceID",
			GoFieldType:        "string",
			JSONFieldName:      "device_id",
			ProtobufFieldName:  "device_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ts",
			Comment:            `时间戳`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "Ts",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "ts",
			ProtobufFieldName:  "ts",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "kpi_id",
			Comment:            `kpi_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "KpiID",
			GoFieldType:        "string",
			JSONFieldName:      "kpi_id",
			ProtobufFieldName:  "kpi_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "value",
			Comment:            `值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "FLOAT4",
			DatabaseTypePretty: "FLOAT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "FLOAT4",
			ColumnLength:       -1,
			GoFieldName:        "Value",
			GoFieldType:        "float64",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (k *Kpi_metrics_5m) TableName() string {
	return "f_kpi_metrics_5m"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (k *Kpi_metrics_5m) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (k *Kpi_metrics_5m) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (k *Kpi_metrics_5m) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (k *Kpi_metrics_5m) TableInfo() *TableInfo {
	return Kpi_metrics_5mTableInfo
}
