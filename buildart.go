package asciiart

import "strings"

func BuildArt(str string, text []string) string {
	validated := ValidateInput(str)
	splistr := Splitting(validated)
	var words strings.Builder
	for _, char := range splistr {
		if char == "" {
			words.WriteString("\n")
			continue
		}
		for i := range 8 {
			for _, cha := range char {
				stat := (int(cha)-32)*9 + 1
				word := text[stat : stat+8]
				words.WriteString(word[i])

			}
			words.WriteString("\n")
		}
	}
	return words.String()
}
