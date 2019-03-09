package spinlock

import (
	"sync"
	"testing"
)

func benchmarkParallelLocker(b *testing.B, locker sync.Locker) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			locker.Lock()
			locker.Unlock()
		}
	})
}

func BenchmarkMutexParallel(b *testing.B) {
	benchmarkParallelLocker(b, &sync.Mutex{})
}

func BenchmarkSpinlockParallel(b *testing.B) {
	benchmarkParallelLocker(b, &Locker{})
}

func benchmarkSingleLocker(b *testing.B, locker sync.Locker) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		locker.Lock()
		locker.Unlock()
	}
}

func BenchmarkMutexSingle(b *testing.B) {
	benchmarkSingleLocker(b, &sync.Mutex{})
}

func BenchmarkSpinlockSingle(b *testing.B) {
	benchmarkSingleLocker(b, &Locker{})
}
