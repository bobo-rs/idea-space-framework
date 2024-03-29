package encrypt

import (
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"testing"
)

func TestMd5_Encrypt(t *testing.T) {
	s := `sjfhjsfhjsdhfs`
	fmt.Println(NewM([]byte(s)...).EncryptBytes())

	fmt.Println(NewM().MustEncryptString(s))

	fmt.Println(NewM().Encrypt(23))
	fmt.Println(NewM().Encrypt(123456))
	fmt.Println(gmd5.MustEncrypt(123456))
	fmt.Println(gmd5.MustEncryptString(`123456`))
	fmt.Println(NewM().EncryptString(`123456`))
}
