package ifile

import (
	"gitee.com/bobo-rs/idea-space-framework/pkg/img"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"image"
	_ "image/gif"  // 导入gif解码器
	_ "image/jpeg" // 导入jpeg解码器
	_ "image/png"  // 导入png解码器
	"os"
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