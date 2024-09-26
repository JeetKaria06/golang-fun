package main

import (
	"testing"
)

func BenchmarkFooPBV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooPBV(obj)
	}
}

func BenchmarkFooPBP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooPBP(&obj)
	}
}
