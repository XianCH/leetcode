package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, "")
	if len(pattern) != len(words) {
		return false
	}

	p2s := make(map[byte]string)
	s2p := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p, w := pattern[i], words[i]
		if p2s[p] != "" && p2s[p] != w || s2p[w] != 0 && s2p[w] != p {
			return false
		}
		p2s[p] = w
		s2p[w] = p
	}

	return true
}
