package leetcode

import "testing"

func TestLongCharacter(t *testing.T) {
	testCase := []struct {
		input  string
		expect int
	}{
		{"abcabccc", 3},
		{"ccccccccc", 1},
		{"pwwkeww", 3},
	}

	for _, tc := range testCase {
		result := LongCharacter(tc.input)
		if result != tc.expect {
			t.Errorf("Input:%s,Expected:%d,Got:%d", tc.input, tc.expect, result)
		}
	}
}
