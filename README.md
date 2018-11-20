go test hello.go hello_test.go

cd iteration
go test -bench=BenchmarkRepeat
go test repeat.go repeat_test.go -v