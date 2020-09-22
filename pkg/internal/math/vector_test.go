package math

import "testing"

func BenchmarkCross(b *testing.B) {
	l := &Vector{0, 0}
	r := &Vector{1, 0}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cross(l, r)
	}
}

func BenchmarkSize(b *testing.B) {
	v := Vector{0, 0}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Size()
	}
}
