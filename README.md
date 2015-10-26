```
$ GOPATH=$(pwd) go run main.go && GOPATH=$(pwd) go test -test.bench=.
Arthur Dent has id 42. The encoded data is 48 bytes long.
testing: warning: no tests to run
PASS
BenchmarkWrite-4        10000000               219 ns/op               0 B/op          0 allocs/op
BenchmarkRead-4         20000000                67.6 ns/op             0 B/op          0 allocs/op
BenchmarkRoundtrip-4     5000000               285 ns/op               0 B/op          0 allocs/op
...
```
