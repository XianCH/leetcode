package leetcode

import (
	"reflect"
	"testing"
)

func TestTwonum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 6, 2}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		got := Twonum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Twonum(%v,%d)=%v,want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
