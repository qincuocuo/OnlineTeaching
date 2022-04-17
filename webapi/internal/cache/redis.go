package cache

import (
	"context"
	"strings"
	"webapi/config"

	"git.moresec.cn/moresec/go-common/mclusterrds"
)

var RedisCli *mclusterrds.RedisCli

func RedisInitPool(redis config.RedisCfg) error {
	defaultCfg := mclusterrds.NewDefaultOption()
	defaultCfg.Password = redis.RedisPasswd
	defaultCfg.MinIdleConns = 10
	defaultCfg.PoolSize = 100
	hosts := strings.Split(redis.RedisHost, ",")
	if redis.Cluster {
		defaultCfg.AddrList = hosts
		RedisCli = mclusterrds.NewClusterCli(defaultCfg)
	} else {
		defaultCfg.Addr = hosts[0]
		RedisCli = mclusterrds.NewRedisCli(defaultCfg)
	}
	if _, err := RedisCli.RawCli().Ping(context.Background()).Result(); err != nil {
		return err
	}
	return nil
}
