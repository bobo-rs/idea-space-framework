package encrypt

import (
	"crypto/md5"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type Md5 struct {
	bytes []byte
}

// EncryptBytes 使用Md5加密
func (m Md5) EncryptBytes() (string, error) {
	h := md5.New()
	// 写入数据
	if _, err := h.Write(m.bytes); err != nil {
		return "", gerror.New(`hash.Write failed`)
	}
	// 拼接秘传
	return fmt.Sprintf(`%x`, h.Sum(nil)), nil
}

// EncryptString 使用Md5加密
func (m Md5) EncryptString(s string) (string, error) {
	m.bytes = []byte(s)
	return m.EncryptBytes()
}

// MustEncryptString 使用Md5加密
func (m Md5) MustEncryptString(s string) string {
	hash, _ := m.EncryptString(s)
	return hash
}

// Encrypt 使用Md5加密
func (m Md5) Encrypt(data interface{}) (string, error) {
	m.bytes = gconv.Bytes(data)
	return m.EncryptBytes()
}

// MustEncrypt 使用Md5加密
func (m Md5) MustEncrypt(data interface{}) string {
	hash, _ := m.Encrypt(data)
	return hash
}