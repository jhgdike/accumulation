package heap_util

import (
	"container/heap"
	"fmt"
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_intHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_int64Heap() {
	h := &Int64Heap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, int64(3))
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

func Example_stringHeap() {
	h := &StringHeap{"abc", "bca", "cba"}
	heap.Init(h)
	heap.Push(h,"cab")
	fmt.Printf("minimum: %s\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%s ", heap.Pop(h))
	}
	// Output:
	// minimum: abc
	// abc bca cab cba
}
