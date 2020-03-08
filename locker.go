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

// Locker is a spinlock-based implementation of sync.Locker.
type Locker struct {
	state int32
}

// LockDo is just a wrapper for `Lock` + `Unlock`.
// It locks the locker, runs the function `fn` and then unlocks the locker.
func (l *Locker) LockDo(fn func()) {
	l.Lock()
	defer l.Unlock()

	fn()
}

// IsLocked returns true if the locker is currently locked.
func (l *Locker) IsLocked() bool {
	return atomic.LoadInt32(&l.state) == locked
}

// TryLock performs a non-blocking attempt to lock the locker and returns true if successful.
func (l *Locker) TryLock() bool {
	return atomic.CompareAndSwapInt32(&l.state, unlocked, locked)
}

// Lock waits until the locker with be unlocked (if it is not) and then locks it.
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

// Unlock unlocks the locker.
func (l *Locker) Unlock() {
	if !atomic.CompareAndSwapInt32(&l.state, locked, unlocked) {
		panic(`Unlock()-ing non-locked locker`)
	}
}

// SetUnlocked resets the state of the locker.
//
// Use case: if you copy an object, you may want to reset this value.
func (l *Locker) SetUnlocked() {
	atomic.StoreInt32(&l.state, unlocked)
}
