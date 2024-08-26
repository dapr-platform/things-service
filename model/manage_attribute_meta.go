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


Table: p_manage_attribute_meta
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] label                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] type                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] unit                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "xEfDFTgMhrjcHkVhZtoZOLtfk",    "name": "hwZJbtkKiJJlyHuQXGyprbIaE",    "label": "miVirGhxuoUgpgRakPjwnufZo",    "type": 33,    "unit": "JmQvBUoQbInGkOQbtgfEMLpES"}



*/

var (
	Manage_attribute_meta_FIELD_NAME_id = "id"

	Manage_attribute_meta_FIELD_NAME_name = "name"

	Manage_attribute_meta_FIELD_NAME_label = "label"

	Manage_attribute_meta_FIELD_NAME_type = "type"

	Manage_attribute_meta_FIELD_NAME_unit = "unit"
)

// Manage_attribute_meta struct is a row record of the p_manage_attribute_meta table in the thingsdb database
type Manage_attribute_meta struct {
	ID    string `json:"id"`    //唯一标识
	Name  string `json:"name"`  //名称
	Label string `json:"label"` //中文名
	Type  int32  `json:"type"`  //字段类型(string:1,integer:2,float:3boolean:4)
	Unit  string `json:"unit"`  //单位

}

var Manage_attribute_metaTableInfo = &TableInfo{
	Name: "p_manage_attribute_meta",
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
			Name:               "name",
			Comment:            `名称`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "label",
			Comment:            `中文名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Label",
			GoFieldType:        "string",
			JSONFieldName:      "label",
			ProtobufFieldName:  "label",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
			Comment:            `字段类型(string:1,integer:2,float:3boolean:4)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "unit",
			Comment:            `单位`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Unit",
			GoFieldType:        "string",
			JSONFieldName:      "unit",
			ProtobufFieldName:  "unit",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Manage_attribute_meta) TableName() string {
	return "p_manage_attribute_meta"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Manage_attribute_meta) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Manage_attribute_meta) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Manage_attribute_meta) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Manage_attribute_meta) TableInfo() *TableInfo {
	return Manage_attribute_metaTableInfo
}
