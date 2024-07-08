package leetcode

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	total, lastint, num := 0, 0, 0
	if s == "" {
		return 0
	}
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = romanMap[char]
		if num < lastint {
			total -= num
		} else {
			total += num
		}
		lastint = num
	}
	return total
}
