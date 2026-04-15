package main

import (
	"fmt"
	"goreloaded/transformations"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run . input.txt output.txt")
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	output := Process(string(data))

	os.WriteFile(outputFile, []byte(output), 0644)
	fmt.Println("Modification succesful")
}

func Process(text string) string {
	result := transformations.Hex(string(text))
	result = transformations.Bin(string(result))
	result = transformations.Up(string(result))
	result = transformations.Low(string(result))
	result = transformations.Cap(string(result))
	result = transformations.Articles(string(result))
	result = transformations.Quote(string(result))
	result = transformations.Nword(string(result))
	result = transformations.Punct(string(result))

	return result
}
