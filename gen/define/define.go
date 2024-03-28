package define

type Enums struct {
	// 枚举的字段名称
	Key string `json:"key"`
	// 枚举值
	Value int32 `json:"value"`
	// 备注信息
	Note string `json:"note"`
}

// enum模板项
type EnumsItem struct {
	// 组名称
	Group string `json:"group"`
	// 枚举名称
	Name string `json:"name"`
	// 枚举注释
	Note string `json:"note"`
	// 数据项
	Items []*Enums `json:"items"`
}

// enum模板数据源
type EnumsSource struct {
	// 组名称
	Group string
	// 数据项
	Enums []*EnumsItem
}
