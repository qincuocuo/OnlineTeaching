package basic

import (
	"strings"

	"webapi/config"
)

var BasicIgnoreUrlList = []string{
	"/",
	"/auth/login",
	"/auth/verifycode",
	"/auth/register",
}

func RegisterIgnoreURLs(urls []string) {
	BasicIgnoreUrlList = config.IrisConf.Web.IgnoreUrls
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
