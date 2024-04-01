package enums

type (
	// SpaceImagePrecisionEnums 图片空间精度枚举类型
	SpaceImagePrecisionEnums uint
)

const (
	SpaceImagePrecisionZero SpaceImagePrecisionEnums = iota
	SpaceImagePrecisionOne
	SpaceImagePrecisionTwo
	SpaceImagePrecisionThree
	SpaceImagePrecisionFour
	SpaceImagePrecisionFive
	SpaceImagePrecisionSix
	SpaceImagePrecisionSeven
	SpaceImagePrecisionEight
	SpaceImagePrecisionNine
	SpaceImagePrecisionTen
)

// Format 格式化图片空间精度
func (e SpaceImagePrecisionEnums) Format() string {
	switch e {
	case SpaceImagePrecisionZero,SpaceImagePrecisionOne, SpaceImagePrecisionTwo:
		return `低精度`
	case SpaceImagePrecisionThree, SpaceImagePrecisionFour, SpaceImagePrecisionFive:
		return `中精度`
	case SpaceImagePrecisionTen:
		return `超高精度`
	default:
		return `高精度`
	}
}
