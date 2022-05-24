package mclusterrds

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func (rc *RedisCli) AppendWithCtx(ctx context.Context, key, value string) error {
	return rc.redisCli.Append(ctx, key, value).Err()
}

//Decr 实现DECR key
func (rc *RedisCli) DecrWithCtx(ctx context.Context, key string) (int64, error) {
	return rc.redisCli.Decr(ctx, key).Result()
}

//Decrby 实现DECRBY
func (rc *RedisCli) DecrByWithCtx(ctx context.Context, key string, value int64) (int64, error) {
	return rc.redisCli.DecrBy(ctx, key, value).Result()
}

//Get 实现GET
func (rc *RedisCli) GetWithCtx(ctx context.Context, key string) (string, error) {
	return rc.redisCli.Get(ctx, key).Result()
}

func (rc *RedisCli) IncrWithCtx(ctx context.Context, key string) (int64, error) {
	return rc.redisCli.Incr(ctx, key).Result()
}

func (rc *RedisCli) IncrByWithCtx(ctx context.Context, key string, value int64) (int64, error) {
	return rc.redisCli.IncrBy(ctx, key, value).Result()
}

func (rc *RedisCli) IncrbyFloatWithCtx(ctx context.Context, key string, value float64) (float64, error) {
	return rc.redisCli.IncrByFloat(ctx, key, value).Result()
}

func (rc *RedisCli) SetWithCtx(ctx context.Context, key string, value interface{}) (int64, error) {
	rsp, err := rc.redisCli.Set(ctx, key, value, 0).Result()
	fmt.Println(rsp, " == ", err)
	if err == nil && rsp == "OK" {
		return 1, nil
	}
	return 0, err
}

func (rc *RedisCli) SetExWithCtx(ctx context.Context, key string, value interface{}, second int32) (int64, error) {
	rsp, err := rc.redisCli.Set(ctx, key, value, time.Duration(second)*time.Second).Result()
	if err == nil && rsp == "OK" {
		return 1, nil
	}
	return 0, err
}

func (rc *RedisCli) SetNXWithCtx(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error) {
	return rc.redisCli.SetNX(ctx, key, value, expire).Result()
}

/// key 操作.
func (rc *RedisCli) DelWithCtx(ctx context.Context, key string) error {
	return rc.redisCli.Del(ctx, key).Err()
}

func (rc *RedisCli) ExistsWithCtx(ctx context.Context, key ...string) (bool, error) {
	rsp, err := rc.redisCli.Exists(ctx, key...).Result()
	if err == nil {
		return rsp == 1, err
	}
	return false, err
}

func (rc *RedisCli) ExpireWithCtx(ctx context.Context, key string, second int32) error {
	_, err := rc.redisCli.Expire(ctx, key, time.Duration(second)*time.Second).Result()
	return err
}

func (rc *RedisCli) TTLWithCtx(ctx context.Context, key string) (int64, error) {
	rsp, err := rc.redisCli.TTL(ctx, key).Result()
	if err == nil {
		return int64(rsp.Seconds()), nil
	}
	if err == redis.Nil {
		return -2, err
	}
	return -1, err
}

func (rc *RedisCli) ExpireAtWithCtx(ctx context.Context, key string, timestamp int64) error {
	_, err := rc.redisCli.ExpireAt(ctx, key, time.Unix(timestamp, 0)).Result()
	return err
}

// keys
func (rc *RedisCli) KeysWithCtx(ctx context.Context, pattern string) ([]string, error) {
	return rc.redisCli.Keys(ctx, pattern).Result()
}
