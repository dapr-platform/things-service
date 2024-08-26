package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: f_device_data
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] device_identifier                              VARCHAR(128)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 128     default: []
[ 2] property_identifier                            VARCHAR(128)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 128     default: []
[ 3] ts                                             TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] vtype                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] unit                                           VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] f_value                                        FLOAT8               null: true   primary: false  isArray: false  auto: false  col: FLOAT8          len: -1      default: []
[ 7] i_value                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 8] t_value                                        TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 9] s_value                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "DSmWXCkkKiFbOVxktQyXZabEP",    "device_identifier": "jxMBAGdNghtXKXguLfffQPXFH",    "property_identifier": "HxAeRgJZNjhcUYeRvdaGbgxSV",    "ts": 95,    "vtype": 90,    "unit": "spvCsMRuZfYxRvfJgHHHAQYko",    "f_value": 0.4026527,    "i_value": 66,    "t_value": 34,    "s_value": "kHdRtyWUbleOyKACHUcUKIbZZ"}



*/

var (
	Device_data_FIELD_NAME_id = "id"

	Device_data_FIELD_NAME_device_identifier = "device_identifier"

	Device_data_FIELD_NAME_property_identifier = "property_identifier"

	Device_data_FIELD_NAME_ts = "ts"

	Device_data_FIELD_NAME_vtype = "vtype"

	Device_data_FIELD_NAME_unit = "unit"

	Device_data_FIELD_NAME_f_value = "f_value"

	Device_data_FIELD_NAME_i_value = "i_value"

	Device_data_FIELD_NAME_t_value = "t_value"

	Device_data_FIELD_NAME_s_value = "s_value"
)

// Device_data struct is a row record of the f_device_data table in the  database
type Device_data struct {
	ID                 string           `json:"id"`                  //唯一id(device_identifier+property_identifier+ts md5)
	DeviceIdentifier   string           `json:"device_identifier"`   //device_identifier
	PropertyIdentifier string           `json:"property_identifier"` //property_identifier
	Ts                 common.LocalTime `json:"ts"`                  //创建时间
	Vtype              int32            `json:"vtype"`               //值类型1 float,2:int,3:ts,4:string
	Unit               string           `json:"unit"`                //单位
	FValue             float64          `json:"f_value"`             //float值
	IValue             int32            `json:"i_value"`             //int值
	TValue             common.LocalTime `json:"t_value"`             //时间值
	SValue             string           `json:"s_value"`             //string值

}

var Device_dataTableInfo = &TableInfo{
	Name: "f_device_data",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `唯一id(device_identifier+property_identifier+ts md5)`,
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
			Name:               "device_identifier",
			Comment:            `device_identifier`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       128,
			GoFieldName:        "DeviceIdentifier",
			GoFieldType:        "string",
			JSONFieldName:      "device_identifier",
			ProtobufFieldName:  "device_identifier",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "property_identifier",
			Comment:            `property_identifier`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       128,
			GoFieldName:        "PropertyIdentifier",
			GoFieldType:        "string",
			JSONFieldName:      "property_identifier",
			ProtobufFieldName:  "property_identifier",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ts",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "Ts",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "ts",
			ProtobufFieldName:  "ts",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "vtype",
			Comment:            `值类型1 float,2:int,3:ts,4:string`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Vtype",
			GoFieldType:        "int32",
			JSONFieldName:      "vtype",
			ProtobufFieldName:  "vtype",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "unit",
			Comment:            `单位`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "Unit",
			GoFieldType:        "string",
			JSONFieldName:      "unit",
			ProtobufFieldName:  "unit",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "f_value",
			Comment:            `float值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "FLOAT8",
			DatabaseTypePretty: "FLOAT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "FLOAT8",
			ColumnLength:       -1,
			GoFieldName:        "FValue",
			GoFieldType:        "float64",
			JSONFieldName:      "f_value",
			ProtobufFieldName:  "f_value",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "i_value",
			Comment:            `int值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "IValue",
			GoFieldType:        "int32",
			JSONFieldName:      "i_value",
			ProtobufFieldName:  "i_value",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "t_value",
			Comment:            `时间值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "TValue",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "t_value",
			ProtobufFieldName:  "t_value",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "s_value",
			Comment:            `string值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "SValue",
			GoFieldType:        "string",
			JSONFieldName:      "s_value",
			ProtobufFieldName:  "s_value",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Device_data) TableName() string {
	return "f_device_data"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Device_data) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Device_data) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Device_data) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Device_data) TableInfo() *TableInfo {
	return Device_dataTableInfo
}
