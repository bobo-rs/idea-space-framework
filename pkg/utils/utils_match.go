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