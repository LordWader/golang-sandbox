package main

import "fmt"

type DHeapNode struct {
	elem, priority, arrayIndex, heapIndex int
}

type DHeap struct {
	Width   int
	Heap    []DHeapNode
	Address map[int]DHeapNode
}

func NewDHeap(d int) *DHeap {
	return &DHeap{
		Width:   d,
		Heap:    make([]DHeapNode, 0),
		Address: make(map[int]DHeapNode),
	}
}

func (dh *DHeap) bubbleUp(startIndex int) {
	lastIndex := startIndex
	for lastIndex != 0 {
		toCmp := (lastIndex - (lastIndex % dh.Width)) / dh.Width
		if dh.Heap[lastIndex].priority >= dh.Heap[toCmp].priority {
			dh.Heap[lastIndex], dh.Heap[toCmp] = dh.Heap[toCmp], dh.Heap[lastIndex]
			dh.Address[lastIndex] = dh.Heap[toCmp]
			dh.Address[toCmp] = dh.Heap[lastIndex]
			lastIndex = toCmp
		} else {
			break
		}
	}
}

func (dh *DHeap) Insert(elem, priority, index int) {
	newNode := DHeapNode{
		elem:       elem,
		priority:   priority,
		arrayIndex: index,
	}
	dh.Heap = append(dh.Heap, newNode)
	dh.Address[len(dh.Heap)-1] = newNode
	dh.bubbleUp(len(dh.Heap) - 1)
}

func (dh *DHeap) Top() DHeapNode {
	carry := dh.Heap[0]
	dh.Heap = dh.Heap[1:]
	if len(dh.Heap) > 0 {
		dh.bubbleUp(len(dh.Heap) - 1)
	}
	return carry
}

func (dh *DHeap) Modify(priority, index int) {
	el, ok := dh.Address[index]
	if !ok {
		panic("No node in heap!")
	}
	newNode := DHeapNode{
		elem:       el.elem,
		priority:   priority,
		arrayIndex: el.arrayIndex,
	}
	dh.Heap[el.heapIndex] = newNode
	dh.bubbleUp(el.heapIndex)
}

func main() {
	h := NewDHeap(2)
	arr := []int{5, 6, 1, 2, 7}
	for ind, el := range arr {
		h.Insert(ind, el, ind)
	}
	for len(h.Heap) > 0 {
		fmt.Println(h.Top())
	}
	fmt.Println("========================")
	for ind, el := range arr {
		h.Insert(ind, el, ind)
	}
	for i := 0; i < 3; i++ {
		h.Modify(i*3, i)
	}
	for len(h.Heap) > 0 {
		fmt.Println(h.Top())
	}
}
