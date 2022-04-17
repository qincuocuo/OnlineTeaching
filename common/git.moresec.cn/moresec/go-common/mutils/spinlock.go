package mutils

import (
	"runtime"
	"sync/atomic"
)

// this is a good candiate for a lock-free structure.

type SpinLock struct{ lock uintptr }

func (l *SpinLock) Lock() {
	for !atomic.CompareAndSwapUintptr(&l.lock, 0, 1) {
		runtime.Gosched()
	}
}
func (l *SpinLock) Unlock() {
	atomic.StoreUintptr(&l.lock, 0)
}
