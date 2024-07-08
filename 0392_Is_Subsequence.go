package leetcode

func isSubsequence(s string, t string) bool {
	currentIndex := 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[currentIndex] {
			currentIndex++
		}
		if currentIndex == len(s)-1 {
			return true
		}
	}
	return false
}

func isSubsequence2(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)

}
