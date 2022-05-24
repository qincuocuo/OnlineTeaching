package mclusterrds

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func (rc *RedisCli) ZAdd(key string, score float64, member string) (int64, error) {
	return rc.redisCli.ZAdd(context.TODO(), key, &redis.Z{Score: score, Member: member}).Result()
}

func (rc *RedisCli) ZScore(key, member string) (float64, error) {
	return rc.redisCli.ZScore(context.TODO(), key, member).Result()
}

func (rc *RedisCli) ZRem(key string, member ...interface{}) (int64, error) {
	return rc.redisCli.ZRem(context.TODO(), key, member...).Result()
}

func (rc *RedisCli) ZRemRangeByRank(key string, start, end int) (int64, error) {
	return rc.redisCli.ZRemRangeByRank(context.TODO(), key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ZCard(key string) (int64, error) {
	return rc.redisCli.ZCard(context.TODO(), key).Result()
}

func (rc *RedisCli) ZRank(key, member string) (int64, error) {
	ret, err := rc.redisCli.ZRank(context.TODO(), key, member).Result()
	if err == redis.Nil {
		return -1, err
	}
	return ret, err
}

func (rc *RedisCli) ZRevRank(key, member string) (int64, error) {
	return rc.redisCli.ZRevRank(context.TODO(), key, member).Result()
}

func (rc *RedisCli) ZRange(key string, start, end int) ([]string, error) {
	return rc.redisCli.ZRange(context.TODO(), key, int64(start), int64(end)).Result()
}

//TODO:  测试
func (rc *RedisCli) ZRangeByScore(key string, min, max float64) ([]string, error) {
	return rc.redisCli.ZRangeByScore(context.TODO(), key, &redis.ZRangeBy{
		Min: strconv.FormatFloat(min, 'f', 6, 64),
		Max: strconv.FormatFloat(max, 'f', 6, 64)}).Result()
}

func (rc *RedisCli) ZRevRange(key string, start, end int) ([]string, error) {
	return rc.redisCli.ZRevRange(context.TODO(), key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ZRevRangeWithScores(key string, start, end int) ([]string, []string, error) {
	result, err := rc.redisCli.ZRevRangeWithScores(context.TODO(), key, int64(start), int64(end)).Result()
	if err != nil {
		return nil, nil, err
	}
	keys := make([]string, len(result))
	value := make([]string, len(result))
	for i, v := range result {
		keys[i] = v.Member.(string)
		value[i] = strconv.FormatFloat(v.Score, 'f', 6, 64)
	}
	return keys, value, err
}

func (rc *RedisCli) ZIncrBy(key string, score float64, member string) (float64, error) {
	return rc.redisCli.ZIncrBy(context.TODO(), key, score, member).Result()
}

type ZMemInfo struct {
	Member string
	Score  float64
}

func (rc *RedisCli) ZRangeWithScores(key string, start, end int) ([]ZMemInfo, error) {
	result, err := rc.redisCli.ZRangeWithScores(context.TODO(), key, int64(start), int64(end)).Result()
	if err != nil {
		return nil, err
	}
	memInfo := make([]ZMemInfo, len(result))
	for i, v := range result {
		memInfo[i] = ZMemInfo{Member: v.Member.(string), Score: v.Score}
	}
	return memInfo, nil
}

func (rc *RedisCli) ZRemRangeByScore(key string, min, max float64) (int64, error) {
	return rc.redisCli.ZRemRangeByScore(context.TODO(), key, strconv.FormatFloat(min, 'f', 6, 64),
		strconv.FormatFloat(max, 'f', 6, 64)).Result()
}

func (rc *RedisCli) ZCount(key string, min, max float64) (int64, error) {
	return rc.redisCli.ZCount(context.TODO(), key, strconv.FormatFloat(min, 'f', 6, 64),
		strconv.FormatFloat(max, 'f', 6, 64)).Result()
}
