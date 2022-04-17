package mredis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

///// hash操作.
func (rc *RedisCli) HDel(key, field string) (int, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int(conn.Do("HDEL", key, field))
}

func (rc *RedisCli) HExists(key, field string) (bool, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rt, err := redis.Int(conn.Do("HEXISTS", key, field))
	if err != nil {
		return false, err
	}
	return rt == 1, nil
}

func (rc *RedisCli) HGet(key, field string) (string, bool) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rsp, err := redis.String(conn.Do("HGET", key, field))
	if err == redis.ErrNil {
		return "", false // 所查字段不存在.
	}
	if err == nil {
		return rsp, true
	}
	return "", true
}

func (rc *RedisCli) HGetAll(key string) (map[string]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.StringMap(conn.Do("HGETALL", key))
}

func (rc *RedisCli) HIncrBy(key, field string, score int) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("HINCRBY", key, field, strconv.Itoa(score)))
}

func (rc *RedisCli) HIncrByFloat(key, field string, score float32) (string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	floatStr := strconv.FormatFloat(float64(score), 'f', 6, 32)
	return redis.String(conn.Do("HINCRBYFLOAT", key, field, floatStr))
}

func (rc *RedisCli) HKeys(key string) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Strings(conn.Do("HKEYS", key))
}

func (rc *RedisCli) HLen(key string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Int64(conn.Do("HLEN", key))
}

func (rc *RedisCli) HMGet(key string, fields []string) (map[string]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	result, err := redis.Strings(conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(fields)...))
	if err != nil {
		return nil, err
	}

	if len(fields) != len(result) {
		return nil, errors.New("fields is not equal to result")
	}

	resultMap := make(map[string]string, 10)
	for i := 0; i < len(fields); i++ {
		resultMap[fields[i]] = result[i]
	}
	return resultMap, nil
}

func (rc *RedisCli) HMSet(key string, param map[string]interface{}) (bool, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rst, err := redis.String(conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(param)...))
	if err != nil {
		return false, err
	}
	if rst != "OK" {
		return false, errors.New("response is not ok")
	}
	return true, nil
}

func (rc *RedisCli) HSet(key, field string, value interface{}) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("HSET", key, field, value))
}

func (rc *RedisCli) HSetNX(key, field string, value interface{}) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Int64(conn.Do("HSETNX", key, field, value))
}

func (rc *RedisCli) HVals(key string) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Strings(conn.Do("HVALS", key))
}
