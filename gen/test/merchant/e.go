package merchant

// 商户等级
const (
	// 免费版
	LevelGratis = 1
	// 旗舰版
	LevelUltimate = 2
	// 经典版
	LevelClassic = 3
	// 超店版
	LevelSuperGold = 4
	// 智尊版
	LevelWisdomHonor = 5
)

// 商户等级
var LevelTypes = map[int32]string{
	// 免费版
	LevelGratis: "免费版",
	// 旗舰版
	LevelUltimate: "旗舰版",
	// 经典版
	LevelClassic: "经典版",
	// 超店版
	LevelSuperGold: "超店版",
	// 智尊版
	LevelWisdomHonor: "智尊版",
}

// 商户等级切片
var LevelSlice = []map[int32]string{
	// 免费版
	{LevelGratis: "免费版"},
	// 旗舰版
	{LevelUltimate: "旗舰版"},
	// 经典版
	{LevelClassic: "经典版"},
	// 超店版
	{LevelSuperGold: "超店版"},
	// 智尊版
	{LevelWisdomHonor: "智尊版"},
}

// 商户状态
const (
	// 已使用
	StatusUse = 1
	// 未使用
	StatusUn = 2
)

// 商户状态
var StatusTypes = map[int32]string{
	// 已使用
	StatusUse: "已使用",
	// 未使用
	StatusUn: "未使用",
}

// 商户状态切片
var StatusSlice = []map[int32]string{
	// 已使用
	{StatusUse: "已使用"},
	// 未使用
	{StatusUn: "未使用"},
}
