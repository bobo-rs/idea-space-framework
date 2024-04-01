package maps

var (
	// MaterialTypeMap 素材配图类型MAP
	MaterialTypeMap = CommonUintStringMap{
		0: `素材配图`,
		1: `素材主图`,
	}
	// MaterialPositionMap 素材卖点位置
	MaterialPositionMap = CommonStringStringMap{
		`main`: `主图`,
		`market`: `营销活动卖点`,
		`top`: `顶部卖点`,
		`btm-main`: `底部主卖点`,
		`btm-second`: `底部次卖点`,
		`btm-single`: `底部单独卖点`,
		`detail`: `详情图卖点`,
	}
	// MaterialStatusMap 素材状态MAP
	MaterialStatusMap = CommonUintStringMap{
		0: `正常`,
		1: `废弃`,
	}
)

