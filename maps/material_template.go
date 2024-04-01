package maps


// 创建公共变量
var (
	// MaterialTemplateIdxStatus 创建一个公共的MaterialTemplateIdxStatus变量，其类型为MaterialTemplateStatusMap
	MaterialTemplateIdxStatus = CommonUintStringMap{
		0: `创建中`,
		1: `已完成`,
		2: `中止失败`,
	}

	// MaterialTemplateStatus 创建一个公共的MaterialTemplateStatus变量，其类型为MaterialTemplateStatusMap
	MaterialTemplateStatus = CommonUintStringMap{
		0: `待审核`,
		1: `正常`,
		2: `下架`,
	}

)

