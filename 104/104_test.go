package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		input []*int
		want  int
	}{
		"1": {
			input: []*int{
				intPtr(3),
				intPtr(9),
				intPtr(20),
				nil,
				nil,
				intPtr(15),
				intPtr(7),
			},
			want: 3,
		},
		"2": {
			input: []*int{
				intPtr(1),
				nil,
				intPtr(2),
			},
			want: 2,
		},
		"3": {
			input: []*int{
				intPtr(1),
				intPtr(2),
				intPtr(3),
				intPtr(4),
				intPtr(5),
			},
			want: 3,
		},
		"4": {
			input: []*int{
				intPtr(1),
				intPtr(2),
				intPtr(3),
				nil,
				nil,
				intPtr(4),
				intPtr(5),
			},
			want: 3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// got := maxDepth(tc.grid)

			treeNode := makeTree(tc.input)
			treeNode.String()

			got := maxDepth(&treeNode)
			// fmt.Println(name, tc.input, tc.want, got)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestLeftRight(t *testing.T) {
	tests := map[string]struct {
		length int
		want   []string
	}{
		"2": {
			length: 2,
			want:   []string{"l"},
		},
		"3": {
			length: 3,
			want:   []string{"l", "r"},
		},
		"4": {
			length: 4, // 1あまる
			want:   []string{"l", "r", "ll"},
		},
		"6": {
			length: 6,
			want:   []string{"l", "r", "ll", "lr", "rl"},
		},
		"7": {
			length: 7,
			want:   []string{"l", "r", "ll", "lr", "rl", "rr"},
		},
		"15": {
			length: 15,
			want:   []string{"l", "r", "ll", "lr", "rl", "rr", "lll", "llr", "lrl", "lrr", "rll", "rlr", "rrl", "rrr"},
		},
		"16": {
			length: 16, // 一つ次の段に入る
			want:   []string{"l", "r", "ll", "lr", "rl", "rr", "lll", "llr", "lrl", "lrr", "rll", "rlr", "rrl", "rrr", "llll"},
		},
		"31": {
			length: 31,
			want:   []string{"l", "r", "ll", "lr", "rl", "rr", "lll", "llr", "lrl", "lrr", "rll", "rlr", "rrl", "rrr", "llll", "lllr", "llrl", "llrr", "lrll", "lrlr", "lrrl", "lrrr", "rlll", "rllr", "rlrl", "rlrr", "rrll", "rrlr", "rrrl", "rrrr"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// got := maxDepth(tc.grid)

			got := leftRight(tc.length)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestInsertNode(t *testing.T) {
	tests := map[string]struct {
		root      *TreeNode
		leftRight []string
		val       []*int
		// want   []string
	}{
		"1": {
			root:      &TreeNode{Val: 1},
			leftRight: []string{"l"},
			val:       []*int{intPtr(2)},
			// want:   []string{"l"},
		},
		"2": {
			root:      &TreeNode{Val: 1},
			leftRight: []string{"l", "r"},
			val:       []*int{intPtr(2), intPtr(3)},
			// want:   []string{"l"},
		},
		"3": {
			root:      &TreeNode{Val: 1},
			leftRight: []string{"l", "r", "ll"},
			val:       []*int{intPtr(2), intPtr(3), intPtr(4)},
			// want:   []string{"l"},
		},
		"4": {
			root:      &TreeNode{Val: 1},
			leftRight: []string{"l", "r", "ll", "lr", "rl", "rr"},
			val:       []*int{intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
			// want:   []string{"l"},
		},
		"5": {
			root:      &TreeNode{Val: 1},
			leftRight: []string{"l", "r", "ll", "lr", "rl", "rr", "lll"},
			val:       []*int{intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7), intPtr(8)},
			// want:   []string{"l"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// got := maxDepth(tc.grid)
			treeNode := tc.root
			for i := 0; i < len(tc.leftRight); i++ {
				treeNode.Insert(tc.val[i], tc.leftRight[i])
			}

			// fmt.Println("treeNode", treeNode, treeNode.Left)
			treeNode.String()
			// if !reflect.DeepEqual(got, tc.want) {
			// 	t.Errorf("got: %v, want: %v", got, tc.want)
			// }
		})
	}
}

// length1はrootなのでlength2ならl, 3ならl, rを返す
func leftRight(length int) []string {
	if length == 2 {
		return []string{"l"}
	}
	if length == 3 {
		return []string{"l", "r"}
	}

	var s []string           // 最終的に返す対象
	s2 := []string{"l", "r"} // 次のmakeLeftRightでl, rのappend先になるs

	return makeLeftRight(length-1, s, s2)
}

func makeLeftRight(length int, s, s2 []string) []string {
	if length <= len(s) {
		// fmt.Println("len(s)", len(s))
		return s[:length]
	}

	// fmt.Println("s", s)

	var tmp []string
	for _, v := range s2 {
		tmp = append(tmp, v+"l")
		tmp = append(tmp, v+"r")
	}

	// fmt.Println("s s2", s, s2)

	// おおもとのsにs2をマージ
	// 次のmakeLeftRightでは、tmpがappendの対象となる
	return makeLeftRight(length, append(s, s2...), tmp)
}

func intPtr(i int) *int {
	return &i
}

// input の要素番号を与えると、depthを返す
// // ルートのnodeのdepthは１とカウントする
// func depth(i int) int {
// 	var depth int
// 	for j := i + 1; j > 0; j /= 2 {
// 		depth++
// 	}
// 	return depth
// }

// // そのdepthまでのnode数を返す
// func nodesNumUpToDepth(depth int) int {
// 	// d 1 => 1 = 2^0
// 	// d 2 => 1+2 = 2^0+2^1
// 	// d 3 => 1+2+4 = 2^0+2^1+2^2
// 	// d N => 2^0+2^1+2^2+...2^(d-1)

// 	var nodesNum int
// 	for d := 1; d <= depth; d++ {
// 		nodesNum += 2 ^ (d - 1)
// 	}
// 	return nodesNum
// }

func makeTree(input []*int) TreeNode {

	// rootのnodeを最初に登録
	treeNode := TreeNode{Val: *input[0]}

	if len(input) == 1 {
		return treeNode
	}

	lr := leftRight(len(input)) // lr は0番目の要素が"l"からなので注意
	for i := 0; i < len(input)-1; i++ {
		// fmt.Println("i input", i, input[i+1], lr[i])
		treeNode.Insert(input[i+1], lr[i])
		treeNode.String()
	}

	// depth1 root
	// depth2 l r
	// depth3 ll lr rl rr
	// depth4 lll llr lrl lrr rll rlr rrl rrr

	return treeNode
}

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 1)
}

func dfs(root *TreeNode, depth int) int {

	maxDepth := depth

	if root.Left != nil {
		tmp := dfs(root.Left, depth+1)
		if maxDepth < tmp {
			maxDepth = tmp
		}
	}

	if root.Right != nil {
		tmp := dfs(root.Right, depth+1)
		if maxDepth < tmp {
			maxDepth = tmp
		}
	}
	return maxDepth
}

func (t *TreeNode) Insert(val *int, leftRight string) {

	lr := string(leftRight[0])
	if len(leftRight) == 1 {
		if lr == "l" {
			if val != nil {
				t.Left = &TreeNode{Val: *val}
			}
		} else {
			if val != nil {
				t.Right = &TreeNode{Val: *val}
			}
		}
		return
	}

	if lr == "l" {
		t.Left.Insert(val, leftRight[1:])
	} else {
		t.Right.Insert(val, leftRight[1:])
	}
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
