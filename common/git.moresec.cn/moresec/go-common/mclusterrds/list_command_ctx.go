package mclusterrds

import (
	"context"
	"time"
)

func (rc *RedisCli) ListIndexWithCtx(ctx context.Context, key string, index int64) (string, error) {
	return rc.redisCli.LIndex(ctx, key, index).Result()
}

func (rc *RedisCli) ListLenWithCtx(ctx context.Context, key string) (int64, error) {
	return rc.redisCli.LLen(ctx, key).Result()
}

func (rc *RedisCli) ListLPopWithCtx(ctx context.Context, key string) (string, error) {
	return rc.redisCli.LPop(ctx, key).Result()
}

func (rc *RedisCli) ListLPushWithCtx(ctx context.Context, key string, value ...interface{}) (int64, error) {
	return rc.redisCli.LPush(ctx, key, value...).Result()
}

func (rc *RedisCli) ListLRangeWithCtx(ctx context.Context, key string, start, end int) ([]string, error) {
	return rc.redisCli.LRange(ctx, key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ListLSetWithCtx(ctx context.Context, key string, index int, value interface{}) (string, error) {
	return rc.redisCli.LSet(ctx, key, int64(index), value).Result()
}

func (rc *RedisCli) ListLTrimWithCtx(ctx context.Context, key string, start, end int) (string, error) {
	return rc.redisCli.LTrim(ctx, key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ListRPopWithCtx(ctx context.Context, key string) (string, error) {
	return rc.redisCli.RPop(ctx, key).Result()
}

func (rc *RedisCli) ListRPushWithCtx(ctx context.Context, key string, value ...interface{}) (int64, error) {
	return rc.redisCli.RPush(ctx, key, value...).Result()
}

func (rc *RedisCli) ListBRPopWithCtx(ctx context.Context, key string, timeout time.Duration) ([]string, error) {
	return rc.redisCli.BRPop(ctx, timeout, key).Result()
}

func (rc *RedisCli) ListRPopLPushWithCtx(ctx context.Context, key string, destKey string) (string, error) {
	return rc.redisCli.RPopLPush(ctx, key, destKey).Result()
}
