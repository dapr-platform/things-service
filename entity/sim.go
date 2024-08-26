package entity

type SimRule struct {
	ProductId          string   `json:"product_id"`                     //产品id
	ChooseDeviceType   string   `json:"choose_device_type""`            //all,random,fixed
	RandomCount        int      `json:"random_count,omitempty"`         //随机数量
	FixedDeviceIds     []string `json:"fixed_device_ids,omitempty"`     //固定选择设备id
	SimIntervalSeconds int      `json:"sim_interval_seconds,omitempty"` //sim interval seconds,default 10s
}
