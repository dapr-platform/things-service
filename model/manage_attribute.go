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


Table: p_manage_attribute
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] meta_id                                        VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] parent_id                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] label                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] value                                          VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "uYxPJEhLVXLAcbjpYrKvaBQeb",    "meta_id": "oHobialnTrJEHZcrHgkgeNuEv",    "parent_id": "SHPnPLXvPindHVTCCUUqJCnjS",    "label": "eJXhtpwpoYGLbjgggkgcEGdof",    "value": "ILuAOQfjuECUlBVvyXqLbYWAZ"}



*/

var (
	Manage_attribute_FIELD_NAME_id = "id"

	Manage_attribute_FIELD_NAME_meta_id = "meta_id"

	Manage_attribute_FIELD_NAME_parent_id = "parent_id"

	Manage_attribute_FIELD_NAME_label = "label"

	Manage_attribute_FIELD_NAME_value = "value"
)

// Manage_attribute struct is a row record of the p_manage_attribute table in the thingsdb database
type Manage_attribute struct {
	ID       string `json:"id"`        //唯一标识
	MetaID   string `json:"meta_id"`   //元数据id
	ParentID string `json:"parent_id"` //父id
	Label    string `json:"label"`     //名称
	Value    string `json:"value"`     //值

}

var Manage_attributeTableInfo = &TableInfo{
	Name: "p_manage_attribute",
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
			Name:               "meta_id",
			Comment:            `元数据id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "MetaID",
			GoFieldType:        "string",
			JSONFieldName:      "meta_id",
			ProtobufFieldName:  "meta_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "parent_id",
			Comment:            `父id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ParentID",
			GoFieldType:        "string",
			JSONFieldName:      "parent_id",
			ProtobufFieldName:  "parent_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "label",
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
			GoFieldName:        "Label",
			GoFieldType:        "string",
			JSONFieldName:      "label",
			ProtobufFieldName:  "label",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "value",
			Comment:            `值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Value",
			GoFieldType:        "string",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Manage_attribute) TableName() string {
	return "p_manage_attribute"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Manage_attribute) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Manage_attribute) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Manage_attribute) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Manage_attribute) TableInfo() *TableInfo {
	return Manage_attributeTableInfo
}
