package exchangecounter

import (
	"fmt"
	"regexp"
	"strings"
)

// ExtractAndConvertCurrencies extracts and converts currency strings from the input text
func ExtractAndConvertCurrencies(input string) []string {
	// Define the regex pattern for matching the currency formats
	pattern := `\$(\d{1,3}(,\d{3})*(\.\d{2})?)|(\d{1,3}(,\d{3})*\.\d{2})\s*美元`
	re := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(input, -1)

	// Process matches to remove currency symbols and commas
	var result []string
	for _, match := range matches {
		if match[1] != "" {
			result = append(result, strings.ReplaceAll(match[1], ",", ""))
		} else if match[4] != "" {
			result = append(result, strings.ReplaceAll(match[4], ",", ""))
		}
	}

	return result
}

func Test() {
	// Example input text
	input := "$200.49、$1,999.00、$99、50.00美元"

	// Extract and print matched and converted currency strings
	matches := ExtractAndConvertCurrencies(input)
	for _, match := range matches {
		fmt.Println(match)
	}
}
