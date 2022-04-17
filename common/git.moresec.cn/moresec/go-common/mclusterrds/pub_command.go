package mclusterrds

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

// 发布模式.
func (rc *RedisCli) Publish(key, msg string) (int64, error) {
	if !rc.isCluster() {
		client := rc.redisCli.(*redis.Client)
		return client.Publish(context.TODO(), key, msg).Result()
	}
	client := rc.redisCli.(*redis.ClusterClient)
	return client.Publish(context.TODO(), key, msg).Result()
}

func (rc *RedisCli) SubscribeWithTimeout(channel string, timeout time.Duration) (*redis.Message, error) {
	subFunc := func(pubSub *redis.PubSub) (*redis.Message, error) {
		defer pubSub.Close()
		for {
			msg, err := pubSub.ReceiveTimeout(context.TODO(), timeout)
			if err != nil {
				return nil, err
			}
			switch data := msg.(type) {
			case *redis.Subscription:
			case *redis.Message:
				return data, nil
			default:
				return nil, errors.New("not receive")
			}
		}
	}

	if rc.isCluster() {
		conn := rc.redisCli.(*redis.ClusterClient)
		pubSub := conn.Subscribe(context.TODO(), channel)
		return subFunc(pubSub)
	} else {
		conn := rc.redisCli.(*redis.Client)
		pubSub := conn.Subscribe(context.TODO(), channel)
		return subFunc(pubSub)
	}
}

func (rc *RedisCli) Subscribe(callback func(message redis.Message), close <-chan struct{}, channel ...string) error {
	subFunc := func(pubSub *redis.PubSub) error {
		defer pubSub.Close()
		for {
			select {
			case <-close:
				pubSub.Unsubscribe(context.TODO(), channel...)
				return nil
			default:
				_, err := pubSub.Receive(context.TODO())
				if err != nil {
					return err
				}
				ch := pubSub.Channel()
				for msg := range ch {
					callback(*msg)
				}
			}
		}
	}

	if rc.isCluster() {
		conn := rc.redisCli.(*redis.ClusterClient)
		pubSub := conn.Subscribe(context.TODO(), channel...)
		return subFunc(pubSub)
	} else {
		conn := rc.redisCli.(*redis.Client)
		pubSub := conn.Subscribe(context.TODO(), channel...)
		return subFunc(pubSub)
	}
}
