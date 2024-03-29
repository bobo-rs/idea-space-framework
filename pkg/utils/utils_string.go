package utils

import (
	"github.com/bobo-rs/idea-space-framework/pkg/encrypt"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
	"unicode"
)

// IsSeparator 返回一个布尔值，表示rune是否为分隔符
func IsSeparator(r rune) bool {
	// 检查是否是逗号（中英文）、空格、制表符、回车符或换行符
	return r == ',' || r == '，' || unicode.IsSpace(r)
}

// SeparatorStringToArray 风格拆分字符串为数组字符串
func SeparatorStringToArray(s string) []string {
	// 分隔符
	slices := strings.FieldsFunc(s, IsSeparator)
	if len(slices) == 0 {
		return nil
	}
	return slices
}

// HashMd5 Md5加密
func HashMd5(slices ...string) (string, error) {
	if len(slices) == 0 {
		return "", gerror.New(`缺少Hash参数`)
	}
	// 拼接hash字符串
	return encrypt.NewM().MustEncryptString(strings.Join(slices, "")), nil
}

// HashMd5String Md5加密字符串
func HashMd5String(slices ...string) string {
	hash, _ := HashMd5(slices...)
	return hash
}