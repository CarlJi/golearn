package channel

import "testing"

// 测试命令: go test -bench . -benchmem
// 测试结果:
// testing: warning: no tests to run
// BenchmarkChanBool-4     	20000000	        78.1 ns/op	       0 B/op       0 allocs/op
// BenchmarkChanStruct-4   	20000000	        75.5 ns/op	       0 B/op       0 allocs/op
// PASS
// ok  	carlji.com/experiments/channel	3.239s
//
// 测试结论: 单纯作为信号场景使用，空的struct要比bool快一些
func BenchmarkChanBool(b *testing.B) {
	done := make(chan bool, 1)
	for n := 0; n < b.N; n++ {
		done <- true
		<-done
	}
}

func BenchmarkChanStruct(b *testing.B) {
	done := make(chan struct{}, 1)
	for n := 0; n < b.N; n++ {
		done <- struct{}{}
		<-done
	}
}
