package mclusterrds

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

func (rc *RedisCli) PublishWithCtx(ctx context.Context, key, msg string) (int64, error) {
	if !rc.isCluster() {
		client := rc.redisCli.(*redis.Client)
		return client.Publish(ctx, key, msg).Result()
	}
	client := rc.redisCli.(*redis.ClusterClient)
	return client.Publish(ctx, key, msg).Result()
}

func (rc *RedisCli) SubscribeWithTimeoutAndCtx(ctx context.Context, channel string, timeout time.Duration) (*redis.Message, error) {
	subFunc := func(pubSub *redis.PubSub) (*redis.Message, error) {
		defer pubSub.Close()
		for {
			msg, err := pubSub.ReceiveTimeout(ctx, timeout)
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
		pubSub := conn.Subscribe(ctx, channel)
		return subFunc(pubSub)
	} else {
		conn := rc.redisCli.(*redis.Client)
		pubSub := conn.Subscribe(ctx, channel)
		return subFunc(pubSub)
	}
}

func (rc *RedisCli) SubscribeWithCtx(ctx context.Context, callback func(message redis.Message), close <-chan struct{}, channel ...string) error {
	subFunc := func(pubSub *redis.PubSub) error {
		defer pubSub.Close()
		for {
			select {
			case <-close:
				pubSub.Unsubscribe(ctx, channel...)
				return nil
			default:
				_, err := pubSub.Receive(ctx)
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
		pubSub := conn.Subscribe(ctx, channel...)
		return subFunc(pubSub)
	} else {
		conn := rc.redisCli.(*redis.Client)
		pubSub := conn.Subscribe(ctx, channel...)
		return subFunc(pubSub)
	}
}
