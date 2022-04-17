package mclusterrds

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type Option struct {
	// AddrList for normal redis addr.
	Addr string

	// AddrList cluster redis addr.
	AddrList []string

	Password string

	SentinelPassword string

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration
	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge time.Duration

	MasterName string

	Hooks []redis.Hook
}

//NewDefaultOption 默认参数
func NewDefaultOption() *Option {
	return &Option{
		Addr:         "localhost:6379",
		AddrList:     []string{"localhost:6379"},
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 2,
		MaxConnAge:   10 * time.Minute,
	}
}

type RedisCli struct {
	redisCli redis.Cmdable
	cluster  bool
	option   *Option // 连接参数.
	mu       sync.Mutex
}

//NewRedisCli 新建 redis 客户端.
func NewRedisCli(option *Option) *RedisCli {
	rawCli := redis.NewClient(&redis.Options{
		Addr:         option.Addr,
		Password:     option.Password,
		DialTimeout:  option.DialTimeout,
		ReadTimeout:  option.ReadTimeout,
		WriteTimeout: option.WriteTimeout,
		PoolSize:     option.PoolSize,
		MinIdleConns: option.MinIdleConns,
		MaxConnAge:   option.MaxConnAge,
	})

	for _, hook := range option.Hooks {
		rawCli.AddHook(hook)
	}

	cli := &RedisCli{
		option:   option,
		cluster:  false,
		redisCli: rawCli,
	}
	return cli
}

// NewClusterCli 新建 redis 集群客户端.
func NewClusterCli(option *Option) *RedisCli {
	rawCli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        option.AddrList,
		Password:     option.Password,
		DialTimeout:  option.DialTimeout,
		ReadTimeout:  option.ReadTimeout,
		WriteTimeout: option.WriteTimeout,
		PoolSize:     option.PoolSize,
		MinIdleConns: option.MinIdleConns,
		MaxConnAge:   option.MaxConnAge,
	})

	for _, hook := range option.Hooks {
		rawCli.AddHook(hook)
	}

	cli := &RedisCli{
		option:   option,
		cluster:  true,
		redisCli: rawCli,
	}
	return cli
}

func NewSentinelCli(option *Option) *RedisCli {
	rawCli := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       option.MasterName,
		SentinelAddrs:    option.AddrList,
		Password:         option.Password,
		SentinelPassword: option.SentinelPassword,
		DialTimeout:      option.DialTimeout,
		ReadTimeout:      option.ReadTimeout,
		WriteTimeout:     option.WriteTimeout,
		PoolSize:         option.PoolSize,
		MinIdleConns:     option.MinIdleConns,
		MaxConnAge:       option.MaxConnAge,
	})

	for _, hook := range option.Hooks {
		rawCli.AddHook(hook)
	}

	return &RedisCli{
		option:   option,
		cluster:  false,
		redisCli: rawCli,
	}
}

func (rc *RedisCli) isCluster() bool {
	return rc.cluster
}

func (rc *RedisCli) RedisCmd() redis.Cmdable {
	return rc.redisCli
}

//Close 关闭连接.
func (rc *RedisCli) Close() {
	if rc.isCluster() {
		rs, ok := rc.redisCli.(*redis.ClusterClient)
		if ok {
			rs.Close()
		}
	} else {
		rs, ok := rc.redisCli.(*redis.Client)
		if ok {
			rs.Close()
		}
	}
}
