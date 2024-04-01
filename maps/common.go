package maps

type (
	// 定义一个新的整型对应字符串的MAP，它的底层是一个map[uint]string
	CommonUintStringMap map[uint]string

	// 定义一个新的整型对应字符串的MAP，它的底层是一个map[string]string
	CommonStringStringMap map[string]string

)
var (
	// ChannelId 渠道平台MAP
	ChannelId = CommonUintStringMap{
		0: `平台`,
		1: `企业`,
		2: `用户`,
	}

	// VipType 素材Vip类型
	VipType = CommonUintStringMap{
		0: `免费`,
		1: `VIP`,
	}

	// Language 语言
	Language = CommonStringStringMap{
		`zh-CN`: `简体中文`,
		`zh-HK`: `繁体中文`,
		`en`: `英语`,
	}


)

