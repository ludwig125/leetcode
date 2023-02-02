package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		"1": {
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2},
		},
		"2": {
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := intersection(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func intersection(nums1 []int, nums2 []int) []int {

	m1 := make(map[int]bool)
	for _, v := range nums1 {
		if _, ok := m1[v]; !ok {
			m1[v] = true
		}
	}
	// fmt.Println(m1)

	m2 := make(map[int]bool)
	for _, v := range nums2 {
		if _, ok := m2[v]; !ok {
			m2[v] = true
		}
	}
	// fmt.Println(m2)

	var res []int
	for k := range m1 {
		if m2[k] {
			res = append(res, k)
		}
	}
	return res
}
