package mclusterrds

import (
	"context"
	"github.com/go-redis/redis/v8"
)

///// hash操作.
func (rc *RedisCli) HDel(key string, field ...string) (int64, error) {
	return rc.redisCli.HDel(context.TODO(), key, field...).Result()
}

func (rc *RedisCli) HExists(key, field string) (bool, error) {
	return rc.redisCli.HExists(context.TODO(), key, field).Result()
}

func (rc *RedisCli) HGet(key, field string) (string, bool) {
	rsp, err := rc.redisCli.HGet(context.TODO(), key, field).Result()
	if err != nil || err == redis.Nil {
		return "", false
	}
	return rsp, true
}

func (rc *RedisCli) HGetAll(key string) (map[string]string, error) {
	return rc.redisCli.HGetAll(context.TODO(), key).Result()
}

func (rc *RedisCli) HIncrBy(key, field string, score int64) (int64, error) {
	return rc.redisCli.HIncrBy(context.TODO(), key, field, score).Result()
}

func (rc *RedisCli) HIncrByFloat(key, field string, score float64) (float64, error) {
	return rc.redisCli.HIncrByFloat(context.TODO(), key, field, score).Result()
}

func (rc *RedisCli) HKeys(key string) ([]string, error) {
	return rc.redisCli.HKeys(context.TODO(), key).Result()
}

func (rc *RedisCli) HLen(key string) (int64, error) {
	return rc.redisCli.HLen(context.TODO(), key).Result()
}

func (rc *RedisCli) HMSet(key string, param map[string]interface{}) (bool, error) {
	rsp, err := rc.redisCli.HSet(context.TODO(), key, param).Result()
	if err == nil && rsp == int64(len(param)) {
		return true, nil
	}
	return false, err
}

func (rc *RedisCli) HSet(key, field string, value interface{}) (bool, error) {
	num, err := rc.redisCli.HSet(context.TODO(), key, field, value).Result()
	if err != nil {
		return false, err
	}

	return num == 1, nil
}

func (rc *RedisCli) HSetNX(key, field string, value interface{}) (bool, error) {
	return rc.redisCli.HSetNX(context.TODO(), key, field, value).Result()
}

func (rc *RedisCli) HVals(key string) ([]string, error) {
	return rc.redisCli.HVals(context.TODO(), key).Result()
}
