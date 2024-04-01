package enums

type (
	// SexEnums 枚举性别类型
	SexEnums uint
)

const (
	SexMan SexEnums = iota
	SexWoman
	SexSecrecy
)

// Format 性别格式化
func (e SexEnums) Format() string {
	switch e {
	case SexMan:
		return `男`
	case SexWoman:
		return `女`
	default:
		return `保密`
	}
}