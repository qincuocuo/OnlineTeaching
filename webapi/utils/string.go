package utils

import (
	"strconv"
	"strings"

	"git.moresec.cn/moresec/go-common/mbase"
	randomId "github.com/satori/go.uuid"
)

type str struct {
}

var String str

func (str) Compare(str1, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}

func (str) GetRandomString(len int) string {
	return mbase.MD5String(randomId.NewV1().String(), len)
}

func (str) Int32ToString(src int32) (dst string) {
	dst = strconv.Itoa(int(src))
	return dst
}
