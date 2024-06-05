package main

import (
	"fmt"
)

type DHeapNode struct {
	priority int
}

type DHeap struct {
	Width int
	Heap  []*DHeapNode
}

func intMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Min D-Heap implementation
Modify could be implemented, but it depends on some complex "business logic"
*/

func NewDHeap(d, length int) *DHeap {
	return &DHeap{
		Width: d,
		Heap:  make([]*DHeapNode, 0, length),
	}
}

func (dh *DHeap) swap(first, second int) {
	dh.Heap[first], dh.Heap[second] = dh.Heap[second], dh.Heap[first]
}

func (dh *DHeap) bubbleUp(startIndex int) {
	lastIndex := startIndex
	for lastIndex != 0 {
		toCmp := (lastIndex - (intMax(1, lastIndex%dh.Width))) / dh.Width
		if dh.Heap[lastIndex].priority < dh.Heap[toCmp].priority {
			dh.swap(lastIndex, toCmp)
			lastIndex = toCmp
		} else {
			break
		}
	}
}

func (dh *DHeap) getMaxPriority(idx int) (*DHeapNode, int) {
	// dummy start
	carry, ind := dh.Heap[idx*dh.Width+1], idx*dh.Width+1
	for i := idx*dh.Width + 2; i < intMin((idx+1)*dh.Width+1, len(dh.Heap)); i++ {
		if dh.Heap[i].priority < carry.priority {
			carry = dh.Heap[i]
			ind = i
		}
	}
	return carry, ind
}

func (dh *DHeap) getLeafIndex(idx int) bool {
	if idx*dh.Width+1 > len(dh.Heap)-1 {
		return true
	}
	return false
}

func (dh *DHeap) siftDown(start int) {
	for !dh.getLeafIndex(start) {
		carry, ind := dh.getMaxPriority(start)
		if dh.Heap[start].priority > carry.priority {
			dh.swap(start, ind)
			start = ind
		} else {
			return
		}
	}
}

func (dh *DHeap) Insert(priority int) {
	newNode := &DHeapNode{
		priority: priority,
	}
	dh.Heap = append(dh.Heap, newNode)
	dh.bubbleUp(len(dh.Heap) - 1)
}

func (dh *DHeap) Top() *DHeapNode {
	carry := dh.Heap[0]
	dh.Heap[0] = dh.Heap[len(dh.Heap)-1]
	dh.Heap = dh.Heap[:len(dh.Heap)-1]
	if len(dh.Heap) > 0 {
		dh.siftDown(0)
	}
	return carry
}

func main() {
	arr := []int{89205, 44720, 70877, -30824, -75881, -1732, 3873, 34959, 55048, -72593}
	h := NewDHeap(4, len(arr))
	for _, el := range arr {
		h.Insert(el)
	}
	for len(h.Heap) > 0 {
		fmt.Println(h.Top())
	}
}
