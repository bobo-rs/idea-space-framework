package utils

// NewArray 实例数组
func NewArray(data interface{}) *sArray {
	return &sArray{
		Data: data,
	}
}
