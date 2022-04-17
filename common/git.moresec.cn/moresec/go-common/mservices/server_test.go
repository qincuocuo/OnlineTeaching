package mservices

import (
	"git.moresec.cn/moresec/go-common/mclusterrds"
	"git.moresec.cn/moresec/go-common/mlog"
	"git.moresec.cn/moresec/go-common/mlog/conf"
	"git.moresec.cn/moresec/go-common/mlog/zaplog"
	"git.moresec.cn/moresec/go-common/mservices/register"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebugger_Serve(t *testing.T) {
	mlog.SetLogger(zaplog.New(
		conf.WithProjectName("server_debugger"),
		conf.WithIsStdOut("yes"),
	))

	debugger := NewService(WithType(ServiceDebugger), WithRegister(&register.Redis{
		RedisCli: mclusterrds.NewRedisCli(&mclusterrds.Option{
			Addr:     "192.168.120.140:6379",
			Password: "moresec#sec",
		}),
	}))

	assert.Equal(t, nil, debugger.Serve())

	prof := NewProfilingService("127.0.0.1:8080")
	prof.Serve()
}


func TestNewProfilingService(t *testing.T) {
	prof := NewProfilingService("127.0.0.1:8080")
	prof.Serve()
}

func TestProxy_Serve(t *testing.T) {
	mlog.SetLogger(zaplog.New(
		conf.WithProjectName("server_proxy"),
		conf.WithIsStdOut("yes"),
	))

	proxy := NewService(WithType(ServiceProxy), WithRegister(&register.Redis{
		RedisCli: mclusterrds.NewRedisCli(&mclusterrds.Option{
			Addr:     "192.168.120.140:6379",
			Password: "moresec#sec",
		}),
	}), WithHost("127.0.0.1:6969"))

	assert.Equal(t, nil, proxy.Serve())
}
