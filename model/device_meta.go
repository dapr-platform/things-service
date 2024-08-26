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


Table: p_device_meta
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] descript                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] chart_data                                     TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] meta_data                                      TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 9] encode_script                                  TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] decode_script                                  TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[11] category                                       VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] vendor                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] model                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[14] cover_file                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[15] direct_connect                                 INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[16] mqtt_receive_topic_prefix                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[17] mqtt_send_topic_prefix                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "KmbuRRKifpaRjdvckpUMvNfgO",    "created_by": "laYohYeoEocnvDmQEWoAPUqAm",    "created_time": 8,    "updated_by": "PWaOjTBFmpWMDXoJLvhCTwZDt",    "updated_time": 5,    "descript": "xVnmXHtkVhnAJdxpvLwQyvnvK",    "name": "woQFSjHCWtYJsxnJvOjYWorPP",    "chart_data": "qTSFqgPLaaurNKjRoUASUxqLO",    "meta_data": "BLJimedIxQVUPgaFnZxISRdkT",    "encode_script": "BqGKTxlWDkSovnRlxilafoaPY",    "decode_script": "MSwkurmAmnqdgxgJKVGQUYrMJ",    "category": "AUeaiCRtQUrOkuoenAmNnRDlv",    "vendor": "pxBaMroOaZCwOdJpmfxZLNpLi",    "model": "fDxLZPjGpVJpgrcZMJrmCKLSR",    "cover_file": "bpJckrREwOmwFFitHOMhPurqc",    "direct_connect": 48,    "mqtt_receive_topic_prefix": "CKIoYIEhwNSVvadeCyLXKxwnq",    "mqtt_send_topic_prefix": "ccphRngFMXMpccYeAEqeGtMfO"}



*/

var (
	Device_meta_FIELD_NAME_id = "id"

	Device_meta_FIELD_NAME_created_by = "created_by"

	Device_meta_FIELD_NAME_created_time = "created_time"

	Device_meta_FIELD_NAME_updated_by = "updated_by"

	Device_meta_FIELD_NAME_updated_time = "updated_time"

	Device_meta_FIELD_NAME_descript = "descript"

	Device_meta_FIELD_NAME_name = "name"

	Device_meta_FIELD_NAME_chart_data = "chart_data"

	Device_meta_FIELD_NAME_meta_data = "meta_data"

	Device_meta_FIELD_NAME_encode_script = "encode_script"

	Device_meta_FIELD_NAME_decode_script = "decode_script"

	Device_meta_FIELD_NAME_category = "category"

	Device_meta_FIELD_NAME_vendor = "vendor"

	Device_meta_FIELD_NAME_model = "model"

	Device_meta_FIELD_NAME_cover_file = "cover_file"

	Device_meta_FIELD_NAME_direct_connect = "direct_connect"

	Device_meta_FIELD_NAME_mqtt_receive_topic_prefix = "mqtt_receive_topic_prefix"

	Device_meta_FIELD_NAME_mqtt_send_topic_prefix = "mqtt_send_topic_prefix"
)

// Device_meta struct is a row record of the p_device_meta table in the thingsdb database
type Device_meta struct {
	ID                     string           `json:"id"`                        //唯一标识
	CreatedBy              string           `json:"created_by"`                //创建人
	CreatedTime            common.LocalTime `json:"created_time"`              //创建时间
	UpdatedBy              string           `json:"updated_by"`                //更新人
	UpdatedTime            common.LocalTime `json:"updated_time"`              //更新时间
	Descript               string           `json:"descript"`                  //描述
	Name                   string           `json:"name"`                      //名称
	ChartData              string           `json:"chart_data"`                //图表呈现元数据
	MetaData               string           `json:"meta_data"`                 //元数据(json定义,属性、服务、事件)
	EncodeScript           string           `json:"encode_script"`             //发送数据脚本处理
	DecodeScript           string           `json:"decode_script"`             //接收数据脚本处理
	Category               string           `json:"category"`                  //分类
	Vendor                 string           `json:"vendor"`                    //厂商
	Model                  string           `json:"model"`                     //型号
	CoverFile              string           `json:"cover_file"`                //封面base64
	DirectConnect          int32            `json:"direct_connect"`            //是否直连（1:直连，2:网关转发）
	MqttReceiveTopicPrefix string           `json:"mqtt_receive_topic_prefix"` //接收消息队列前缀
	MqttSendTopicPrefix    string           `json:"mqtt_send_topic_prefix"`    //发送消息队列前缀

}

var Device_metaTableInfo = &TableInfo{
	Name: "p_device_meta",
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "chart_data",
			Comment:            `图表呈现元数据`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "ChartData",
			GoFieldType:        "string",
			JSONFieldName:      "chart_data",
			ProtobufFieldName:  "chart_data",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "meta_data",
			Comment:            `元数据(json定义,属性、服务、事件)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "MetaData",
			GoFieldType:        "string",
			JSONFieldName:      "meta_data",
			ProtobufFieldName:  "meta_data",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "encode_script",
			Comment:            `发送数据脚本处理`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "EncodeScript",
			GoFieldType:        "string",
			JSONFieldName:      "encode_script",
			ProtobufFieldName:  "encode_script",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "decode_script",
			Comment:            `接收数据脚本处理`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "DecodeScript",
			GoFieldType:        "string",
			JSONFieldName:      "decode_script",
			ProtobufFieldName:  "decode_script",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "category",
			Comment:            `分类`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "Category",
			GoFieldType:        "string",
			JSONFieldName:      "category",
			ProtobufFieldName:  "category",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "vendor",
			Comment:            `厂商`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Vendor",
			GoFieldType:        "string",
			JSONFieldName:      "vendor",
			ProtobufFieldName:  "vendor",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "model",
			Comment:            `型号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Model",
			GoFieldType:        "string",
			JSONFieldName:      "model",
			ProtobufFieldName:  "model",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "cover_file",
			Comment:            `封面base64`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "direct_connect",
			Comment:            `是否直连（1:直连，2:网关转发）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "DirectConnect",
			GoFieldType:        "int32",
			JSONFieldName:      "direct_connect",
			ProtobufFieldName:  "direct_connect",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "mqtt_receive_topic_prefix",
			Comment:            `接收消息队列前缀`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "MqttReceiveTopicPrefix",
			GoFieldType:        "string",
			JSONFieldName:      "mqtt_receive_topic_prefix",
			ProtobufFieldName:  "mqtt_receive_topic_prefix",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "mqtt_send_topic_prefix",
			Comment:            `发送消息队列前缀`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "MqttSendTopicPrefix",
			GoFieldType:        "string",
			JSONFieldName:      "mqtt_send_topic_prefix",
			ProtobufFieldName:  "mqtt_send_topic_prefix",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Device_meta) TableName() string {
	return "p_device_meta"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Device_meta) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Device_meta) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Device_meta) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Device_meta) TableInfo() *TableInfo {
	return Device_metaTableInfo
}
