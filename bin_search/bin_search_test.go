package main

import "testing"

var input = GenRandomSortedSlice(1_000_000)

func BenchmarkBatchSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BatchSearch(input, 1)
	}
}

func BenchmarkNormalBinSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalBinSearch(input, 1)
	}
}
