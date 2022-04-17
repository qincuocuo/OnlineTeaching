package mnet

import (
	"fmt"
	"testing"
)

func TestNetHead(t *testing.T) {
	str := string("hello world")
	data, _ := SetPacket([]byte(str), NetHeadTypeProtoBuf)
	fmt.Println(data)
	head, _ := GetHead(data)
	fmt.Println(head)

	pkg, pkgType, _ := GetPacket(data)
	str1 := string(pkg)
	fmt.Println(str1)
	fmt.Println(pkgType)
}
