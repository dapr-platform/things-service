package entity

import "time"

type MonitorTotal struct {
	Total []MonitorTotalItem `json:"total"`
}

type MonitorTotalItem struct {
	Name   string `json:"name"`   //名称
	Total  int    `json:"total"`  //总数
	Normal int    `json:"normal"` //正常数
	Error  int    `json:"error"`  //异常数
}

type MonitorListHostItem struct {
	Name     string  `json:"name"`   //名称
	Ip       string  `json:"ip"`     //ip
	Status   int     `json:"status"` //状态,1:正常,0:离线
	Os       string  `json:"os"`
	Kernel   string  `json:"kernel"`
	CpuUsage float64 `json:"cpu_usage"` //cpu使用率
	MemUsage float64 `json:"mem_usage"` //内存使用率
}

type MonitorListDbItem struct {
	Name            string  `json:"name"`             //名称
	Server          string  `json:"server"`           //服务器
	ConnectionCount int     `json:"connection_count"` //连接数
	CacheHitRate    float64 `json:"cache_hit_rate"`   //缓存命中率
	DeadlockCount   int     `json:"deadlock_count"`   //死锁数
	ConflictCount   int     `json:"conflict_count"`   //冲突数
	Status          int     `json:"status"`           //运行状态，1：正常，0：停止
}

type MonitorListGatewayItem struct {
	Identifier  string `json:"identifier"`    //标识符
	Name        string `json:"name"`          //名称
	AIPortCount int    `json:"ai_port_count"` //ai端口数
	DIPortCount int    `json:"di_port_count"` //di端口数
	AOPortCount int    `json:"ao_port_count"` //ao端口数
	DOPortCount int    `json:"do_port_count"` //do端口数
	Status      int    `json:"status"`        //运行状态，1：正常，0：停止
}

type MonitorListInterfaceItem struct {
	Name             string `json:"name"`               //名称
	ThirdPartySystem string `json:"third_party_system"` //第三方系统
	Protocol         string `json:"protocol"`           //接口协议
	CallFrequency    int    `json:"call_frequency"`     //调用频率
	AlertCount       int    `json:"alert_count"`        //告警次数
	Status           int    `json:"status"`             //状态，1：正常，0：异常
	LatestDataTime   string `json:"latest_data_time"`
}

type MonitorGatewayDetail struct {
	Identifier string                      `json:"identifier"`
	Name       string                      `json:"name"`
	Status     int                         `json:"status"`
	Tags       []string                    `json:"tags"`
	Points     []MonitorGatewayDetailPoint `json:"points"`
}
type MonitorGatewayDetailPoint struct {
	Name             string `json:"name"`
	DeviceIdentifier string `json:"device_identifier"`
	AccessType       string `json:"access_type"` //接入类型 ， IO/485
	AddressIO        string `json:"address_io"`  //如果是io，则是AI1,DI1
	Address485       string `json:"address_485"` //如果是485，则是地址
	Value            string `json:"value"`       //值
	UpdateTime       string `json:"update_time"` //格式 2006-01-02 15:04:05
	Status           int    `json:"status"`      //状态，1：正常，0：异常
}

type MetricItem struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}
type MonitorListServiceItem struct {
	Name        string `json:"name"`         //名称
	Host        string `json:"host"`         //主机
	ThreadCount int    `json:"thread_count"` //线程数
	MemoryUsage int    `json:"memory_usage"` //占用内存数
	Status      int    `json:"status"`       //运行状态，1：正常，0：停止
}

type MonitorLog struct {
	Level string `json:"level"`
	Count string `json:"count"`
}
type QueryMetrics struct {
	Query string `json:"query"`
	Time  string `json:"time"`
}
type QueryRangeMetrics struct {
	Query string `json:"query"`
	// The boundaries of the time range.
	Start string `json:"start"` //格式 2006-01-02 15:04:05
	End   string `json:"end"`   //格式 2006-01-02 15:04:05
	// The maximum time between two slices within the boundaries.
	Step int `json:"step"` //单位秒
}

type QueryFormItem struct {
	Start string `json:"start" binding:"required"` //格式 2006-01-02 15:04:05
	End   string `json:"end" binding:"required"`   //格式 2006-01-02 15:04:05
	Step  int    `json:"step" binding:"required"`  //单位秒
	Query string `json:"query" binding:"required"`
}

type BatchQueryForm struct {
	Queries []QueryFormItem `json:"queries" binding:"required"`
}
type Range struct {
	// The boundaries of the time range.
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	// The maximum time between two slices within the boundaries.
	Step time.Duration `json:"step"`
}
