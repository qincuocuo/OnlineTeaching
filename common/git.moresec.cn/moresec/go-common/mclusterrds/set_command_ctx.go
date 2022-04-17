package mclusterrds

import "context"

func (rc *RedisCli) SAddWithCtx(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return rc.redisCli.SAdd(ctx, key, members...).Result()
}

func (rc *RedisCli) SCardWithCtx(ctx context.Context, key string) (int64, error) {
	return rc.redisCli.SCard(ctx, key).Result()
}

func (rc *RedisCli) SIsMemberWithCtx(ctx context.Context, key, value string) (bool, error) {
	return rc.redisCli.SIsMember(ctx, key, value).Result()
}

func (rc *RedisCli) SMembersWithCtx(ctx context.Context, key string) ([]string, error) {
	return rc.redisCli.SMembers(ctx, key).Result()
}

func (rc *RedisCli) SPopWithCtx(ctx context.Context, key string) (string, error) {
	return rc.redisCli.SPop(ctx, key).Result()
}

func (rc *RedisCli) SRemWithCtx(ctx context.Context, key string, member ...interface{}) (int64, error) {
	return rc.redisCli.SRem(ctx, key, member...).Result()
}

func (rc *RedisCli) SUnionWithCtx(ctx context.Context, key ...string) ([]string, error) {
	return rc.redisCli.SUnion(ctx, key...).Result()
}
