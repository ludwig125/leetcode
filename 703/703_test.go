package main

// TODO: heapで書き直す

import (
	"reflect"
	"sort"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		k    int
		nums []int
		adds []int
		want []int
	}{
		"1": {
			k:    3,
			nums: []int{4, 5, 8, 2},
			adds: []int{3, 5, 10, 9, 4},
			want: []int{4, 5, 5, 8, 8},
		},
		"2": {
			k:    1,
			nums: []int{},
			adds: []int{-3, -2, -4, 0, 4},
			want: []int{-3, -2, -2, 0, 4},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			kTh := Constructor(tc.k, tc.nums)

			for i, v := range tc.adds {
				got := kTh.Add(v)
				if got != tc.want[i] {
					t.Errorf("got: %v, want: %v", got, tc.want[i])
				}
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := map[string]struct {
		nums []int
		add  int
		want []int
	}{
		"1": {
			nums: []int{8, 5, 4, 2},
			add:  3,
			want: []int{8, 5, 4, 3, 2},
		},
		"2": {
			nums: []int{8, 5, 4, 3, 2},
			add:  10,
			want: []int{10, 8, 5, 4, 3, 2},
		},
		"3": {
			nums: []int{8, 5, 4, 3, 2},
			add:  1,
			want: []int{8, 5, 4, 3, 2, 1},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			kTh := Constructor(3, tc.nums)

			kTh.Insert(tc.add)
			got := kTh.Nums
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

type KthLargest struct {
	K    int
	Nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return KthLargest{
		K:    k,
		Nums: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	this.Insert(val)

	// 問題に以下のようにあるので不要
	// It is guaranteed that there will be at least k elements in the array when you search for the kth element.
	// if len(this.Nums) < this.K {
	// 	return 0
	// }
	return this.Nums[this.K-1]
}

func (this *KthLargest) Insert(val int) {
	if len(this.Nums) == 0 {
		this.Nums = append(this.Nums, val)
		return
	}

	// // valがNumsの要素番号0の数値より大きい場合、要素番号0に挿入
	if val > this.Nums[0] {
		this.Nums = append([]int{val}, this.Nums[0:]...)
		return
	}

	// valがNumsの要素番号 len(Nums)-1の数値（一番小さい数値）より小さい場合、要素番号len(Nums)に挿入
	if val <= this.Nums[len(this.Nums)-1] {
		this.Nums = append(this.Nums, val)
		return
	}

	// valがNumsの要素番号iの数値より小さく、i+1の数値より大きい場合、要素番号i+1に挿入
	for i := 0; i < len(this.Nums); i++ {
		if this.Nums[i] >= val && this.Nums[i+1] < val {
			// https://golang.shop/post/go-slice-tricks-ja/
			this.Nums = append(this.Nums[:i+1], append([]int{val}, this.Nums[i+1:]...)...)
			return
		}
	}

}
