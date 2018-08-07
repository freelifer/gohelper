package utils

import (
	"testing"
)

func Test_RandomString(t *testing.T) {
	result, err := RandomString(12)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 12 {
		t.Error("len != 12")
	}
}

func Test_NewSessionID(t *testing.T) {
	sessionId := NewSessionID()

	t.Log(sessionId)
}
func Test_Guid(t *testing.T) {
	sessionId := Guid()

	t.Log(sessionId)
}

// go test -bench=".*" -count=5
func Benchmark_RandomString32(b *testing.B) {
	// 必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		RandomString(4)
	}
}

func Benchmark_NewSessionID(b *testing.B) {
	// 必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		NewSessionID()
	}
}
func Benchmark_Guid(b *testing.B) {
	// 必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		Guid()
	}
}

// go test -bench=”.”
// 测试并发效率
// func BenchmarkLoopsParallel(b *testing.B) {
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			NewSessionID()
// 		}
// 	})
// }
