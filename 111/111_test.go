package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		input func() TreeNode
		want  int
	}{
		"1": {
			input: func() TreeNode {
				tr := TreeNode{Val: 3}
				tr.Left = &TreeNode{Val: 9}
				tr.Right = &TreeNode{Val: 20}
				tr.Right.Left = &TreeNode{Val: 15}
				tr.Right.Right = &TreeNode{Val: 7}
				return tr
			},
			want: 2,
		},
		"2": {
			input: func() TreeNode {
				tr := TreeNode{Val: 2}
				tr.Right = &TreeNode{Val: 3}
				tr.Right.Right = &TreeNode{Val: 4}
				tr.Right.Right.Right = &TreeNode{Val: 5}
				tr.Right.Right.Right.Right = &TreeNode{Val: 6}
				return tr
			},
			want: 5,
		},
		"3": {
			input: func() TreeNode {
				tr := TreeNode{Val: 1}
				tr.Left = &TreeNode{Val: 2}
				tr.Right = &TreeNode{Val: 3}
				tr.Left.Left = &TreeNode{Val: 4}
				tr.Left.Right = &TreeNode{Val: 5}
				return tr
			},
			want: 2,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			treeNode := tc.input()
			treeNode.String()

			got := minDepth(&treeNode)
			// fmt.Println(name, tc.input, tc.want, got)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// String prints a visual representation of the tree
func (t *TreeNode) String() {
	fmt.Println("------------------------------------------------")
	stringify(t, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *TreeNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Val)
		stringify(n.Right, level)
	}
}

func minDepth(root *TreeNode) int {
	return bfs(root)
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	type q struct {
		treeNode *TreeNode
		depth    int
	}
	var que []q
	que = append(que, q{treeNode: root, depth: 1})

	for i := 0; i < len(que); i++ {
		front := que[i].treeNode
		frontD := que[i].depth

		// fmt.Println("front", i, front.Val, frontD)

		if front.Left == nil && front.Right == nil {
			// fmt.Println("end front", front.Val, frontD)
			depth = frontD
			break
		}
		// depth++
		if front.Left != nil {
			// fmt.Println("front.Left", front.Left.Val, frontD)
			que = append(que, q{treeNode: front.Left, depth: frontD + 1})
		}
		if front.Right != nil {
			// fmt.Println("front.Right", front.Right.Val, frontD)
			que = append(que, q{treeNode: front.Right, depth: frontD + 1})
		}
	}
	return depth
}
