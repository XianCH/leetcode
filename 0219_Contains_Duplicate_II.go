package leetcode

func containsNearByDuplicate(nums []int, k int) bool {
	char := make(map[int]int)
	for i, num := range nums {
		if j, ok := char[num]; ok && i-j <= k {
			return true
		}
		char[num] = i
	}
	return false
}
