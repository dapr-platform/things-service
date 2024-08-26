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


Table: v_tag_with_product_id
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] rel_id                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] key                                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] value                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] editable                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] rel_type                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] user_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] product_id                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []


JSON Sample
-------------------------------------
{    "id": "cqIisJfTTEFGxipAmRwSDmOXO",    "rel_id": "jgecijZMuLtfSDKxSMDvYWCHA",    "key": "FgqxlNdNtwbvSoWjtFqlIKkmy",    "value": "JCwhsQIUUeKweuQZSPEwlvxiv",    "editable": 96,    "rel_type": 1,    "user_id": "OAKqwaExIlSKHbRQgChweysjT",    "product_id": "dHmSsxYtjPfYWCrFkUjwsnBNr"}


Comments
-------------------------------------
[ 0] Warning table: v_tag_with_product_id does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_tag_with_product_id primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Tag_with_product_id_FIELD_NAME_id = "id"

	Tag_with_product_id_FIELD_NAME_rel_id = "rel_id"

	Tag_with_product_id_FIELD_NAME_key = "key"

	Tag_with_product_id_FIELD_NAME_value = "value"

	Tag_with_product_id_FIELD_NAME_editable = "editable"

	Tag_with_product_id_FIELD_NAME_rel_type = "rel_type"

	Tag_with_product_id_FIELD_NAME_user_id = "user_id"

	Tag_with_product_id_FIELD_NAME_product_id = "product_id"
)

// Tag_with_product_id struct is a row record of the v_tag_with_product_id table in the  database
type Tag_with_product_id struct {
	ID        string `json:"id"`         //id
	RelID     string `json:"rel_id"`     //rel_id
	Key       string `json:"key"`        //key
	Value     string `json:"value"`      //value
	Editable  int32  `json:"editable"`   //editable
	RelType   int32  `json:"rel_type"`   //rel_type
	UserID    string `json:"user_id"`    //user_id
	ProductID string `json:"product_id"` //product_id

}

var Tag_with_product_idTableInfo = &TableInfo{
	Name: "v_tag_with_product_id",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_tag_with_product_id does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_tag_with_product_id primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "rel_id",
			Comment:            `rel_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RelID",
			GoFieldType:        "string",
			JSONFieldName:      "rel_id",
			ProtobufFieldName:  "rel_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "key",
			Comment:            `key`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
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
			Comment:            `value`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "editable",
			Comment:            `editable`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Editable",
			GoFieldType:        "int32",
			JSONFieldName:      "editable",
			ProtobufFieldName:  "editable",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "rel_type",
			Comment:            `rel_type`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "RelType",
			GoFieldType:        "int32",
			JSONFieldName:      "rel_type",
			ProtobufFieldName:  "rel_type",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "user_id",
			Comment:            `user_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "UserID",
			GoFieldType:        "string",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "product_id",
			Comment:            `product_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ProductID",
			GoFieldType:        "string",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Tag_with_product_id) TableName() string {
	return "v_tag_with_product_id"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Tag_with_product_id) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Tag_with_product_id) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Tag_with_product_id) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Tag_with_product_id) TableInfo() *TableInfo {
	return Tag_with_product_idTableInfo
}
