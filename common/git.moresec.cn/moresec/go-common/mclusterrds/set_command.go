package mclusterrds

import "context"

func (rc *RedisCli) SAdd(key string, members ...interface{}) (int64, error) {

	return rc.redisCli.SAdd(context.TODO(), key, members...).Result()
}

func (rc *RedisCli) SCard(key string) (int64, error) {

	return rc.redisCli.SCard(context.TODO(), key).Result()
}

func (rc *RedisCli) SIsMember(key, value string) (bool, error) {

	return rc.redisCli.SIsMember(context.TODO(), key, value).Result()
}

func (rc *RedisCli) SMembers(key string) ([]string, error) {

	return rc.redisCli.SMembers(context.TODO(), key).Result()
}

func (rc *RedisCli) SPop(key string) (string, error) {

	return rc.redisCli.SPop(context.TODO(), key).Result()
}

func (rc *RedisCli) SRem(key string, member ...interface{}) (int64, error) {

	return rc.redisCli.SRem(context.TODO(), key, member...).Result()
}

func (rc *RedisCli) SUnion(key ...string) ([]string, error) {

	return rc.redisCli.SUnion(context.TODO(), key...).Result()
}
