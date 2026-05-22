package ascii

func BuildArt(str string, text []string) string {
	validated := ValidateInput(str)
	splistr := Splitting(validated)
	txt := ""
	for _, char := range splistr {
		if char == "" {
			//fmt.Println("\n")
			return ""
		}
		for i := range 8 {
			for _, cha := range char {
				stat := (int(cha)-32)*9 + 1
				text1 := text[stat : stat+8]
				txt = text1[i]
				print(txt)
			}
			println()
		}
	}
	return txt
}
