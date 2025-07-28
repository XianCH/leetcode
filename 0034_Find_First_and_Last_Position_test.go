package leetcode

import "testing"

func TestSearch_034(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}

	for _, test := range tests {
		result := search_034(test.nums, test.target)
		if len(result) != 2 || result[0] != test.expect[0] || result[1] != test.expect[1] {
			t.Errorf("search_034(%v, %d) = %v; expected %v", test.nums, test.target, result, test.expect)
		}
	}
}
