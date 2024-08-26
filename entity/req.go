package entity

type ServiceInvokeReq struct {
	DeviceIdentifier  string `json:"device_identifier"`
	ServiceIdentifier string `json:"service_identifier"`
	InputParam        any    `json:"input_param,omitempty"`
}
type PropertySetReq struct {
	DeviceIdentifier   string `json:"device_identifier"`
	PropertyIdentifier string `json:"property_identifier"`
	Value              any    `json:"value"`
}

type SimDeviceMirrorReq struct {
	DeviceIdentifier string `json:"device_identifier"`
	JsonData         string `json:"json_data"`
}

type FileUploadReq struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}
