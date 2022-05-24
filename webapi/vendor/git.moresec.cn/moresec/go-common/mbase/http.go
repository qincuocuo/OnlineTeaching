package mbase

import (
	"context"
	"net"
	"net/http"
	"time"
)

//NewHTTPClient 生成带有超时机制http客户端.
// 用法： NewHttpClient(2) // 就是2秒超时. 不可这样 NewHttpClient(2*time.Second)
func NewHTTPClient(timeout int64) *http.Client {
	DefaultClient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(time.Duration(timeout) * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
				if err != nil {
					return nil, err
				}
				err = c.SetDeadline(deadline)
				return c, err
			},
			ResponseHeaderTimeout: time.Second * time.Duration(timeout),
		},
	}
	return DefaultClient
}
