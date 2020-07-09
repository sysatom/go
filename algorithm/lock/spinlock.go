package lock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type SpinLock struct {
	_    sync.Mutex
	lock uintptr
}

func NewSpinLock() *SpinLock {
	return &SpinLock{lock: 0}
}

func (l *SpinLock) Lock() {
	for !atomic.CompareAndSwapUintptr(&l.lock, 0, 1) {
		runtime.Gosched()
	}
}

func (l *SpinLock) Unlock() {
	atomic.StoreUintptr(&l.lock, 0)
}
