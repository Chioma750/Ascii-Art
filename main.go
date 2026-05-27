package main

import (
	"ascii-art/asciiart"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(`ensure you enter: go run . "Hello" "standard"`)
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Argument should not be more than 3")
		return
	}

	inputtext := os.Args[1]
	inputfile := os.Args[2]

	data, err := asciiart.LoadBanner(inputfile + ".txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	result := asciiart.BuildArt(inputtext, data)

	fmt.Print(result)
}
