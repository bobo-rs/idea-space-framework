package ifile

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"image"
	_ "image/gif"  // 导入gif解码器
	_ "image/jpeg" // 导入jpeg解码器
	_ "image/png"  // 导入png解码器
	"io"
	"os"
)

// ReadHasher 读取Hash值内容
func ReadHasher(fs *os.File) (string, error) {
	// 创建Hash字符串
	hasher := sha256.New()

	// 将内容赋值到hasher中
	if _, err := io.Copy(hasher, fs); err != nil {
		return "", gerror.Wrapf(err, `将内容赋值到hash失败%s`, err.Error())
	}

	// 将byte转换为十六进制
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// ReadContent 读取文件内容
func ReadContent(fs *os.File) ([]byte, error) {
	// 创建一个byte切片
	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, fs); err != nil {
		return nil, gerror.Wrapf(err, `读取文件内容失败%s`, err.Error())
	}

	// 读取为[]byte
	return buffer.Bytes(), nil
}


// ReadImageWH 读取图片文件高宽度
func ReadImageWH(fs *os.File) (w, h int, err error) {
	// 创建Image
	im, _, err := image.Decode(fs)
	if err != nil {
		return 0, 0, gerror.Wrapf(err, `打开图片失败%s`, err.Error())
	}

	// 获取图片边框
	bounds := im.Bounds()
	return bounds.Dx(), bounds.Dy(), nil
}