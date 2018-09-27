# golang_unlock_benchmark

Shows the non-trivial cost of deferring Unlock with sync.Mutex

## Benchmarking

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
