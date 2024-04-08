package main

import "fmt"

/*
		   0
          123
    345   678  91011
*/

type DHeapNode struct {
	elem, priority int
}

type DHeap struct {
	Width   int
	Heap    []DHeapNode
	Address map[*DHeapNode]int
}

func NewDHeap(d int) *DHeap {
	return &DHeap{
		Width: d,
		Heap:  make([]DHeapNode, 0),
	}
}

func (dh *DHeap) bubbleUp() {
	var toCmp int
	lastIndex := len(dh.Heap) - 1
	for lastIndex != 0 {
		for i := 0; i < dh.Width; i++ {
			if (lastIndex-i)%dh.Width == 0 {
				toCmp = (lastIndex - i) / dh.Width
			}
		}
		if dh.Heap[lastIndex].priority > dh.Heap[toCmp].priority {
			dh.Heap[lastIndex], dh.Heap[toCmp] = dh.Heap[toCmp], dh.Heap[lastIndex]
			lastIndex = toCmp
		} else {
			break
		}
	}
}

func (dh *DHeap) Insert(elem, priority int) {
	dh.Heap = append(dh.Heap, DHeapNode{
		elem:     elem,
		priority: priority,
	})
	dh.bubbleUp()
}

func (dh *DHeap) Top() DHeapNode {
	carry := dh.Heap[0]
	dh.Heap = dh.Heap[1:]
	if len(dh.Heap) > 0 {
		dh.bubbleUp()
	}
	return carry
}

func main() {
	h := NewDHeap(5)
	arr := []int{5, 6, 1, 2, 7}
	for ind, el := range arr {
		h.Insert(ind, el)
	}
	for len(h.Heap) > 0 {
		fmt.Println(h.Top())
	}
}
