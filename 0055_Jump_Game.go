package leetcode

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	reachmax := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= reachmax {
			reachmax = i
		}
	}
	return reachmax == 0
}

//贪心算法
