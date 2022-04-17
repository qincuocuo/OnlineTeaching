package mredis

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

// 参数.
type Option struct {
	URL             string //0.0.0.0:19000
	DB              int
	Password        string
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	DialConnTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	UsePool         bool // 默认提供连接池.

}

//NewDefaultOption 默认参数
func NewDefaultOption() *Option {
	return &Option{
		URL:             "127.0.0.1:6379",
		DB:              0,
		MaxIdle:         10,
		MaxActive:       100,
		IdleTimeout:     180 * time.Second,
		DialConnTimeout: 200 * time.Millisecond,
		ReadTimeout:     100 * time.Millisecond,
		WriteTimeout:    100 * time.Millisecond,
		UsePool:         true,
	}
}

type RedisCli struct {
	redisPool *redis.Pool // 连接池.

	option *Option // 连接参数.
	mu     sync.Mutex
}

//NewRedisCli 新建 redis 客户端.
func NewRedisCli(option *Option) *RedisCli {
	cli := &RedisCli{
		option: option,
	}
	// 使用连接池.
	cli.redisPool = &redis.Pool{
		MaxIdle:     option.MaxIdle,
		MaxActive:   option.MaxActive,
		IdleTimeout: option.IdleTimeout,
		Wait:        true,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				option.URL,
				redis.DialConnectTimeout(option.DialConnTimeout*time.Millisecond),
				redis.DialReadTimeout(option.ReadTimeout*time.Millisecond),
				redis.DialWriteTimeout(option.WriteTimeout*time.Millisecond),
				redis.DialDatabase(option.DB),
				redis.DialPassword(option.Password))
			if err != nil {
				fmt.Println("error:", err.Error())
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return cli
}

func NewRedisCliMock(option *Option, conn *redigomock.Conn) *RedisCli {
	cli := &RedisCli{
		option: option,
	}

	cli.redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return conn, nil
		},
		MaxIdle: option.MaxIdle,
	}

	return cli
}

//GetConn 返回连接及是否使用连接池.
func (rc *RedisCli) GetConn() redis.Conn {
	return rc.redisPool.Get()
}

//Close 关闭连接.
func (rc *RedisCli) Close() {
	_ = rc.redisPool.Close()
}
