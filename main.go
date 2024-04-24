package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	transformation "final/transform_words"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error! Expected: go run main.go sample.txt result.txt")
	}

	inputfilename := os.Args[1]
	outputfilename := os.Args[2]

	inputfile, err := os.Open(inputfilename)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer inputfile.Close()

	outputfile, err := os.Create(outputfilename)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer outputfile.Close()

	scanner := bufio.NewScanner(inputfile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		// Input functions here
		words = transformation.Vowels(words)
		words = transformation.Capitalize(words)
		words = transformation.Uppercase(words)
		words = transformation.Lowercase(words)
		words = transformation.Binary(words)
		words = transformation.Hexadecimal(words)
		words = transformation.Apostrophe(words)
		words = transformation.Punctuation(words)

		finalstr := strings.Join(words, " ")
		_, err := outputfile.WriteString(finalstr + "\n")
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
