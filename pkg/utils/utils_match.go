package utils

import (
	"github.com/bobo-rs/idea-space-framework/consts"
	"regexp"
)

// ScanIsEmpty 扫描数据验证是否为空
func ScanIsEmpty(errorMsg string) bool {
	re := regexp.MustCompile(consts.MatchScanIsEmpty)
	return re.MatchString(errorMsg)
}

// AlphaNum 验证是否字母和数字
func AlphaNum(s string) bool {
	re := regexp.MustCompile(consts.AlphaNum)
	return re.MatchString(s)
}

// ChsDash 匹配汉字、字母、数字特殊字符
func ChsDash(s string) bool {
	re := regexp.MustCompile(consts.ChsDash)
	return re.MatchString(s)
}

// ChsAlphaNum 匹配汉字、字母、数字
func ChsAlphaNum(s string) bool {
	re := regexp.MustCompile(consts.ChsAlphaNum)
	return re.MatchString(s)
}

