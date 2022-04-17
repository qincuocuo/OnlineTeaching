package register

import (
	"context"
	"fmt"
	"git.moresec.cn/moresec/go-common/mclusterrds"
	"time"
)

type Redis struct {
	*mclusterrds.RedisCli
	KeyPrefix    string
	ExpireSecond int32
}

func (r *Redis) key(name string) string {
	if r.KeyPrefix == "" {
		r.KeyPrefix = "debug:pprof:"
	}

	return r.KeyPrefix + name
}

func (r *Redis) TTL() time.Duration {
	if r.ExpireSecond > 0 {
		return time.Duration(r.ExpireSecond) * time.Second
	}

	r.ExpireSecond = 30
	return 30 * time.Second
}

func (r *Redis) Registered(ctx context.Context, name, host string) {
	b, err := r.SetNX(r.key(name), host, r.TTL())
	if err != nil || !b {
		panic(err)
	}

	for {
		select {
		case <-time.After(r.TTL() / 3):
			err = r.Expire(r.key(name), r.ExpireSecond)
			if err != nil {
				fmt.Println("expire error:", err)
			}
		case <-ctx.Done():
			fmt.Println("register done")
			break
		}
	}
}

func (r *Redis) Query(name string) (string, error) {
	return r.Get(r.key(name))
}
