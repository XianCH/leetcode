package leetcode

func minSubArrayLen(target int, nums []int) int {
	start, end, sum, minLen := 0, 0, 0, len(nums)+1
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			if end-start+1 < minLen {
				minLen = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
