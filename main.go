package main

import (
	"ascii-art/asciiart"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("ensure you enter: go run . Hello  standared")
		return 
	}

	inputtext := os.Args[1]
	inputfile := os.Args[2]

	data, _ := asciiart.LoadBanner(inputfile + ".txt")
	result := asciiart.BuildArt(inputtext, data)

	fmt.Print(result)
}
