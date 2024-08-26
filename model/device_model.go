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


Table: o_device_model
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] descript                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] categories                                     VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] json_data                                      TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 9] service_script                                 TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] property_script                                TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[11] event_script                                   TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[12] cover_file                                     TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "SDudIIvgKpGfSKRIhcvglTwjW",    "created_by": "psFsbAWfmvadFRbVlBxjtNwrq",    "created_time": 80,    "updated_by": "tCmRlyTQBDYBmrSYrWxmGFrqJ",    "updated_time": 57,    "name": "eZMporJxnvQpMrvIRAbMXYdAv",    "descript": "uyoaKPvFQFFJjuYceVOsQUigC",    "categories": "lqmWAyiRtPLbsQdxBkDlcHvbD",    "json_data": "PUXpMWSPMedLNgIvNgQTWtkJn",    "service_script": "mvPnRCEvPfdbefNklmgBmpPjB",    "property_script": "EAOZMRtyWdpgsXvKByAVmlbwx",    "event_script": "TeWPmlCjQjglleuhIWRgaiepM",    "cover_file": "BSOdhoEbxXWHsYiWGsdBiADqU"}



*/

var (
	Device_model_FIELD_NAME_id = "id"

	Device_model_FIELD_NAME_created_by = "created_by"

	Device_model_FIELD_NAME_created_time = "created_time"

	Device_model_FIELD_NAME_updated_by = "updated_by"

	Device_model_FIELD_NAME_updated_time = "updated_time"

	Device_model_FIELD_NAME_name = "name"

	Device_model_FIELD_NAME_descript = "descript"

	Device_model_FIELD_NAME_categories = "categories"

	Device_model_FIELD_NAME_json_data = "json_data"

	Device_model_FIELD_NAME_service_script = "service_script"

	Device_model_FIELD_NAME_property_script = "property_script"

	Device_model_FIELD_NAME_event_script = "event_script"

	Device_model_FIELD_NAME_cover_file = "cover_file"
)

// Device_model struct is a row record of the o_device_model table in the  database
type Device_model struct {
	ID             string           `json:"id"`              //唯一标识
	CreatedBy      string           `json:"created_by"`      //创建人
	CreatedTime    common.LocalTime `json:"created_time"`    //创建时间
	UpdatedBy      string           `json:"updated_by"`      //更新人
	UpdatedTime    common.LocalTime `json:"updated_time"`    //更新时间
	Name           string           `json:"name"`            //名称
	Descript       string           `json:"descript"`        //描述
	Categories     string           `json:"categories"`      //分类（A/B/C 手动填写)
	JSONData       string           `json:"json_data"`       //物模型数据(json定义,属性、服务、事件)
	ServiceScript  string           `json:"service_script"`  //服务数据转换脚本处理
	PropertyScript string           `json:"property_script"` //属性数据转换脚本处理
	EventScript    string           `json:"event_script"`    //事件转换脚本处理
	CoverFile      string           `json:"cover_file"`      //封面base64

}

var Device_modelTableInfo = &TableInfo{
	Name: "o_device_model",
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
			Name:               "created_by",
			Comment:            `创建人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "CreatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "created_by",
			ProtobufFieldName:  "created_by",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "created_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "created_time",
			ProtobufFieldName:  "created_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "updated_by",
			Comment:            `更新人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "UpdatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "updated_by",
			ProtobufFieldName:  "updated_by",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "updated_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "descript",
			Comment:            `描述`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Descript",
			GoFieldType:        "string",
			JSONFieldName:      "descript",
			ProtobufFieldName:  "descript",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "categories",
			Comment:            `分类（A/B/C 手动填写)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Categories",
			GoFieldType:        "string",
			JSONFieldName:      "categories",
			ProtobufFieldName:  "categories",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "json_data",
			Comment:            `物模型数据(json定义,属性、服务、事件)`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "service_script",
			Comment:            `服务数据转换脚本处理`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "ServiceScript",
			GoFieldType:        "string",
			JSONFieldName:      "service_script",
			ProtobufFieldName:  "service_script",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "property_script",
			Comment:            `属性数据转换脚本处理`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "PropertyScript",
			GoFieldType:        "string",
			JSONFieldName:      "property_script",
			ProtobufFieldName:  "property_script",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "event_script",
			Comment:            `事件转换脚本处理`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "EventScript",
			GoFieldType:        "string",
			JSONFieldName:      "event_script",
			ProtobufFieldName:  "event_script",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "cover_file",
			Comment:            `封面base64`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "CoverFile",
			GoFieldType:        "string",
			JSONFieldName:      "cover_file",
			ProtobufFieldName:  "cover_file",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Device_model) TableName() string {
	return "o_device_model"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Device_model) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Device_model) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Device_model) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Device_model) TableInfo() *TableInfo {
	return Device_modelTableInfo
}
