package example

import "testing"

var BenchInts = map[string]int64{
	"first":   34,
	"second":  12,
	"third":   56,
	"fourth":  78,
	"fifth":   90,
	"sixth":   23,
	"seventh": 56,
	"eighth":  78,
	"ninth":   90,
	"tenth":   23,
}

var Benchfloats = map[string]float64{
	"first":   35.98,
	"second":  26.99,
	"third":   56.98,
	"fourth":  78.98,
	"fifth":   90.98,
	"sixth":   23.98,
	"seventh": 56.98,
	"eighth":  78.98,

	"ninth": 90.98,
	"tenth": 23.98,
}

func BenchmarkSumGenInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumIntsOrFloats(BenchInts)
	}
}
func BenchmarkSumGenFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumIntsOrFloats(Benchfloats)
	}
}
func BenchmarkSumIntLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumInts(BenchInts)
	}
}
func BenchmarkSumFloatLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumFloats(Benchfloats)
	}
}
