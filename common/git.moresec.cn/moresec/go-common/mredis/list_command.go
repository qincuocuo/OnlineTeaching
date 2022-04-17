package mredis

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func (rc *RedisCli) ListIndex(key string, index int) (string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.String(conn.Do("LINDEX", key, strconv.Itoa(index)))
}

func (rc *RedisCli) ListLen(key string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Int64(conn.Do("LLEN", key))
}

func (rc *RedisCli) ListLPop(key string) (string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.String(conn.Do("LPOP", key))
}

func (rc *RedisCli) ListLPush(key string, value interface{}) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Int64(conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(value)...))
}

func (rc *RedisCli) ListLRange(key string, start, end int) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Strings(conn.Do("LRANGE", key, strconv.Itoa(start), strconv.Itoa(end)))
}

func (rc *RedisCli) ListLSet(key string, index int, value string) (bool, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rsp, err := redis.String(conn.Do("LSET", key, strconv.Itoa(index), value))
	if err != nil {
		return false, err
	}
	if rsp != "OK" {
		return false, errors.New("response is not ok")
	}

	return true, nil
}

func (rc *RedisCli) ListLTrim(key string, start, end int) (bool, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rsp, err := redis.String(conn.Do("LTRIM", key, strconv.Itoa(start), strconv.Itoa(end)))
	if err != nil {
		return false, err
	}
	if rsp != "OK" {
		return false, errors.New("response is not ok")
	}
	return true, nil
}

func (rc *RedisCli) ListRPop(key string) (string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.String(conn.Do("RPOP", key))
}

func (rc *RedisCli) ListRPush(key string, value interface{}) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("RPUSH", redis.Args{}.Add(key).AddFlat(value)...))
}

func (rc *RedisCli) ListBRPop(key string, timeout time.Duration) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Strings(redis.DoWithTimeout(conn, timeout, "BRPOP", key, timeout.Seconds()))
}

func (rc *RedisCli) ListRPopLPush(key string, destKey string) (string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.String(conn.Do("RPOPLPUSH", key, destKey))
}
