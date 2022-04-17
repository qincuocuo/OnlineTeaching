package cache

import "fmt"

// RdxJwtWhitelist jwt白名单(用于超时登录)
//类型: string
func RdxJwtWhitelist(token string) string {
	return "jwt:whitelist:" + token
}

// RdxJwtBlacklist jwt黑名单
//类型: string
func RdxJwtBlacklist(token string) string {
	return "jwt:blacklist:" + token
}

func RdxWebToken(uid int) string {
	return fmt.Sprintf("web:token:%d", uid)
}

func RdxCaptchaKey(id string) string {
	return "captcha:" + id
}

// RdxUserLock 用户锁定表(用于错误次数判定)
//类型: string
func RdxUserLock(addr, username string) string {
	return "user:lock:" + addr + ":" + username
}

// RdxPasswordCheckLock 用户密码校验锁定表(用于错误次数判定)
func RdxPasswordCheckLock(addr, username string) string {
	return "password:check:lock:" + addr + ":" + username
}

// RdxImageProc 镜像对应的进程信息.
// 类型: set
func RdxImageProc(image string) string {
	return "proc:image:" + image
}
