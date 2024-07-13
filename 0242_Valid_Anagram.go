package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	char := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		char[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		if _, ok := char[t[j]]; ok {
			if char[t[j]] > 0 {
				char[t[j]]--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
