package main

import (
	"fmt"
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
	h := NewDHeap(4, 1000000)
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
	arraySize := 100000
	a := GenRandomArray(arraySize)
	// reset timer after generating sample
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := NewDHeap(4, arraySize)
		for _, elem := range a {
			h.Insert(elem)
		}
	}
}

func BenchmarkDHeap_Top(b *testing.B) {
	arraySize := 100000
	for i := 0; i < b.N; i++ {
		// Init Heap
		a := GenRandomArray(arraySize)
		h := NewDHeap(4, arraySize)
		for _, elem := range a {
			h.Insert(elem)
		}
		for len(h.Heap) > 0 {
			h.Top()
		}
	}
}

func BenchmarkDHeap_TopForOptimalD(b *testing.B) {
	var table [20]*DHeap
	for i := 0; i < 20; i++ {
		arr := GenRandomArray(1000000)
		h := NewDHeap(i+2, 1000000)
		for _, elem := range arr {
			h.Insert(elem)
		}
		table[i] = h
	}
	for j, heap := range table {
		b.Run(fmt.Sprintf("heap.Top%d", heap.Width), func(b *testing.B) {
			for len(table[j].Heap) > 0 {
				table[j].Top()
			}
		})
	}
}
