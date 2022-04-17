package redis

import (
	"context"
	"fmt"
	"strconv"
	"webapi/internal/cache"
	"webapi/middleware/tracking"
	"webapi/utils"
)

func SetUserWebToken(ctx context.Context, uid int, token string) (err error) {
	span, _ := tracking.RedisTracking(ctx, cache.RdxWebToken(uid), fmt.Sprintf("token: %s", token))
	defer span.End()
	_, err = cache.RedisCli.Set(cache.RdxWebToken(uid), token)
	return
}

func GetUserWebToken(ctx context.Context, uid int) (token string, err error) {
	span, _ := tracking.RedisTracking(ctx, cache.RdxWebToken(uid))
	defer span.End()
	token, err = cache.RedisCli.Get(cache.RdxWebToken(uid))
	return
}

func UserWebTokenIsExists(ctx context.Context, uid int) (isExist bool, err error) {
	span, _ := tracking.RedisTracking(ctx, cache.RdxWebToken(uid))
	defer span.End()
	isExist, err = cache.RedisCli.Exists(cache.RdxWebToken(uid))
	return
}

func SetUserLoginLock(addr, username string, lockTm int32) {
	var count int
	countStr, _ := cache.RedisCli.Get(cache.RdxUserLock(addr, username))
	count, _ = strconv.Atoi(countStr)
	count++
	_, _ = cache.RedisCli.SetEx(cache.RdxUserLock(addr, username), count, lockTm*60)
}

func GetUserLoginLock(addr, username string) (count int) {
	countStr, _ := cache.RedisCli.Get(cache.RdxUserLock(addr, username))
	count, _ = strconv.Atoi(countStr)
	return
}
func RemoveUserLoginLock(addr, username string) {
	_ = cache.RedisCli.Del(cache.RdxUserLock(addr, username))
}

func SetUserPasswordCheckLock(addr string, userRole int32, lockTm int32) {
	var count int
	countStr, _ := cache.RedisCli.Get(cache.RdxPasswordCheckLock(addr, utils.String.Int32ToString(userRole)))
	count, _ = strconv.Atoi(countStr)
	count++
	_, _ = cache.RedisCli.SetEx(cache.RdxPasswordCheckLock(addr, utils.String.Int32ToString(userRole)), count, lockTm*60)
}
func GetUserPasswordCheckLock(addr, username string) (count int) {
	countStr, _ := cache.RedisCli.Get(cache.RdxPasswordCheckLock(addr, username))
	count, _ = strconv.Atoi(countStr)
	return
}
