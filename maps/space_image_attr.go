package maps

var (
	// SpaceImageAttrStorageType 图片空间存储类型
	SpaceImageAttrStorageType = CommonStringStringMap{
		`local`: `本地`,
		`qiniu`: `七牛`,
		`ali_oss`: `阿里云OSS`,
		`tencent_oss`: `腾讯云OSS`,
		`baidu_oss`: `百度云OSS`,
	}
)

