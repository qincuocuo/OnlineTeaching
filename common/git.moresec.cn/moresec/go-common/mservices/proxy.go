package mservices

import (
	"context"
	"git.moresec.cn/moresec/go-common/mlog"
	"go.uber.org/zap"
)

type Proxy struct {
	*Option
}

func (p *Proxy) Name() string {
	return p.name
}

func (p *Proxy) Serve() error {
	err := p.Check()
	if err != nil {
		mlog.Error("option check", zap.Error(err), zap.String("name", p.Name()))
		return err
	}

	mlog.Info("option checked", zap.Any("data", p.Option), zap.String("name", p.Name()), zap.String("name", p.Name()))

	err = p.server.Serve(p.l)
	if err != nil {
		mlog.Error("http serve error", zap.Error(err), zap.String("name", p.Name()))
	}
	return err
}

func (p *Proxy) Stop() {
	p.cancel()
	err := p.server.Shutdown(context.TODO())
	if err != nil {
		mlog.Error("shutdown", zap.Error(err), zap.String("name", p.Name()))
	}
}
