package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		input  []int
		target int
		want   []int
	}{
		"1": {
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		"2": {
			input:  []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		"3": {
			input:  []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := twoSum(tc.input, tc.target)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			// fmt.Println("nums[i]+nums[j]", nums[i], nums[j])
			if nums[i]+nums[j] == target {
				return append([]int{i}, j)
			}
		}
	}
	return nil
}
