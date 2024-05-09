package leetcode

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		got := IntToRoman(tt.num)
		if got != tt.want {
			t.Errorf("IntToRoman(%d) = %s,want :%s", tt.num, got, tt.want)
		}
	}
}
