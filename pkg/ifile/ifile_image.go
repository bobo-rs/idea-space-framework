package ifile

import (
	"gitee.com/bobo-rs/idea-space-framework/pkg/img"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

// ImageWH 打开图片文件获取宽度和高度
func (a FileAttract) ImageWH() (w, h int, err error) {
	// 验证是否图片
	if !img.IsValidImageExt(gfile.ExtName(a.Path)) {
		return 0, 0, nil
	}

	// 打开文件图片
	f, err := a.Open()
	if err != nil {
		return 0, 0, gerror.Wrapf(err, `打开文件错误%s`, err.Error())
	}
	// 关闭文件
	defer func() {
		_ = f.Close()
	}()
	// 读取文件图片
	return ReadImageWH(f)
}
