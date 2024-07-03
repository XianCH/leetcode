package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result, different := 0, math.MaxInt16

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < different {
					different = abs(nums[i] + nums[j] + nums[k] - target)
					result = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func threeSumClosest2(nums []int, target int) int {
	n, result, different := len(nums), 0, math.MaxInt16
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for nums[i] == nums[i+1] && i < n-2 {
			i++
		}
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < different {
				different = abs(sum - target)
				result = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
