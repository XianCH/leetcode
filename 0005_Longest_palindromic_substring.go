package leetcode

// too long n^3
func longestpalindrome(s string) string {
	max := 0
	var result string
	array := make([]string, 0)
	for _, char := range s {
		array = append(array, string(char))
	}
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if isPalindrome(array[i:j+1]) && max < len(array[i:j+1]) {
				result = s[i : j+1]
				max = len(array[i : j+1])
			}
		}
	}
	return result
}

func isPalindrome(s []string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// solution 2
func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		for right+1 < len(s) && s[right+1] == s[left] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			right, left = pr, pl
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}

//Manacher's Algorithm
// to do
