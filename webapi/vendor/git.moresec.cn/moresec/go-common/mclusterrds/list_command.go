package mclusterrds

import (
	"context"
	"time"
)

func (rc *RedisCli) ListIndex(key string, index int64) (string, error) {
	return rc.redisCli.LIndex(context.TODO(), key, index).Result()
}

func (rc *RedisCli) ListLen(key string) (int64, error) {
	return rc.redisCli.LLen(context.TODO(), key).Result()
}

func (rc *RedisCli) ListLPop(key string) (string, error) {
	return rc.redisCli.LPop(context.TODO(), key).Result()
}

func (rc *RedisCli) ListLPush(key string, value ...interface{}) (int64, error) {
	return rc.redisCli.LPush(context.TODO(), key, value...).Result()
}

func (rc *RedisCli) ListLRange(key string, start, end int) ([]string, error) {
	return rc.redisCli.LRange(context.TODO(), key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ListLSet(key string, index int, value interface{}) (string, error) {
	return rc.redisCli.LSet(context.TODO(), key, int64(index), value).Result()
}

func (rc *RedisCli) ListLTrim(key string, start, end int) (string, error) {
	return rc.redisCli.LTrim(context.TODO(), key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ListRPop(key string) (string, error) {
	return rc.redisCli.RPop(context.TODO(), key).Result()
}

func (rc *RedisCli) ListRPush(key string, value ...interface{}) (int64, error) {
	return rc.redisCli.RPush(context.TODO(), key, value...).Result()
}

func (rc *RedisCli) ListBRPop(key string, timeout time.Duration) ([]string, error) {
	return rc.redisCli.BRPop(context.TODO(), timeout, key).Result()
}

func (rc *RedisCli) ListRPopLPush(key string, destKey string) (string, error) {
	return rc.redisCli.RPopLPush(context.TODO(), key, destKey).Result()
}
