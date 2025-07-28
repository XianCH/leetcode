package leetcode

// log(n) time complexity

func search_034(nums []int, target int) []int {
	result := []int{-1, -1}

	// 查找左边界
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 检查左边界是否有效
	if left >= len(nums) || nums[left] != target {
		return result
	}
	result[0] = left

	// 查找右边界
	right = len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	result[1] = right

	return result
}

// error
// runtime error: index out of range
func search_034_error(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if nums[mid-1] == target {
				return []int{mid - 1, mid}
			} else if nums[mid+1] == target {
				return []int{mid, mid + 1}
			}
			return []int{mid, mid}
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
