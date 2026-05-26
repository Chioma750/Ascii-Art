package asciiart

import (
	"fmt"
	"strings"
)

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
