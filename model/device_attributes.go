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


Table: o_device_attributes
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] updated_by                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] updated_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] device_id                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] attribute_name                                 VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] tag                                            VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] value                                          FLOAT8               null: false  primary: false  isArray: false  auto: false  col: FLOAT8          len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "LwauxiKOXeFjVUXsoTFBqtZSc",    "updated_by": "JVSYmBbxHTshkQiBSRQLBeCKf",    "updated_time": 79,    "device_id": "ascrlySYkkKqGunbasnDhYMDp",    "attribute_name": "HEakIQfDKfZTZoNAsrosKGHXg",    "tag": "xvCuiCciirCasrpSdHXLulueq",    "value": 0.52195853}



*/

var (
	Device_attributes_FIELD_NAME_id = "id"

	Device_attributes_FIELD_NAME_updated_by = "updated_by"

	Device_attributes_FIELD_NAME_updated_time = "updated_time"

	Device_attributes_FIELD_NAME_device_id = "device_id"

	Device_attributes_FIELD_NAME_attribute_name = "attribute_name"

	Device_attributes_FIELD_NAME_tag = "tag"

	Device_attributes_FIELD_NAME_value = "value"
)

// Device_attributes struct is a row record of the o_device_attributes table in the thingsdb database
type Device_attributes struct {
	ID            string           `json:"id"`             //唯一标识(device_id+attribute_name)
	UpdatedBy     string           `json:"updated_by"`     //更新人
	UpdatedTime   common.LocalTime `json:"updated_time"`   //更新时间
	DeviceID      string           `json:"device_id"`      //设备唯一标识
	AttributeName string           `json:"attribute_name"` //属性英文名（meta中对应)
	Tag           string           `json:"tag"`            //tag（属性对应的tag）
	Value         float32          `json:"value"`          //当前值

}

var Device_attributesTableInfo = &TableInfo{
	Name: "o_device_attributes",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `唯一标识(device_id+attribute_name)`,
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
			Name:               "updated_by",
			Comment:            `更新人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "UpdatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "updated_by",
			ProtobufFieldName:  "updated_by",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "updated_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "device_id",
			Comment:            `设备唯一标识`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "attribute_name",
			Comment:            `属性英文名（meta中对应)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "AttributeName",
			GoFieldType:        "string",
			JSONFieldName:      "attribute_name",
			ProtobufFieldName:  "attribute_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "tag",
			Comment:            `tag（属性对应的tag）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Tag",
			GoFieldType:        "string",
			JSONFieldName:      "tag",
			ProtobufFieldName:  "tag",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "value",
			Comment:            `当前值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "FLOAT8",
			DatabaseTypePretty: "FLOAT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "FLOAT8",
			ColumnLength:       -1,
			GoFieldName:        "Value",
			GoFieldType:        "float32",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Device_attributes) TableName() string {
	return "o_device_attributes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Device_attributes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Device_attributes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Device_attributes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Device_attributes) TableInfo() *TableInfo {
	return Device_attributesTableInfo
}
