package ifile

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"gitee.com/bobo-rs/idea-space-framework/pkg/img"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"io"
	"os"
)

// FileAttract 文件属性
type FileAttract struct {
	Path string
	Ctx context.Context
}

// FileInfo 文件详情
type FileInfo struct {
	Size int64 // 内存大小
	Width int // 宽度
	Height int // 高度
	Name string // 文件名
	Content []byte // 文件内容
	Hash string // hash值
}

// New 实例
func New(ctx context.Context, path string) *FileAttract {
	return &FileAttract{
		Ctx: ctx,
		Path: path,
	}
}

// Open 打开文件内容
func (a FileAttract) Open() (*os.File, error) {
	return os.Open(a.Path)
}

// Content 读取文件二进制内容
func (a FileAttract) Content() ([]byte, error) {
	// 打开文件
	file, err := a.Open()
	if err != nil {
		return nil, gerror.Wrapf(err, `打开File失败%s`, err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	// 转换成byte
	return ReadContent(file)
}

// MustContentString 读取文件内容转字符串
func (a FileAttract) MustContentString() string {
	buffer, err := a.Content()
	if err != nil {
		g.Log().Debug(a.Ctx, err)
		return ""
	}
	return string(buffer)
}

// Hasher 读取文件Hash值
func (a FileAttract) Hasher() (string, error) {
	// 打开文件
	f, err := a.Open()
	if err != nil {
		return "", gerror.Wrapf(err, `打开File失败%s`, err.Error())
	}

	// 关闭文件
	defer func() {
		_ = f.Close()
	}()

	// 读取hash内容byte并转换为十六进制
	return ReadHasher(f)
}

// MustHasherString 必须创建获取Hash值
func (a FileAttract) MustHasherString() string {
	hasher , err := a.Hasher()
	if err != nil {
		g.Log().Debug(a.Ctx, err)
		return ""
	}
	return hasher
}

// Attract 获取文件详情属性
func (a FileAttract) Attract() (*FileInfo, error) {

	// 打开文件
	f, err := a.Open()
	if err != nil {
		return nil, gerror.Wrapf(err, `打开File失败%s`, err.Error())
	}

	// 关闭文件
	defer func() {
		_ = f.Close()
	}()

	// 读取文件详情
	fs, err := f.Stat()
	if err != nil {
		return nil, gerror.Wrapf(err, `获取File详情失败%s`, err.Error())
	}
	// 初始化
	item := &FileInfo{
		Name: fs.Name(),
		Size: fs.Size(),
	}

	// 获取文件hash值
	item.Hash, err = ReadHasher(f)
	if err != nil {
		return nil, err
	}

	// 读取文件内容
	item.Content, err = ReadContent(f)
	if err != nil {
		return nil, err
	}

	// 图片类获取尺寸
	if img.IsValidImageExt(gfile.ExtName(a.Path)) {
		item.Width, item.Height, err = ReadImageWH(f)
		if err != nil {
			return nil, err
		}
	}

	return item, nil
}

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