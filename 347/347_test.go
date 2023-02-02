package main

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		want []int
	}{
		"1": {
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		// "2": {
		// 	nums: []int{1},
		// 	k:    1,
		// 	want: []int{1},
		// },
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := topKFrequent(tc.nums, tc.k)
			// fmt.Println("got", got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func topKFrequent(nums []int, k int) []int {
	m := makeMap(nums)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	i := 0
	for k, v := range m {
		item := &Item{
			Value: k,
			Count: v,
			Index: i,
		}
		i++
		// fmt.Printf("Hist %#v\n", item)
		heap.Push(&pq, item)
		// pq.update(item, item.Value, 5)
	}

	var res []int
	for i := 0; i < k; i++ {
		v := heap.Pop(&pq)
		res = append(res, v.(*Item).Value)
		// fmt.Printf("koko %#v\n", v)
	}
	// for pq.Len() > 0 {
	// 	v := heap.Pop(&pq)
	// 	fmt.Println(v)
	// }
	return res
}

// 頻出回数ごとにmapに記録する関数
func makeMap(nums []int) map[int]int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	return m
}

// ref: https://pkg.go.dev/container/heap#example-package-HistHeap

type Item struct {
	Value int
	Count int
	Index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	fmt.Println("Less", pq[i].Count, pq[j].Count)
	// We want Pop to give us the highest, not lowest, Count so we use greater than here.
	return pq[i].Count > pq[j].Count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// // update modifies the Count and Value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, Value int, Count int) {
// 	item.Value = Value
// 	item.Count = Count
// 	heap.Fix(pq, item.Index)
// }

// heapの詳細はgithub上の別コードに記載
