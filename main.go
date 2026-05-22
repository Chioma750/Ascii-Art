package main

import (
	"ascii-art/asciiart"
	"fmt"
	"os"
)

func main() {
	inputfile := "standard.txt"
	inputtext := os.Args[1]

	data, _ := asciiart.LoadBanner(inputfile)
	result := asciiart.BuildArt(inputtext, data)

	fmt.Println(result)
}
