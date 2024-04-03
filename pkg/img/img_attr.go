package img

import (
	_ "image/gif"  // 导入gif解码器
	_ "image/jpeg" // 导入jpeg解码器
	_ "image/png"  // 导入png解码器
)

// ImageAttr 图片基础属性
type ImageAttr struct {
	Width int
	Height int
	Size int64
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