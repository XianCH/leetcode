package leetcode

func palindromeNumber(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return false
	}

	array := make([]int, 0, 32)
	for x > 0 {
		array = append(array, x%10)
		x = x / 10
	}

	for i, j := 0, len(array)-1; i <= j; i, j = i+1, j-1 {
		if array[i] != array[j] {
			return false
		}
	}
	return true
}
