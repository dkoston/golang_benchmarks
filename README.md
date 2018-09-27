# golang_lock_benchmarks

Benchmarks on different styles of locking in golang


## Run the benchmarks
```
go test -bench=.
```

### Deferring Unlock() vs explicit Unlock()

Many people like to defer Unlock() like:

```
defer mu.Unlock()
```

Rather than adding an explicit Unlock() when there are multiple
logic branches in a function like:

```
func updateCache() {
  mu.Lock()
  defer mu.Unlock()

  if branch1 {
    return
  }

  if branch2 {
    return
  }

  return
}

```

Does the `defer` have a real cost?

#### Benchmarks

```
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkDeferFunction-8            	 5000000	       274 ns/op
BenchmarkExplicitUnlockFunction-8   	10000000	       223 ns/op
PASS
ok  	_golang_unlock_benchmark	4.141s

$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkDeferFunction-8            	 5000000	       274 ns/op
BenchmarkExplicitUnlockFunction-8   	10000000	       224 ns/op
PASS
ok  	_golang_unlock_benchmark	4.137s

$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkDeferFunction-8            	 5000000	       268 ns/op
BenchmarkExplicitUnlockFunction-8   	10000000	       235 ns/op
PASS
ok  	_golang_unlock_benchmark	4.210s
```

### Catch-all RWMutex vs Mutex

Sometimes, you'll create a RWMutex rather than a Mutex as you don't
know all the use cases where your shared data structure will be used later.

Is is ok to use RWMutex without losing performance or should you start with Mutex
and refactor if you need read locks?

#### Benchmarks

```
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkDeferFunction-8            	 5000000	       279 ns/op
BenchmarkExplicitUnlockFunction-8   	10000000	       220 ns/op
BenchmarkKeyLock-8                  	10000000	       232 ns/op
BenchmarkKeyRWLock-8                	 5000000	       258 ns/op
PASS
ok  	_golang_unlock_benchmark	8.263s

$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkDeferFunction-8            	 5000000	       272 ns/op
BenchmarkExplicitUnlockFunction-8   	10000000	       225 ns/op
BenchmarkKeyLock-8                  	 5000000	       239 ns/op
BenchmarkKeyRWLock-8                	 5000000	       260 ns/op
PASS
ok  	_golang_unlock_benchmark	7.152s

$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkDeferFunction-8            	 5000000	       287 ns/op
BenchmarkExplicitUnlockFunction-8   	10000000	       220 ns/op
BenchmarkKeyLock-8                  	10000000	       232 ns/op
BenchmarkKeyRWLock-8                	 5000000	       254 ns/op
PASS
ok  	_golang_unlock_benchmark	8.255s
```
