package mclusterrds

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func (rc *RedisCli) ZAddWithCtx(ctx context.Context, key string, score float64, member string) (int64, error) {
	return rc.redisCli.ZAdd(ctx, key, &redis.Z{Score: score, Member: member}).Result()
}

func (rc *RedisCli) ZScoreWithCtx(ctx context.Context, key, member string) (float64, error) {
	return rc.redisCli.ZScore(ctx, key, member).Result()
}

func (rc *RedisCli) ZRemWithCtx(ctx context.Context, key string, member ...interface{}) (int64, error) {
	return rc.redisCli.ZRem(ctx, key, member...).Result()
}

func (rc *RedisCli) ZRemRangeByRankWithCtx(ctx context.Context, key string, start, end int) (int64, error) {
	return rc.redisCli.ZRemRangeByRank(ctx, key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ZCardWithCtx(ctx context.Context, key string) (int64, error) {
	return rc.redisCli.ZCard(ctx, key).Result()
}

func (rc *RedisCli) ZRankWithCtx(ctx context.Context, key, member string) (int64, error) {
	ret, err := rc.redisCli.ZRank(ctx, key, member).Result()
	if err == redis.Nil {
		return -1, err
	}
	return ret, err
}

func (rc *RedisCli) ZRevRankWithCtx(ctx context.Context, key, member string) (int64, error) {
	return rc.redisCli.ZRevRank(ctx, key, member).Result()
}

func (rc *RedisCli) ZRangeWithCtx(ctx context.Context, key string, start, end int) ([]string, error) {
	return rc.redisCli.ZRange(ctx, key, int64(start), int64(end)).Result()
}

//TODO:  测试
func (rc *RedisCli) ZRangeByScoreWithCtx(ctx context.Context, key string, min, max float64) ([]string, error) {
	return rc.redisCli.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: strconv.FormatFloat(min, 'f', 6, 64),
		Max: strconv.FormatFloat(max, 'f', 6, 64)}).Result()
}

func (rc *RedisCli) ZRevRangeWithCtx(ctx context.Context, key string, start, end int) ([]string, error) {
	return rc.redisCli.ZRevRange(ctx, key, int64(start), int64(end)).Result()
}

func (rc *RedisCli) ZRevRangeWithScoresAndCtx(ctx context.Context, key string, start, end int) ([]string, []string, error) {
	result, err := rc.redisCli.ZRevRangeWithScores(ctx, key, int64(start), int64(end)).Result()
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

func (rc *RedisCli) ZIncrByWithCtx(ctx context.Context, key string, score float64, member string) (float64, error) {
	return rc.redisCli.ZIncrBy(ctx, key, score, member).Result()
}

func (rc *RedisCli) ZRangeWithScoresAndCtx(ctx context.Context, key string, start, end int) ([]ZMemInfo, error) {
	result, err := rc.redisCli.ZRangeWithScores(ctx, key, int64(start), int64(end)).Result()
	if err != nil {
		return nil, err
	}
	memInfo := make([]ZMemInfo, len(result))
	for i, v := range result {
		memInfo[i] = ZMemInfo{Member: v.Member.(string), Score: v.Score}
	}
	return memInfo, nil
}

func (rc *RedisCli) ZRemRangeByScoreWithCtx(ctx context.Context, key string, min, max float64) (int64, error) {
	return rc.redisCli.ZRemRangeByScore(ctx, key, strconv.FormatFloat(min, 'f', 6, 64),
		strconv.FormatFloat(max, 'f', 6, 64)).Result()
}

func (rc *RedisCli) ZCountWithCtx(ctx context.Context, key string, min, max float64) (int64, error) {
	return rc.redisCli.ZCount(ctx, key, strconv.FormatFloat(min, 'f', 6, 64),
		strconv.FormatFloat(max, 'f', 6, 64)).Result()
}
