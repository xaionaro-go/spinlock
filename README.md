[![GoDoc](https://godoc.org/github.com/xaionaro-go/spinlock?status.svg)](https://pkg.go.dev/github.com/xaionaro-go/spinlock?tab=doc)
[![go report](https://goreportcard.com/badge/github.com/xaionaro-go/spinlock)](https://goreportcard.com/report/github.com/xaionaro-go/spinlock)
<p xmlns:dct="http://purl.org/dc/terms/" xmlns:vcard="http://www.w3.org/2001/vcard-rdf/3.0#">
  <a rel="license"
     href="http://creativecommons.org/publicdomain/zero/1.0/">
    <img src="http://i.creativecommons.org/p/zero/1.0/88x31.png" style="border-style: none;" alt="CC0" />
  </a>
</p>

# spinlock

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
