package mredis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 发布模式.
func (rc *RedisCli) Publish(key, msg string) (int64, error) {
	conn := rc.GetConn()
	defer func() {
		conn.Close()
	}()
	return redis.Int64(conn.Do("PUBLISH", key, msg))
}

func (rc *RedisCli) Subscribe(callback func(redis.Message) error, close <-chan struct{}, channel ...string) error {
	psc := redis.PubSubConn{Conn: rc.redisPool.Get()}
	if err := psc.Subscribe(redis.Args{}.AddFlat(channel)...); err != nil {
		return err
	}

	done := make(chan error, 1)
	go func() {
		defer psc.Close()
		for {
			switch msg := psc.Receive().(type) {
			case redis.Message:
				if err := callback(msg); err != nil {
					done <- err
					return
				}
			case error:
				done <- fmt.Errorf("redis pubsub receive err: %v", msg)
				return
			case redis.Subscription:
				if msg.Count == 0 {
					done <- nil
					return
				}
			}
		}
	}()

	// health check.
	ticker := time.NewTimer(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-close:
			if err := psc.Unsubscribe(); err != nil {
				return fmt.Errorf("redis pubsub unsubscribe err: %v", err)
			}
			return nil
		case err := <-done:
			return err
		case <-ticker.C:
			if err := psc.Ping(""); err != nil {
				return err
			}
		}
	}
}
