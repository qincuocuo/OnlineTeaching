package mnet

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	aes := &Aes{cipher: []byte("0123456789abcdef")}

	data, _ := aes.Encrypt("0123456789")
	str := string(data)
	fmt.Println("encrypt:", data)
	fmt.Println("len:", len(str))
}
