package main

import (
	"strings"
)

func ValidateInput(input string) string {
	input = strings.ReplaceAll(input, "\n", "\n")
	if input == "" {
		return ""
	}
	if input == "\n" {
		return "\n"
	}
	for _, word := range input {
		if word < 32 || word > 126 {
			return "invalid character"
		}
	}
	return input
}
