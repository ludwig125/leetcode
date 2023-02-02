package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"1": {
			input: "()",
			want:  true,
		},
		"2": {
			input: "()[]{}",
			want:  true,
		},
		"3": {
			input: "(]",
			want:  false,
		},
		"4": {
			input: "{[]}",
			want:  true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isValid(tc.input); got != tc.want {
				t.Errorf("got: %t, want: %t", got, tc.want)
			}
		})
	}
}

// https://www.rapidtables.com/code/text/ascii-table.html
// 	ascii code
//  '(' // 40
// 	')' // 41
// 	'{' // 123
// 	'}' // 125
// 	'[' // 91
// 	']' // 93

func isValid(str string) bool {
	var s StackLinkedList
	s.Push(str[0]) // 一つ目は先に登録する
	for i := 1; i < len(str); i++ {
		// s.Print()

		// Stackの先頭を取り出して、
		// str[i]がそれと対になるものであれば、その先頭は削除する
		// 対にならないものであれば、Stackに追加する
		if IsPair(s.Peak(), str[i]) {
			s.Pop()
			continue
		}

		s.Push(str[i])
	}

	// 最後にStackに残りがあれば、
	// 対にならない括弧が残っているということなので有効ではない
	return s.IsEmpty()
}

func IsPair(val1, val2 byte) bool {
	switch val1 {
	case '(':
		if val2 == ')' {
			return true
		}
	case '{':
		if val2 == '}' {
			return true
		}
	case '[':
		if val2 == ']' {
			return true
		}
	}
	return false
}

// ref: https://www.educative.io/answers/how-to-implement-a-stack-using-a-linked-list-in-go
type Node struct {
	Value byte
	Next  *Node
}

type StackLinkedList struct {
	Head *Node
	Size int
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.Size == 0
}

func (s *StackLinkedList) Push(b byte) {
	if s.IsEmpty() {
		headNode := &Node{Value: b}
		s.Head = headNode
		s.Size = 1
		return
	}

	s.Head = &Node{Value: b, Next: s.Head}
	s.Size++
}

func (s *StackLinkedList) Pop() {
	if s.IsEmpty() {
		return
	}
	s.Head = s.Head.Next
	s.Size--
}

func (s *StackLinkedList) Peak() byte {
	if s.IsEmpty() {
		return 0 // IsPairで照合するのが０だと問題になるが今回のケースではOK
	}
	return s.Head.Value
}

func (s *StackLinkedList) Print() {
	nodes := s.GetNodes()
	fmt.Println(nodes)
}

func (s *StackLinkedList) GetNodes() []byte {
	if s.IsEmpty() {
		return nil
	}

	var nodes []byte
	tail := s.Head
	for {
		// fmt.Println("tail.Value", tail.Value)
		nodes = append(nodes, tail.Value)
		if tail.Next == nil {
			break
		}
		tail = tail.Next
	}

	return nodes
}
