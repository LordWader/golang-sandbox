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

func (dh *DHeap) getMaxPriority(idx int) (DHeapNode, int) {
	// identify leaf node
	if idx*dh.Width+1 > len(dh.Heap)-1 {
		return DHeapNode{}, -1
	}
	// dummy start
	carry, ind := dh.Heap[idx*dh.Width+1], idx*dh.Width+1
	for i := idx*dh.Width + 2; i < (idx+1)*dh.Width+1; i++ {
		if i > len(dh.Heap)-1 {
			break
		}
		if dh.Heap[i].priority >= carry.priority {
			carry = dh.Heap[i]
			ind = i
		}
	}
	return carry, ind
}

func (dh *DHeap) siftDown(start int) {
	for start != len(dh.Heap) {
		carry, ind := dh.getMaxPriority(start)
		if ind == -1 {
			return
		}
		if dh.Heap[start].priority <= carry.priority {
			dh.Heap[start], dh.Heap[ind] = dh.Heap[ind], dh.Heap[start]
			start = ind
		} else {
			return
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
	dh.Heap[0] = dh.Heap[len(dh.Heap)-1]
	dh.Heap = dh.Heap[:len(dh.Heap)-1]
	if len(dh.Heap) > 0 {
		dh.siftDown(0)
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
	h := NewDHeap(4)
	arr := []int{5, 6, 1, 2, 7}
	for ind, el := range arr {
		h.Insert(ind, el, ind)
	}
	for len(h.Heap) > 0 {
		fmt.Println(h.Top())
	}
	//fmt.Println("========================")
	//for ind, el := range arr {
	//	h.Insert(ind, el, ind)
	//}
	////for i := 0; i < 3; i++ {
	//h.Modify(0, 0)
	////}
	//for len(h.Heap) > 0 {
	//	fmt.Println(h.Top())
	//}
}
