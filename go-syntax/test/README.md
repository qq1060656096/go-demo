

```sh
go test -bench=基准测试名 -benchmem -count=次数 -run=none -v go-syntax/test/assert_test.go  
# 运行基准测试
go test -bench=BenchmarkStringJoinStringBuilder -benchmem -count=1 -cpu 1,2,4,8,12 -run=nonexxxxx -v go-syntax/test/assert_test.go  


goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
                                -GOMAXPROCS 运行次数        每次平均消耗时间         每次测试分配了多少内存   每次测试内存分配次数
BenchmarkStringJoinStringBuilder
BenchmarkStringJoinStringBuilder             859           1202804 ns/op         5863408 B/op         34 allocs/op
BenchmarkStringJoinStringBuilder-2           904           1155478 ns/op         5863408 B/op         34 allocs/op
BenchmarkStringJoinStringBuilder-4           962           1089219 ns/op         5863409 B/op         34 allocs/op
BenchmarkStringJoinStringBuilder-8           963           1096528 ns/op         5863420 B/op         34 allocs/op
BenchmarkStringJoinStringBuilder-12          919           1108123 ns/op         5863419 B/op         34 allocs/op
PASS
ok      command-line-arguments  6.276s
```

