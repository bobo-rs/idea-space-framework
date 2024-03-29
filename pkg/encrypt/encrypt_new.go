package encrypt

// NewM 实例Md5加密
func NewM(b ...byte) *Md5 {
	return &Md5{
		bytes: b,
	}
}