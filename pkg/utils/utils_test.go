package utils

import (
	"fmt"
	"testing"
)

func TestSeparatorStringToArray(t *testing.T) {
	s := `hshhs, shfsf, jjj，上课还是分开，fhsfhhf
sfhskfhksdf
shfsjdfhj	sdfh ddddd ddd  d`
	fmt.Println(SeparatorStringToArray(s), len(SeparatorStringToArray(s)))
}

func TestHashMd5(t *testing.T) {
	fmt.Println(HashMd5(`淘宝`,`12`))
}

func TestAlphaNum(t *testing.T) {
	fmt.Println(AlphaNum(`手机手机放还减肥是减肥`))

	fmt.Println(AlphaNum(`fhjsfhjkshfjs12312312`))
}
