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


Table: o_model_meta
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] identifier                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] json_data                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] meta_type                                      VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
[ 5] category                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] type                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] tag                                            INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "jGBhJmjVETnfOErSygCJBBruS",    "name": "jTNFVZFcumtavtRyXmMvbnWvi",    "identifier": "qNMCBIPgFnAyUmoalIIgnfEKX",    "json_data": "rZqtRNGALqqxxgMZLkoZhVvKi",    "meta_type": "moTdFjcQWlbulLnqeIxGCfuAU",    "category": "cnZymHIlPQZGtNKGxMgTOUmYu",    "type": "lBDMZrZcTcIlyRqyhJeUfDepW",    "tag": 74}



*/

var (
	Model_meta_FIELD_NAME_id = "id"

	Model_meta_FIELD_NAME_name = "name"

	Model_meta_FIELD_NAME_identifier = "identifier"

	Model_meta_FIELD_NAME_json_data = "json_data"

	Model_meta_FIELD_NAME_meta_type = "meta_type"

	Model_meta_FIELD_NAME_category = "category"

	Model_meta_FIELD_NAME_type = "type"

	Model_meta_FIELD_NAME_tag = "tag"
)

// Model_meta struct is a row record of the o_model_meta table in the  database
type Model_meta struct {
	ID         string `json:"id"`         //唯一标识
	Name       string `json:"name"`       //名称
	Identifier string `json:"identifier"` //标识符
	JSONData   string `json:"json_data"`  //元数据
	MetaType   string `json:"meta_type"`  //类型(attribute,service,event)
	Category   string `json:"category"`   //category
	Type       string `json:"type"`       //品类
	Tag        int64  `json:"tag"`        //更新tag

}

var Model_metaTableInfo = &TableInfo{
	Name: "o_model_meta",
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
			Name:               "identifier",
			Comment:            `标识符`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Identifier",
			GoFieldType:        "string",
			JSONFieldName:      "identifier",
			ProtobufFieldName:  "identifier",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "json_data",
			Comment:            `元数据`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "meta_type",
			Comment:            `类型(attribute,service,event)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       20,
			GoFieldName:        "MetaType",
			GoFieldType:        "string",
			JSONFieldName:      "meta_type",
			ProtobufFieldName:  "meta_type",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "category",
			Comment:            `category`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Category",
			GoFieldType:        "string",
			JSONFieldName:      "category",
			ProtobufFieldName:  "category",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "type",
			Comment:            `品类`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "tag",
			Comment:            `更新tag`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "Tag",
			GoFieldType:        "int64",
			JSONFieldName:      "tag",
			ProtobufFieldName:  "tag",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *Model_meta) TableName() string {
	return "o_model_meta"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *Model_meta) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *Model_meta) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *Model_meta) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *Model_meta) TableInfo() *TableInfo {
	return Model_metaTableInfo
}
