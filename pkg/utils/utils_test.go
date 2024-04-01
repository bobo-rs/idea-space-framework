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

func TestArrayColumns(t *testing.T) {
	slices := make([]map[string]interface{}, 0)
	slices = append(slices, map[string]interface{}{
		`id`: 4,
		`name`: `哈哈哈哈`,
	})
	fmt.Println(NewArray(slices).ArrayColumns( "id"))

	sMap, _ := NewArray(slices).ColumnMap( `name`)
	sMap1, _ := NewArray(slices).ColumnMap( `id`)
	fmt.Println(sMap, sMap[`哈哈哈哈`])
	fmt.Println(sMap1, sMap1[4])
}
