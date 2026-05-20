package main

import "strings"

func Splitting (input string) []string {
	if input == "" {
		return []string{}
	}

	lines := strings.Split(input, "\\n")
	return lines
}