package mservices

import (
	"context"
	"git.moresec.cn/moresec/go-common/mlog"
	"go.uber.org/zap"
)

type Debugger struct {
	*Option
}

func (d *Debugger) Name() string {
	return d.name
}

func (d *Debugger) Serve() error {
	err := d.Check()
	if err != nil {
		mlog.Error("option check", zap.Error(err), zap.String("name", d.Name()))
		return err
	}

	mlog.Info("option checked", zap.Any("data", d.Option), zap.String("name", d.Name()))

	go d.register.Registered(d.ctx, d.Name(), d.Host)

	err = d.server.Serve(d.l)
	if err != nil {
		mlog.Error("http serve error", zap.Error(err), zap.String("name", d.Name()))
	}
	return err
}

func (d *Debugger) Stop() {
	d.cancel()
	err := d.server.Shutdown(context.TODO())
	if err != nil {
		mlog.Error("shutdown", zap.Error(err), zap.String("name", d.Name()))
	}
}
