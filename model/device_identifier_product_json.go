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


Table: v_device_identifier_product_json
[ 0] identifier                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] json_data                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "identifier": "YHjxoITgycUjxvgcXRFIeSDID",    "status": 47,    "json_data": "uYNLrKPZTOlxyZFWkXvFcVJfW"}


Comments
-------------------------------------
[ 0] Warning table: v_device_identifier_product_json does not have a primary key defined, setting col position 1 identifier as primary key
Warning table: v_device_identifier_product_json primary key column identifier is nullable column, setting it as NOT NULL




*/

var (
	Device_identifier_product_json_FIELD_NAME_identifier = "identifier"

	Device_identifier_product_json_FIELD_NAME_status = "status"

	Device_identifier_product_json_FIELD_NAME_json_data = "json_data"
)

// Device_identifier_product_json struct is a row record of the v_device_identifier_product_json table in the  database
type Device_identifier_product_json struct {
	Identifier string `json:"identifier"` //identifier
	Status     int32  `json:"status"`     //status
	JSONData   string `json:"json_data"`  //json_data

}

var Device_identifier_product_jsonTableInfo = &TableInfo{
	Name: "v_device_identifier_product_json",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "identifier",
			Comment: `identifier`,
			Notes: `Warning table: v_device_identifier_product_json does not have a primary key defined, setting col position 1 identifier as primary key
Warning table: v_device_identifier_product_json primary key column identifier is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Identifier",
			GoFieldType:        "string",
			JSONFieldName:      "identifier",
			ProtobufFieldName:  "identifier",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "status",
			Comment:            `status`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "json_data",
			Comment:            `json_data`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "JSONData",
			GoFieldType:        "string",
			JSONFieldName:      "json_data",
			ProtobufFieldName:  "json_data",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Device_identifier_product_json) TableName() string {
	return "v_device_identifier_product_json"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Device_identifier_product_json) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Device_identifier_product_json) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Device_identifier_product_json) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Device_identifier_product_json) TableInfo() *TableInfo {
	return Device_identifier_product_jsonTableInfo
}
