package leetcode

func canConstruct(ransomNot string, magazine string) bool {
	charMap := make(map[rune]int)
	for _, char := range magazine {
		charMap[char]++
	}

	for _, char := range ransomNot {
		if _, ok := charMap[char]; ok {
			if charMap[char] > 0 {
				charMap[char]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
