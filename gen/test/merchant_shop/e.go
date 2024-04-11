package merchant_shop

// 门店状态
const (
	// 已使用
	StatusUse = 1
	// 未使用
	StatusUn = 2
)

// 门店状态
var StatusTypes = map[int32]string{
	// 已使用
	StatusUse: "已使用",
	// 未使用
	StatusUn: "未使用",
}

// 门店状态切片
var StatusSlice = []map[int32]string{
	// 已使用
	{StatusUse: "已使用"},
	// 未使用
	{StatusUn: "未使用"},
}
