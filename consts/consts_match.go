package consts

const (
	ChsDash = `^[\x{4e00}-\x{9fa5}0-9a-zA-Z\s.,，!?;:\'\"-]+$` // 匹配中英文数字以及中英文标点符号
	ChsAlphaNum = `^[\x{4e00}-\x{9fa5}0-9a-zA-Z]+$` // 汉字、字母、数字
	AlphaNum = `^[0-9a-zA-Z]+$` // 字母、数字
	MatchScanIsEmpty = `sql: no rows in result set` // 匹配Scan数据是否为空
)

