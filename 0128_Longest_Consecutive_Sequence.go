package leetcode

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	longest := 0
	for num := range m {
		if !m[num-1] {
			currentNum := num
			currentStreak := 1
			for m[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}
