package entity

import "github.com/dapr-platform/common"

type DeviceInfo struct {
	Id             string           `json:"id"`
	CreatedBy      string           `json:"created_by"`
	CreatedTime    common.LocalTime `json:"created_time"`
	UpdatedBy      string           `json:"updated_by"`
	UpdatedTime    common.LocalTime `json:"updated_time"`
	Name           string           `json:"name"`
	Type           int              `json:"type"`
	Status         int              `json:"status"`
	ParentId       string           `json:"parent_id"`
	GroupId        string           `json:"group_id"`
	ProductId      string           `json:"product_id"`
	ProtocolConfig string           `json:"protocol_config"`
	Enabled        int              `json:"enabled"`
	Identifier     string           `json:"identifier"`
	Tags           []string         `json:"tags"`
	ProductName    string           `json:"product_name"`
	JsonData       string           `json:"json_data"`
}

type DeviceStatusStatistics struct {
	Fault  int `json:"fault"`  //故障
	Alert  int `json:"alert"`  //告警
	Loss   int `json:"loss"`   //失联
	Normal int `json:"normal"` //正常
}

type UserDeviceInfo struct {
	ID          string           `json:"id"`           //id
	UserID      string           `json:"user_id"`      //user_id
	DeviceID    string           `json:"device_id"`    //device_id
	CreatedTime common.LocalTime `json:"created_time"` //created_time
	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time
	Index       int32            `json:"index"`        //index
	Locked      int32            `json:"locked"`       //locked
	Name        string           `json:"name"`         //name
	Type        int32            `json:"type"`         //type
	Enabled     int32            `json:"enabled"`      //enabled
	Identifier  string           `json:"identifier"`   //identifier
	Status      int32            `json:"status"`       //status
	Tags        []string         `json:"tags"`         //tags
	JSONData    string           `json:"json_data"`    //json_data
}
