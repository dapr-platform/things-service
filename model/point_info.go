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


Table: v_point_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] gateway_id                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] gateway_identifier                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] device_id                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] device_identifier                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] ts                                             TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 7] value                                          FLOAT8               null: true   primary: false  isArray: false  auto: false  col: FLOAT8          len: -1      default: []
[ 8] tags                                           _TEXT                null: true   primary: false  isArray: false  auto: false  col: _TEXT           len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "MRYgpnNuQyKkBBUCtFEFuafWg",    "gateway_id": "yHLTjaTxbNLXfhrmvKoHpePNj",    "gateway_identifier": "gHDVPueHsuxrubUAkuHrXErbm",    "device_id": "rpaPEssJkgqDSBCFRaNHnmbbf",    "device_identifier": "kBprpCrSyamfiXRxBdwEjvSBM",    "name": "qDylrWPjFihSdKCutBFMJBHOn",    "ts": 84,    "value": 0.27541125,    "tags": "CnmpyOBvqRmXZZFHjodqUCYkJ"}


Comments
-------------------------------------
[ 0] Warning table: v_point_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_point_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Point_info_FIELD_NAME_id = "id"

	Point_info_FIELD_NAME_gateway_id = "gateway_id"

	Point_info_FIELD_NAME_gateway_identifier = "gateway_identifier"

	Point_info_FIELD_NAME_device_id = "device_id"

	Point_info_FIELD_NAME_device_identifier = "device_identifier"

	Point_info_FIELD_NAME_name = "name"

	Point_info_FIELD_NAME_ts = "ts"

	Point_info_FIELD_NAME_value = "value"

	Point_info_FIELD_NAME_tags = "tags"
)

// Point_info struct is a row record of the v_point_info table in the  database
type Point_info struct {
	ID                string           `json:"id"`                 //id
	GatewayID         string           `json:"gateway_id"`         //gateway_id
	GatewayIdentifier string           `json:"gateway_identifier"` //gateway_identifier
	DeviceID          string           `json:"device_id"`          //device_id
	DeviceIdentifier  string           `json:"device_identifier"`  //device_identifier
	Name              string           `json:"name"`               //name
	Ts                common.LocalTime `json:"ts"`                 //ts
	Value             float64          `json:"value"`              //value
	Tags              string           `json:"tags"`               //tags

}

var Point_infoTableInfo = &TableInfo{
	Name: "v_point_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_point_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_point_info primary key column id is nullable column, setting it as NOT NULL
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
			Nullable:           true,
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
			Name:               "gateway_identifier",
			Comment:            `gateway_identifier`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "GatewayIdentifier",
			GoFieldType:        "string",
			JSONFieldName:      "gateway_identifier",
			ProtobufFieldName:  "gateway_identifier",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "device_id",
			Comment:            `device_id`,
			Notes:              ``,
			Nullable:           true,
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
			Name:               "device_identifier",
			Comment:            `device_identifier`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DeviceIdentifier",
			GoFieldType:        "string",
			JSONFieldName:      "device_identifier",
			ProtobufFieldName:  "device_identifier",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "ts",
			Comment:            `ts`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "value",
			Comment:            `value`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Point_info) TableName() string {
	return "v_point_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Point_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Point_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Point_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Point_info) TableInfo() *TableInfo {
	return Point_infoTableInfo
}
