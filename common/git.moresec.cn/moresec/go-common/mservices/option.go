package mservices

import (
	"context"
	"git.moresec.cn/moresec/go-common/mlog"
	"git.moresec.cn/moresec/go-common/mservices/healthcheck"
	"git.moresec.cn/moresec/go-common/mservices/register"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type OptionHandler func(option *Option)

type Option struct {
	name       string
	Host       string // register value
	ServerType ServiceType
	register   register.Register
	hc         healthcheck.HealthChecker
	server     *http.Server

	ctx    context.Context
	cancel context.CancelFunc
	l      net.Listener
}

func WithName(name string) OptionHandler {
	return func(option *Option) {
		option.name = name
	}
}

func WithType(t ServiceType) OptionHandler {
	return func(option *Option) {
		option.ServerType = t
	}
}

func WithHost(h string) OptionHandler {
	return func(option *Option) {
		option.Host = h
	}
}

func WithHealthChecker(hc healthcheck.HealthChecker) OptionHandler {
	return func(option *Option) {
		option.hc = hc
	}
}

func WithRegister(r register.Register) OptionHandler {
	return func(option *Option) {
		option.register = r
	}
}

func WithServer(s *http.Server) OptionHandler {
	return func(option *Option) {
		option.server = s
	}
}

func WithCtx(ctx context.Context) OptionHandler {
	return func(option *Option) {
		option.ctx, option.cancel = context.WithCancel(ctx)
	}
}

type Checker interface {
	Check() error
}

type Updater interface {
	Update(o *Option)
}

func (o *Option) Check() (err error) {
	o.l, err = net.Listen("tcp", o.Host)
	if err != nil {
		mlog.Error("net listen error", zap.Error(err))
		return
	}

	if o.name == "" {
		o.name = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	}

	if o.Host == "" {
		o.Host = "127.0.0.1"
	}
	if !strings.Contains(o.Host, ":") {
		o.Host = net.JoinHostPort(o.Host, strconv.Itoa(o.l.Addr().(*net.TCPAddr).Port))
	}

	if o.register == nil {
		o.register = &register.Default{}
	}

	if o.ctx == nil {
		o.ctx, o.cancel = context.WithCancel(context.TODO())
	}

	if o.server == nil {
		o.server, err = initRouter(o)
		if err != nil {
			return
		}
	}

	if o.hc == nil {
		o.hc = &healthcheck.Default{}
	}

	return
}
