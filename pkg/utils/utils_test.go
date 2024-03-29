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

