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


Table: f_point_data
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] ts                                             TIMESTAMP            null: false  primary: true   isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 2] key                                            VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] value                                          FLOAT8               null: true   primary: false  isArray: false  auto: false  col: FLOAT8          len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "dmmWoKSrpFawDLcLqCnudjpWL",    "ts": 39,    "key": "OEqDsjLbBEZlQapISJXkouoMv",    "value": 0.958242}



*/

var (
	Point_data_FIELD_NAME_id = "id"

	Point_data_FIELD_NAME_ts = "ts"

	Point_data_FIELD_NAME_key = "key"

	Point_data_FIELD_NAME_value = "value"
)

// Point_data struct is a row record of the f_point_data table in the thingsdb database
type Point_data struct {
	ID    string           `json:"id"`    //点位id
	Ts    common.LocalTime `json:"ts"`    //创建时间
	Key   string           `json:"key"`   //key
	Value float64          `json:"value"` //值

}

var Point_dataTableInfo = &TableInfo{
	Name: "f_point_data",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `点位id`,
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
			Name:               "ts",
			Comment:            `创建时间`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "key",
			Comment:            `key`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "Key",
			GoFieldType:        "string",
			JSONFieldName:      "key",
			ProtobufFieldName:  "key",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "value",
			Comment:            `值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "FLOAT8",
			DatabaseTypePretty: "FLOAT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "FLOAT8",
			ColumnLength:       -1,
			GoFieldName:        "Value",
			GoFieldType:        "float64",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Point_data) TableName() string {
	return "f_point_data"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Point_data) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Point_data) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Point_data) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Point_data) TableInfo() *TableInfo {
	return Point_dataTableInfo
}
