package exchangecounter

import (
	"reflect"
	"testing"
)

func TestExtractAndConvertCurrencies(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "$200.49、$1,999.00、$99、50.00美元",
			expected: []string{"200.49", "1999.00", "99", "50.00"},
		},
		{
			input:    "Nothing to match here.",
			expected: []string{},
		},
		{
			input:    "$100 and 50.00美元 are both valid.",
			expected: []string{"100", "50.00"},
		},
		{
			input:    "Incorrect formats: 100$, 999.00$, 50元.",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		got := ExtractAndConvertCurrencies(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("ExtractAndConvertCurrencies(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
