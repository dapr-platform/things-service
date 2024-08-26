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


Table: o_kpi_info
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] label                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] description                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] unit                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] product_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] interval                                       INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] type                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 8] calc_script                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 9] summary_type                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] value_type                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] org_id                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "XxZAxVNOVJQUmCrPZrfDLeWNV",    "name": "cOEIvRbwUleaKLQnGirKhCjFP",    "label": "QEyKXrAGrLndvsxbRbpKuHNcQ",    "description": "FONuwxPWdUmENhETvoUMBmPuT",    "unit": "lnTIfOCswAMcXhvPEncftKktY",    "product_name": "XZITIoOijsAhhefCyrxHZBZkR",    "interval": 35,    "type": 16,    "calc_script": "SAPTuiYwZdpTcOrKXscBCImrl",    "summary_type": "dZbrZiCZYKCMkfwLCFHMFRPai",    "value_type": "UFvrsgxIvoUEqtnAijLBJeDXg",    "org_id": "bPXmjjnldWfltuKuCakaCfuTr"}



*/

var (
	Kpi_info_FIELD_NAME_id = "id"

	Kpi_info_FIELD_NAME_name = "name"

	Kpi_info_FIELD_NAME_label = "label"

	Kpi_info_FIELD_NAME_description = "description"

	Kpi_info_FIELD_NAME_unit = "unit"

	Kpi_info_FIELD_NAME_product_name = "product_name"

	Kpi_info_FIELD_NAME_interval = "interval"

	Kpi_info_FIELD_NAME_type = "type"

	Kpi_info_FIELD_NAME_calc_script = "calc_script"

	Kpi_info_FIELD_NAME_summary_type = "summary_type"

	Kpi_info_FIELD_NAME_value_type = "value_type"

	Kpi_info_FIELD_NAME_org_id = "org_id"
)

// Kpi_info struct is a row record of the o_kpi_info table in the  database
type Kpi_info struct {
	ID          string `json:"id"`           //唯一标识
	Name        string `json:"name"`         //英文名
	Label       string `json:"label"`        //中文名
	Description string `json:"description"`  //描述
	Unit        string `json:"unit"`         //单位
	ProductName string `json:"product_name"` //计算产品名称
	Interval    int32  `json:"interval"`     //分钟粒度 1,5,15,30,60
	Type        int32  `json:"type"`         //type
	CalcScript  string `json:"calc_script"`  //计算公式
	SummaryType string `json:"summary_type"` //统计类型
	ValueType   string `json:"value_type"`   //值类型
	OrgID       string `json:"org_id"`       //组织id

}

var Kpi_infoTableInfo = &TableInfo{
	Name: "o_kpi_info",
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
			Comment:            `英文名`,
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
			Name:               "label",
			Comment:            `中文名`,
			Notes:              ``,
			Nullable:           true,
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
			Name:               "description",
			Comment:            `描述`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "unit",
			Comment:            `单位`,
			Notes:              ``,
			Nullable:           true,
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

		&ColumnInfo{
			Index:              5,
			Name:               "product_name",
			Comment:            `计算产品名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ProductName",
			GoFieldType:        "string",
			JSONFieldName:      "product_name",
			ProtobufFieldName:  "product_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "interval",
			Comment:            `分钟粒度 1,5,15,30,60`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Interval",
			GoFieldType:        "int32",
			JSONFieldName:      "interval",
			ProtobufFieldName:  "interval",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "type",
			Comment:            `type`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "calc_script",
			Comment:            `计算公式`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "CalcScript",
			GoFieldType:        "string",
			JSONFieldName:      "calc_script",
			ProtobufFieldName:  "calc_script",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "summary_type",
			Comment:            `统计类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "SummaryType",
			GoFieldType:        "string",
			JSONFieldName:      "summary_type",
			ProtobufFieldName:  "summary_type",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "value_type",
			Comment:            `值类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ValueType",
			GoFieldType:        "string",
			JSONFieldName:      "value_type",
			ProtobufFieldName:  "value_type",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "org_id",
			Comment:            `组织id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "OrgID",
			GoFieldType:        "string",
			JSONFieldName:      "org_id",
			ProtobufFieldName:  "org_id",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (k *Kpi_info) TableName() string {
	return "o_kpi_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (k *Kpi_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (k *Kpi_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (k *Kpi_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (k *Kpi_info) TableInfo() *TableInfo {
	return Kpi_infoTableInfo
}
