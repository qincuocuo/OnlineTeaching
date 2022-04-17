package basic

import (
	"strings"
)

var BasicIgnoreUrlList = []string{
	"/",
	"/auth/login",
	"/auth/verifycode",
}

// CheckURL 判定路径是否存在于配置文件的忽略列表中
func CheckURL(path string) bool {
	for _, item := range BasicIgnoreUrlList {
		if path == item || strings.Contains(path, "static") {
			return true
		}
	}
	return false
}
