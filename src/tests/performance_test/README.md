# Performance tests

Example run:

```sh
$ go test -bench=.
goos: linux
goarch: amd64
pkg: main/tests/performance_test
cpu: 12th Gen Intel(R) Core(TM) i7-1260P
BenchmarkDatabase-16    	       1	12775914913 ns/op
PASS
ok  	main/tests/performance_test	12.779s
```

Each file in `./outputs/` is the `ns/op` results over multiple runs against the code at the time of the output's commit.
