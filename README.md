This package implements `sync.Lock` as a spinlock.

```
import (
	"github.com/xaionaro-go/atomicmap/spinlock"
)

var (
	locker = &spinlock.Locker{}
)

...

	locker.Lock()
	[doSomething]
	locker.Unlock()
```

```
BenchmarkMutexParallel-4        20000000                75.7 ns/op             0 B/op          0 allocs/op
BenchmarkSpinlockParallel-4     50000000                22.3 ns/op             0 B/op          0 allocs/op
BenchmarkMutexSingle-4          100000000               17.1 ns/op             0 B/op          0 allocs/op
BenchmarkSpinlockSingle-4       100000000               16.9 ns/op             0 B/op          0 allocs/op

go version go1.11 linux/amd64
```
