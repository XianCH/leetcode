package leetcode

func LongCharacter(s string) int {

	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastL, ok := lastOccurred[ch]; ok && lastL >= start {
			start = lastL + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
