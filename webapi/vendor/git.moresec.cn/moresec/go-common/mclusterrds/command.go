package mclusterrds

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//Append 实现 APPEND 命令
func (rc *RedisCli) Append(key, value string) error {
	return rc.redisCli.Append(context.TODO(), key, value).Err()
}

//Decr 实现DECR key
func (rc *RedisCli) Decr(key string) (int64, error) {
	return rc.redisCli.Decr(context.TODO(), key).Result()
}

//Decrby 实现DECRBY
func (rc *RedisCli) DecrBy(key string, value int64) (int64, error) {
	return rc.redisCli.DecrBy(context.TODO(), key, value).Result()
}

//Get 实现GET
func (rc *RedisCli) Get(key string) (string, error) {
	return rc.redisCli.Get(context.TODO(), key).Result()
}

func (rc *RedisCli) GetInt(key string) (int, error) {
	return rc.redisCli.Get(context.TODO(), key).Int()
}

func (rc *RedisCli) GetInt64(key string) (int64, error) {
	return rc.redisCli.Get(context.TODO(), key).Int64()
}

func (rc *RedisCli) GetUInt64(key string) (uint64, error) {
	return rc.redisCli.Get(context.TODO(), key).Uint64()
}

func (rc *RedisCli) GetBool(key string) (bool, error) {
	return rc.redisCli.Get(context.TODO(), key).Bool()
}

func (rc *RedisCli) Incr(key string) (int64, error) {
	return rc.redisCli.Incr(context.TODO(), key).Result()
}

func (rc *RedisCli) IncrBy(key string, value int64) (int64, error) {
	return rc.redisCli.IncrBy(context.TODO(), key, value).Result()
}

func (rc *RedisCli) IncrbyFloat(key string, value float64) (float64, error) {
	return rc.redisCli.IncrByFloat(context.TODO(), key, value).Result()
}

func (rc *RedisCli) Set(key string, value interface{}) (int64, error) {
	rsp, err := rc.redisCli.Set(context.TODO(), key, value, 0).Result()
	if err == nil && rsp == "OK" {
		return 1, nil
	}
	return 0, err
}

func (rc *RedisCli) SetEx(key string, value interface{}, second int32) (int64, error) {
	rsp, err := rc.redisCli.Set(context.TODO(), key, value, time.Duration(second)*time.Second).Result()
	if err == nil && rsp == "OK" {
		return 1, nil
	}
	return 0, err
}

func (rc *RedisCli) SetNX(key string, value interface{}, expire time.Duration) (bool, error) {
	return rc.redisCli.SetNX(context.TODO(), key, value, expire).Result()
}

/// key 操作.
func (rc *RedisCli) Del(key string) error {
	return rc.redisCli.Del(context.TODO(), key).Err()
}

func (rc *RedisCli) Exists(key ...string) (bool, error) {
	rsp, err := rc.redisCli.Exists(context.TODO(), key...).Result()
	if err == nil {
		return rsp == 1, err
	}
	return false, err
}

func (rc *RedisCli) Expire(key string, second int32) error {
	_, err := rc.redisCli.Expire(context.TODO(), key, time.Duration(second)*time.Second).Result()
	return err
}

func (rc *RedisCli) TTL(key string) (int64, error) {
	rsp, err := rc.redisCli.TTL(context.TODO(), key).Result()
	if err == nil {
		return int64(rsp.Seconds()), nil
	}
	if err == redis.Nil {
		return -2, err
	}
	return -1, err
}

func (rc *RedisCli) ExpireAt(key string, timestamp int64) error {
	_, err := rc.redisCli.ExpireAt(context.TODO(), key, time.Unix(timestamp, 0)).Result()
	return err
}

// keys
func (rc *RedisCli) Keys(pattern string) ([]string, error) {
	return rc.redisCli.Keys(context.TODO(), pattern).Result()
}

func (rc *RedisCli) RawCli() redis.Cmdable {
	return rc.redisCli
}
