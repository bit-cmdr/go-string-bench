# Go String Concatenation Benchmarking

This is to benchmark different string concatenation methods in Go

## Usage

Pull down locally and run

```sh
$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/bit-cmdr/go-string-bench
BenchmarkSprintf-12            2000000         751 ns/op       263 B/op        7 allocs/op
BenchmarkPlusConcat-12         3000000         519 ns/op       167 B/op        3 allocs/op
BenchmarkStringBuilder-12      2000000         606 ns/op       304 B/op        7 allocs/op
BenchmarkBytesBuffer-12        2000000         668 ns/op       455 B/op        5 allocs/op
PASS
ok    github.com/bit-cmdr/go-string-bench 8.186s
```
