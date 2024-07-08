package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	s = removeSymbols(s)

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func removeSymbols(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, s)
}
