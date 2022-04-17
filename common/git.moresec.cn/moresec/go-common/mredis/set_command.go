package mredis

import "github.com/gomodule/redigo/redis"

func (rc *RedisCli) SAdd(key string, value interface{}) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("SADD", redis.Args{}.Add(key).AddFlat(value)...))
}

func (rc *RedisCli) SCard(key string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("SCARD", key))
}

func (rc *RedisCli) SIsMember(key, value string) (bool, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rst, err := redis.Int64(conn.Do("SISMEMBER", key, value))
	if err != nil {
		return false, err
	}
	return rst == 1, nil
}

func (rc *RedisCli) SMembers(key string) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Strings(conn.Do("SMEMBERS", key))
}

func (rc *RedisCli) SPop(key string) (string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	reply, err := redis.String(conn.Do("SPOP", key))
	if err == redis.ErrNil {
		return "", nil
	}
	return reply, err
}

func (rc *RedisCli) SRem(key string, member interface{}) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("SREM", redis.Args{}.Add(key).AddFlat(member)...))
}

func (rc *RedisCli) SUnion(key ...string) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Strings(conn.Do("SUNION", redis.Args{}.AddFlat(key)...))
}
