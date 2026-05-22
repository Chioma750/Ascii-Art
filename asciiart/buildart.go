package asciiart

import "fmt"

func BuildArt(str string, text []string) string {
	validated := ValidateInput(str)
	splistr := Splitting(validated)
	txt := ""
	for _, char := range splistr {
		if char == "" {
			fmt.Print("\n")
			return ""
		}
		for i := 0; i < 8; i++ {
			for _, cha := range char {
				stat := (int(cha)-32)*9 + 1
				text1 := text[stat+i]
				//txt = text1[i]
				print(text1)
			}
			println()
		}
	}
	return txt
}
