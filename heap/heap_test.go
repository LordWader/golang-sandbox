package main

import (
	"math/rand"
	"testing"
)

func GenRandomArray(size int) []int {
	mi := -100_000
	mx := 100_000
	arr := make([]int, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(mx-mi)+mi)
	}
	return arr
}

func TestDHeap_Top(t *testing.T) {
	// Initialize data
	a := GenRandomArray(1000000)
	h := NewDHeap(4)
	for _, elem := range a {
		h.Insert(elem)
	}
	carry := h.Top()
	// testing min d-heap
	for len(h.Heap) > 0 {
		c := h.Top()
		if c.priority < carry.priority {
			t.Fail()
		}
		carry = c
	}
}

func BenchmarkDHeap_Insert(b *testing.B) {
	a := GenRandomArray(100)
	// reset timer after generating sample
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := NewDHeap(4)
		for _, elem := range a {
			h.Insert(elem)
		}
	}
}

func BenchmarkDHeap_Top(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Init Heap
		a := GenRandomArray(100)
		h := NewDHeap(4)
		for _, elem := range a {
			h.Insert(elem)
		}
		for len(h.Heap) > 0 {
			h.Top()
		}
	}
}
