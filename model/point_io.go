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


Table: o_point_io
[ 0] id                                             VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] gateway_id                                     VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] device_id                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] endpoint                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "tYQxMBmnNFQujIdwmKlQdedfE",    "gateway_id": "VWpxmRUBRrHWjeSdWUYIxuXlq",    "device_id": "apvLRsEwAkXeEQoHXQSnsJDKv",    "name": "ackPVkwUPrVeLJKidFQvBtbXW",    "endpoint": "bihMsWGtXHhaMUMDvuslSrtXE"}


Comments
-------------------------------------
[ 0] Warning table: o_point_io does not have a primary key defined, setting col position 1 id as primary key




*/

var (
	Point_io_FIELD_NAME_id = "id"

	Point_io_FIELD_NAME_gateway_id = "gateway_id"

	Point_io_FIELD_NAME_device_id = "device_id"

	Point_io_FIELD_NAME_name = "name"

	Point_io_FIELD_NAME_endpoint = "endpoint"
)

// Point_io struct is a row record of the o_point_io table in the thingsdb database
type Point_io struct {
	ID        string `json:"id"`         //id
	GatewayID string `json:"gateway_id"` //gateway_id
	DeviceID  string `json:"device_id"`  //device_id
	Name      string `json:"name"`       //name
	Endpoint  string `json:"endpoint"`   //端子

}

var Point_ioTableInfo = &TableInfo{
	Name: "o_point_io",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: o_point_io does not have a primary key defined, setting col position 1 id as primary key
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
			Name:               "gateway_id",
			Comment:            `gateway_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "GatewayID",
			GoFieldType:        "string",
			JSONFieldName:      "gateway_id",
			ProtobufFieldName:  "gateway_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "device_id",
			Comment:            `device_id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "name",
			Comment:            `name`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "endpoint",
			Comment:            `端子`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Endpoint",
			GoFieldType:        "string",
			JSONFieldName:      "endpoint",
			ProtobufFieldName:  "endpoint",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Point_io) TableName() string {
	return "o_point_io"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Point_io) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Point_io) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Point_io) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Point_io) TableInfo() *TableInfo {
	return Point_ioTableInfo
}
