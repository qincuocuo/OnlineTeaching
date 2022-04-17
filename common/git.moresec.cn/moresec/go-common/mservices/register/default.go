package register

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Default struct{}

func (d *Default) Registered(ctx context.Context, name, host string) {
	for {
		select {
		case <-time.After(time.Second * 10):
			fmt.Println(errors.New("without register,name:" + name + ",host:" + host))
		case <-ctx.Done():
			fmt.Println("register done")
			break
		}
	}
}

func (d *Default) Query(name string) (string, error) {
	return "", errors.New("without register,name:" + name)
}
