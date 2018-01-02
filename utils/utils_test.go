package utils

import "testing"
import "fmt"

func TestString2Map(t *testing.T) {
	str := `realm="https://r.j3ss.co/auth",service="Docker registry",scope="registry:catalog:*"`
	m := String2Map(str)
	for k, v := range m {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func TestBase64Encode(t *testing.T) {
	o := Base64Encode("123")
	fmt.Println(o)
}
