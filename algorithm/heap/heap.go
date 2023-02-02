package main

import (
	"container/heap"
	"fmt"
)

/*
heapに渡すInterfaceには、sort.Interfaceに含まれた、
Less, Swap, Lenメソッドと、
Push, Popメソッドを用意する必要がある

実装例：https://pkg.go.dev/container/heap#example-package-IntHeap

https://pkg.go.dev/container/heap#Interface
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

https://pkg.go.dev/sort#Interface
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	ih := &intHeap{}

	heap.Init(ih)
	heap.Push(ih, 3)
	heap.Push(ih, 1)
	heap.Push(ih, 10)
	heap.Push(ih, 7)

	for ih.Len() > 0 {
		v := heap.Pop(ih) // heap.Pop経由で値を取り出す
		fmt.Println(v)
	}

}
