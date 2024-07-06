package leetcode

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 1. 首先，我们需要处理 k 值，因为如果 k 大于数组的长度，旋转数组的结果会和旋转 k % len(nums) 的结果相同。所以我们先将 k 更新为 k % len(nums)。
//
// 2. 然后，我们可以将整个数组反转。
//
// 3. 接着，我们将数组的前 k 个元素反转。
//
// 4. 最后，我们将数组的剩余元素反转。
//
// 这个方法的基本思想是利用反转操作来实现旋转。例如，如果我们要将数组 [1,2,3,4,5,6,7] 向右旋转 3 步，我们可以先将整个数组反转得到 [7,6,5,4,3,2,1]，然后将前 3 个元素反转得到 [5,6,7,4,3,2,1]，最后将剩余元素反转得到 [5,6,7,1,2,3,4]，这就是旋转后的结果。
