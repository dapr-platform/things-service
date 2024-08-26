package entity

type DeviceMetaInfo struct {
	Option     DeviceMetaOption     `json:"option"`
	Properties []DeviceMetaProperty `json:"properties"`
	Events     []DeviceMetaEvent    `json:"events"`
	Services   []DeviceMetaService  `json:"services"`
}
type DeviceMetaOption struct {
	Classify string `json:"classify"`
	CatValue string `json:"catValue"`
}

type DeviceMetaProperty struct {
	DataType   string `json:"dataType"`
	DataRange  string `json:"dataRange"`
	StepLength any    `json:"stepLength"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	Unit       string `json:"unit"`
}

type DeviceMetaEvent struct {
	EventName        string `json:"eventName"`
	EventExpr        string `json:"eventExpr"`
	EventDescription string `json:"eventDescription"`
}
type DeviceMetaService struct {
	ServiceName        string `json:"serviceName"`
	Param              string `json:"param"`
	ServiceDescription string `json:"serviceDescription"`
}
