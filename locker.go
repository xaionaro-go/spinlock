package spinlock

import (
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	iterationThreshold = 4 // the more this value the long we wait before begin running of "time.Sleep" in Lock()
)

const (
	// We leave positive values free for future needs (it's required to implement RLock in future; RLock will increment the value).
	unlocked = int32(0)
	locked   = -1
)

type Locker struct {
	state int32
}

// LockDo is just a wrapper for `Lock` + `Unlock`. It lockes the locker, runs the function `fn` and unlockes the locker.
//
// Warning! The `fn()` should not panic, otherwise the locker won't be unlocked
func (l *Locker) LockDo(fn func()) {
	l.Lock()
	fn()
	l.Unlock()
}

func (l *Locker) IsLocked() bool {
	return atomic.LoadInt32(&l.state) == locked
}

func (l *Locker) TryLock() bool {
	return atomic.CompareAndSwapInt32(&l.state, unlocked, locked)
}

func (l *Locker) Lock() {
	i := 0
	for !atomic.CompareAndSwapInt32(&l.state, unlocked, locked) {
		if i > iterationThreshold {
			time.Sleep(time.Nanosecond * time.Duration((50 + rand.Intn(950))))
		} else {
			runtime.Gosched()
		}
		i++
	}
}

func (l *Locker) Unlock() {
	if !atomic.CompareAndSwapInt32(&l.state, locked, unlocked) {
		panic(`Unlock()-ing non-locked locker`)
	}
}

// SetUnlocked resets the state of the locker. Use case: if you copy an object, you may want to reset this value.
func (l *Locker) SetUnlocked() {
	atomic.StoreInt32(&l.state, unlocked)
}
