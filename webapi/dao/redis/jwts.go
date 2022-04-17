package redis

import (
	"strconv"
	"webapi/internal/cache"
)

func InJwtBlacklist(token string) bool {
	_, err := cache.RedisCli.Get(cache.RdxJwtBlacklist(token))
	if err != nil {
		return false
	}
	return true
}

func SetJwtBlacklist(token string, expire int32) (err error) {
	_, err = cache.RedisCli.SetEx(cache.RdxJwtBlacklist(token), 1, expire)
	return
}

func RemoveJwtWhitelist(token string) (err error) {
	err = cache.RedisCli.Del(cache.RdxJwtBlacklist(token))
	return
}

func SetJwtWhitelist(token string, expire int32) (err error) {
	_, err = cache.RedisCli.SetEx(cache.RdxJwtWhitelist(token), expire, expire)
	return
}

func InJwtWhitelist(token string) bool {
	_, err := cache.RedisCli.Get(cache.RdxJwtWhitelist(token))
	if err != nil {
		return false
	}
	return true
}

func FlushJwtWhitelist(token string) (err error) {
	var expire string
	expire, err = cache.RedisCli.Get(cache.RdxJwtWhitelist(token))
	exp, _ := strconv.Atoi(expire)
	_, err = cache.RedisCli.SetEx(cache.RdxJwtWhitelist(token), exp, int32(exp))
	return
}
