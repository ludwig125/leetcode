package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		head []int
		want []int
	}{
		"1": {
			head: []int{1, 1, 2},
			want: []int{1, 2},
		},
		"2": {
			head: []int{1, 1, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		// "3": {
		// 	head: []int{1},
		// 	pos:  -1,
		// 	want: false,
		// },
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			l := makeList(tc.head)
			l.Print()
			// 		got := hasCycle(tc.head, tc.pos)
			// 		if tc.want != got {
			// 			t.Fatalf("got: %v, want: %v", got, tc.want)
			// 		}

			got := deleteDuplicates(l).getAllNodes()
			// fmt.Println("got")
			// got.Print()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func makeList(heads []int) *ListNode {
	l := &ListNode{Val: heads[0]}
	for i := 1; i < len(heads); i++ {
		l.Add(heads[i])
	}

	fmt.Println("maked list", l)

	return l
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(i int) {
	tail := l
	for {
		// fmt.Println("tail:", tail)
		if tail.Next == nil {
			// fmt.Println("break")
			break
		}
		tail = tail.Next
	}
	// fmt.Println("tail1", tail.Val)
	tail.Next = &ListNode{Val: i}
	// fmt.Println("tail2", tail.Val)
}

func (l *ListNode) Print() {
	nodes := l.getAllNodes()
	fmt.Println("linked list nodes", nodes)
	// l.recursive(list)
}

func (l *ListNode) getAllNodes() []int {
	var visited []*ListNode
	var visitedInt []int

	tail := l
	for {
		if isVisited(tail, visited) {
			fmt.Println("visited", tail.Val)
			break
		}
		visited = append(visited, tail)
		visitedInt = append(visitedInt, tail.Val)
		// fmt.Println("tail``:", tail)
		if tail.Next == nil {
			// fmt.Println("break")
			break
		}
		tail = tail.Next
	}
	return visitedInt
}

func isVisited(node *ListNode, visited []*ListNode) bool {
	for _, n := range visited {
		if n == node {
			return true
		}
	}
	return false
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	tail := head
	for {
		if tail.Next == nil {
			break
		}
		if tail.Next.Val == tail.Val {
			if tail.Next.Next == nil {
				tail.Next = nil
				break
			}
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}
	return head
}
