package mredis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func (rc *RedisCli) ZAdd(key string, score float64, member string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	floatStr := strconv.FormatFloat(score, 'f', 6, 64)
	return redis.Int64(conn.Do("ZADD", key, floatStr, member))
}

func (rc *RedisCli) ZScore(key, member string) (float64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rst, err := redis.String(conn.Do("ZSCORE", key, member))
	if err != nil {
		return -1, err
	}

	return strconv.ParseFloat(rst, 64)
}

func (rc *RedisCli) ZRem(key, member string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("ZREM", key, member))
}

func (rc *RedisCli) ZRemRangeByRank(key string, start, end int) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Int64(conn.Do("ZREMRANGEBYRANK", key, start, end))
}

func (rc *RedisCli) ZCard(key string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Int64(conn.Do("ZCARD", key))
}

func (rc *RedisCli) ZRank(key, member string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rst, err := redis.Int64(conn.Do("ZRANK", key, member))
	if err != nil {
		return -1, err
	}
	return rst, nil
}

func (rc *RedisCli) ZRevRank(key, member string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rst, err := redis.Int64(conn.Do("ZREVRANK", key, member))
	if err != nil {
		return -1, err
	}
	return rst, nil
}

func (rc *RedisCli) ZRange(key string, start, end int) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Strings(conn.Do("ZRANGE", key, strconv.Itoa(start), strconv.Itoa(end)))
}

func (rc *RedisCli) ZRangeByScore(key string, min, max float64) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Strings(conn.Do("ZRANGEBYSCORE", key, min, max))
}

func (rc *RedisCli) ZRevRange(key string, start, end int) ([]string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	result, err := redis.Strings(conn.Do("ZREVRANGE", key, strconv.Itoa(start), strconv.Itoa(end)))
	if err != nil {
		return nil, err
	}
	return result, err
}

func (rc *RedisCli) ZRevRangeWithScores(key string, start, end int) ([]string, []string, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	result, err := redis.Strings(conn.Do("ZREVRANGE", key, strconv.Itoa(start), strconv.Itoa(end), "withscores"))
	if err != nil {
		return nil, nil, err
	}
	var keys []string
	var scores []string

	for i := 0; i < len(result); i += 2 {
		keys = append(keys, result[i])
		scores = append(scores, result[i+1])
	}
	return keys, scores, err
}

func (rc *RedisCli) ZIncrBy(key string, score float64, member string) (float64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	floatStr := strconv.FormatFloat(score, 'f', 6, 64)

	rst, err := redis.String(conn.Do("ZINCRBY", key, floatStr, member))
	if err != nil {
		return -1, err
	}

	return strconv.ParseFloat(rst, 64)
}

type ZMemInfo struct {
	Member string
	Score  float64
}

func (rc *RedisCli) ZRangeWithScores(key string, start, end int) ([]ZMemInfo, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	rst, err := redis.Strings(conn.Do("ZRANGE", key, strconv.Itoa(start), strconv.Itoa(end), "withscores"))
	if err != nil {
		return nil, err
	}

	if len(rst)%2 != 0 {
		return nil, errors.New("response is not even number")
	}

	var memInfo []ZMemInfo

	for i := 0; i < len(rst); i += 2 {
		member := rst[i]
		score, _ := strconv.ParseFloat(rst[i+1], 64)

		memInfo = append(memInfo, ZMemInfo{Member: member, Score: score})
	}

	return memInfo, nil
}

func (rc *RedisCli) ZRemRangeByScore(key string, min, max float64) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("ZREMRANGEBYSCORE", key, min, max))
}

func (rc *RedisCli) ZCount(key string, min, max float64) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()

	return redis.Int64(conn.Do("ZCOUNT", key, min, max))
}
