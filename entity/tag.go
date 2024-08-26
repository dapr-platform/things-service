package entity

type BatchTagInfo struct {
	RelId   string `json:"rel_id"`   // 对象id
	RelType int    `json:"rel_type"` //对象类型 1:设备,2:网关，3:点位
	Tags    []Tag  `json:"tags"`
}

type Tag struct {
	Key      string `json:"key"`      //标签名称
	Value    any    `json:"value"`    //标签值
	Editable int    `json:"editable"` //是否可以修改编辑。（设备导入的标签不可编辑）
}

type TagInfo struct {
	Key    string   `json:"key"`    //tag名称
	Values []string `json:"values"` //tag values
}
