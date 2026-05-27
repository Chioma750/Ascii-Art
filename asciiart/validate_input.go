package asciiart

import (
	"fmt"
	"strings"
)

// Function to validate user inputs to ensure only allowed inputs are supplied
// Takes user inputs as strings and returns strings
func ValidateInput(input string) string {
	input = strings.ReplaceAll(input, "\n", "\n")
	if input == "" {
		return ""
	}
	if input == "\\n" {
		fmt.Print("\n")
		return ""
	}
	for _, word := range input {
		if word < 32 || word > 126 {
			fmt.Print("invalid character")
			return ""
		}
	}
	return input
}
