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


Table: v_device_current_data
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] tags                                           _TEXT                null: true   primary: false  isArray: false  auto: false  col: _TEXT           len: -1      default: []
[ 3] points                                         JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "EpQkGbNQgwRyJatmoRPHCGrZq",    "name": "lPNatvZKJtAHHuFWSERcdeNUd",    "tags": "HMKvidowSdQXJioeKbuYrgbDQ",    "points": 99}


Comments
-------------------------------------
[ 0] Warning table: v_device_current_data does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_device_current_data primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Device_current_data_FIELD_NAME_id = "id"

	Device_current_data_FIELD_NAME_name = "name"

	Device_current_data_FIELD_NAME_tags = "tags"

	Device_current_data_FIELD_NAME_points = "points"
)

// Device_current_data struct is a row record of the v_device_current_data table in the  database
type Device_current_data struct {
	ID     string `json:"id"`     //id
	Name   string `json:"name"`   //name
	Tags   string `json:"tags"`   //tags
	Points any    `json:"points"` //points

}

var Device_current_dataTableInfo = &TableInfo{
	Name: "v_device_current_data",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_device_current_data does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_device_current_data primary key column id is nullable column, setting it as NOT NULL
`,
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
			Name:               "name",
			Comment:            `name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "tags",
			Comment:            `tags`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "_TEXT",
			DatabaseTypePretty: "_TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            true,
			ColumnType:         "_TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Tags",
			GoFieldType:        "string",
			JSONFieldName:      "tags",
			ProtobufFieldName:  "tags",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "points",
			Comment:            `points`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "Points",
			GoFieldType:        "any",
			JSONFieldName:      "points",
			ProtobufFieldName:  "points",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Device_current_data) TableName() string {
	return "v_device_current_data"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Device_current_data) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Device_current_data) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Device_current_data) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Device_current_data) TableInfo() *TableInfo {
	return Device_current_dataTableInfo
}
