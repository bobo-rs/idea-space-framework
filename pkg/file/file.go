package file

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"io"
	"os"
)

// Hash 读取文件内容转换为hash值
func Hash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = f.Close()
	}()

	// 创建Hash值
	hasher := sha256.New()

	// 将内容写入哈希值
	if _, err = io.Copy(hasher, f); err != nil {
		return "", gerror.Wrapf(err, `读取文件内容失败%s`, err.Error())
	}

	// 获取hash字节
	hashBytes := hasher.Sum(nil)

	// 转换十六进制
	return hex.EncodeToString(hashBytes), nil
}

// HashString 获取文件hash值
func HashString(path string) string {
	h , _ := Hash(path)
	return h
}
