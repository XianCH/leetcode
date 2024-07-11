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

func lengthOfLongestSubstring(s string) int {
	start, maxLen := 0, 0
	charMap := make(map[rune]int)
	for end, char := range s {
		if _, ok := charMap[char]; ok {
			for s[start] != byte(char) {
				delete(charMap, rune(s[start]))
				start++
			}
			start++
		} else {
			charMap[char] = end
			if end-start+1 > maxLen {
				maxLen = end - start + 1
			}
		}
	}
	return maxLen
}
