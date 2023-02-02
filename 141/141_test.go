package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		head []int
		pos  int
		want bool
	}{
		"1": {
			head: []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		"2": {
			head: []int{1, 2},
			pos:  0,
			want: true,
		},
		"3": {
			head: []int{1},
			pos:  -1,
			want: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			l := makeList(tc.head, tc.pos)
			l.Print()
			// 		got := hasCycle(tc.head, tc.pos)
			// 		if tc.want != got {
			// 			t.Fatalf("got: %v, want: %v", got, tc.want)
			// 		}

			if got := hasCycle(l); got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func makeList(heads []int, pos int) *ListNode {
	l := &ListNode{Val: heads[0]}
	for i := 1; i < len(heads); i++ {
		l.Add(heads[i])
	}

	fmt.Println("maked list", l)

	if pos >= 0 {
		l.AddPos(pos)
	}
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

func (l *ListNode) AddPos(pos int) {
	var posTarget *ListNode
	cnt := 0
	tail := l
	for {
		if cnt == pos {
			posTarget = tail
		}
		cnt++

		if tail.Next == nil {
			tail.Next = posTarget
			break
		}
		tail = tail.Next
	}
}

func (l *ListNode) Print() {
	var visited []*ListNode

	tail := l
	for {
		if isVisited(tail, visited) {
			fmt.Println("visited", tail.Val)
			break
		}
		visited = append(visited, tail)
		// fmt.Println("tail``:", tail)
		if tail.Next == nil {
			// fmt.Println("break")
			break
		}
		tail = tail.Next
	}

	fmt.Println("linked list nodes", visited)
	// l.recursive(list)
}

func isVisited(node *ListNode, visited []*ListNode) bool {
	for _, n := range visited {
		if n == node {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	tail := head

	var visitedNodes []*ListNode

	for {
		if isVisited(tail, visitedNodes) {
			return true
		}
		visitedNodes = append(visitedNodes, tail)

		fmt.Println("tail", tail.Val, tail.Next)
		if tail.Next == nil {
			break
		}
		tail = tail.Next
	}
	return false
}
