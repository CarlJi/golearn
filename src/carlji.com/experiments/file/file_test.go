package main

import "testing"

// 测试命令: go test -bench .
// 测试结果:
//
// 测试结论: 单纯作为信号场景使用，空的struct要比bool快一些

func BenchmarkFileO_Sync(b *testing.B) {
}
