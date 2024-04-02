package img

import (
	"gitee.com/bobo-rs/idea-space-framework/pkg/file"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"image"
	_ "image/gif"  // 导入gif解码器
	_ "image/jpeg" // 导入jpeg解码器
	_ "image/png"  // 导入png解码器
	"os"
)

// ImageAttr 图片基础属性
type ImageAttr struct {
	Width int
	Height int
	Size int64
	Hash string
}

// ReadImageAttr 读取图片详情属性
func ReadImageAttr(path string) (attr *ImageAttr, err error) {
	attr = &ImageAttr{}
	// 验证扩展
	if IsValidImageExt(gfile.ExtName(path)) == false {
		return nil, gerror.New(`非图片扩展后缀`)
	}
	// 打开图片
	f , err := os.Open(path)
	if err != nil {
		return nil, gerror.Wrap(err, `图片文件打不开`)
	}
	defer func() {
		_ = f.Close()
	}()

	// 获取文件详情
	fileDetail, err := f.Stat()
	if err != nil {
		return nil, gerror.Wrap(err, `读取文件内容失败`)
	}

	// 读取图片内容
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, gerror.Wrapf(err, `打开图片失败`, err.Error())
	}

	// 读取width, height, size
	bounds := img.Bounds()
	return &ImageAttr{
		Width: bounds.Dx(),
		Height: bounds.Dy(),
		Size: fileDetail.Size(),
		Hash: file.HashString(path),
	}, nil
}

// IsValidImageExt 验证图片扩展
func IsValidImageExt(ext string) bool {
	switch ImageExt(ext) {
	case ImageExtGif, ImageExtJpeg, ImageExtPng, ImageExtJpg:
		return true
	default:
		return false
	}
}