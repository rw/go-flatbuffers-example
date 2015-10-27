```
$ GOPATH=$(pwd) go run main.go && GOPATH=$(pwd) go test -test.bench=.
Arthur Dent has id 42. The encoded data is 48 bytes long.
testing: warning: no tests to run
PASS
BenchmarkWrite-4        10000000               214 ns/op         223.35 MB/s           0 B/op          0 allocs/op
BenchmarkRead-4         20000000                72.4 ns/op       662.90 MB/s           0 B/op          0 allocs/op
BenchmarkRoundtrip-4     5000000               302 ns/op         158.71 MB/s           0 B/op          0 allocs/op
...
```
