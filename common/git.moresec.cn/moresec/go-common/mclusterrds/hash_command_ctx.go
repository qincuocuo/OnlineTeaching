package mclusterrds

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func (rc *RedisCli) HDelWithCtx(ctx context.Context, key string, field ...string) (int64, error) {
	return rc.redisCli.HDel(ctx, key, field...).Result()
}

func (rc *RedisCli) HExistsWithCtx(ctx context.Context, key, field string) (bool, error) {
	return rc.redisCli.HExists(ctx, key, field).Result()
}

func (rc *RedisCli) HGetWithCtx(ctx context.Context, key, field string) (string, bool) {
	rsp, err := rc.redisCli.HGet(ctx, key, field).Result()
	if err != nil || err == redis.Nil {
		return "", false
	}
	return rsp, true
}

func (rc *RedisCli) HGetAllWithCtx(ctx context.Context, key string) (map[string]string, error) {
	return rc.redisCli.HGetAll(ctx, key).Result()
}

func (rc *RedisCli) HIncrByWithCtx(ctx context.Context, key, field string, score int64) (int64, error) {
	return rc.redisCli.HIncrBy(ctx, key, field, score).Result()
}

func (rc *RedisCli) HIncrByFloatWithCtx(ctx context.Context, key, field string, score float64) (float64, error) {
	return rc.redisCli.HIncrByFloat(ctx, key, field, score).Result()
}

func (rc *RedisCli) HKeysWithCtx(ctx context.Context, key string) ([]string, error) {
	return rc.redisCli.HKeys(ctx, key).Result()
}

func (rc *RedisCli) HLenWithCtx(ctx context.Context, key string) (int64, error) {
	return rc.redisCli.HLen(ctx, key).Result()
}

func (rc *RedisCli) HMSetWithCtx(ctx context.Context, key string, param map[string]interface{}) (bool, error) {
	rsp, err := rc.redisCli.HSet(ctx, key, param).Result()
	if err == nil && rsp == int64(len(param)) {
		return true, nil
	}
	return false, err
}

func (rc *RedisCli) HSetWithCtx(ctx context.Context, key, field string, value interface{}) (bool, error) {
	num, err := rc.redisCli.HSet(ctx, key, field, value).Result()
	if err != nil {
		return false, err
	}

	return num == 1, nil
}

func (rc *RedisCli) HSetNXWithCtx(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return rc.redisCli.HSetNX(ctx, key, field, value).Result()
}

func (rc *RedisCli) HValsWithCtx(ctx context.Context, key string) ([]string, error) {
	return rc.redisCli.HVals(ctx, key).Result()
}
