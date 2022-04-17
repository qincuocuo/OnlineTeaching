package mbase

import (
	"time"
)

type TimerHandler func()
type MTimer struct {
	isrunning bool
	describe  string
	timer     *time.Ticker
	tc        chan int
}

func NewMTimer(desc string) *MTimer {
	timer := &MTimer{
		isrunning: false,
		describe:  desc,
		tc:        make(chan int),
	}
	return timer
}

func (thiz *MTimer) IsRunning() bool {
	return thiz.isrunning
}

func (thiz *MTimer) Start(d time.Duration, f TimerHandler) {

	if thiz.isrunning {
		return
	}

	if thiz.timer == nil {
		thiz.timer = time.NewTicker(d)
	}

	// 定时器回调在单独协程里，要考虑竞争问题
	thiz.isrunning = true
	go func() {
	T:
		for {
			select {
			case <-thiz.tc:
				break T
			case <-thiz.timer.C:
				f()
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
		thiz.timer = nil
		thiz.isrunning = false
	}()
}

func (thiz *MTimer) Stop() {
	if !thiz.isrunning {
		return
	}

	if thiz.timer != nil {
		thiz.timer.Stop()
	}

	if thiz.tc != nil {
		select {
		case thiz.tc <- 0:
		default:
		}
	}
}
