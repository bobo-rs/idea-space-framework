package maps

var (
	// UserStatus 用户状态MAP
	UserStatus = CommonUintStringMap{
		0: `正常`,
		1: `已冻结`,
		2: `待注销`,
	}

	// UserCredentialType 证件类型
	UserCredentialType = CommonUintStringMap{
		0: `身份证`,
	}

	// UserCredentialExpireType 用户证件效期类型
	UserCredentialExpireType = CommonUintStringMap{
		0: `有效期`,
		1: `长期`,
	}

	// UserDeviceDisabled 用户设备禁用
	UserDeviceDisabled = CommonUintStringMap{
		0: `正常`,
		1: `禁用`,
	}

	// UserLoginStatus 用户登录态状态
	UserLoginStatus = CommonUintStringMap{
		0: `正常`,
		1: `下线`,
		2: `超时下线`,
	}

	// UserLoginIsUse 用户登录是否使用
	UserLoginIsUse = CommonUintStringMap{
		0: `是`,
		1: `否`,
	}

	// UserLoginIsNewDevice 用户登录是否新设备
	UserLoginIsNewDevice = CommonUintStringMap{
		0: `常用设备`,
		1: `新设备`,
	}
)



