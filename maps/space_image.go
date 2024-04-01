package maps

var (
	// SpaceImageIsOpen 是否开放访问
	SpaceImageIsOpen = CommonUintStringMap{
		0: `开放`,
		1: `关闭`,
	}

	// SpaceImageDesignType 设计类型
	SpaceImageDesignType = CommonUintStringMap{
		0: `智能设计`,
		1: `手动设计`,
		2: `主动设计`,
	}

	// SpaceImageOcrIsOrigin 图片空间OCR-翻译类型
	SpaceImageOcrIsOrigin = CommonUintStringMap{
		0: `原文`,
		1: `翻译`,
	}
)

