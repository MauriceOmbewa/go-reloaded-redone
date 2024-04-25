package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	transformation "Go-reloaded/transform_words"
)

func main() {
	if len(os.Args) != 3 { // dealing with the error of having more or less than required arguments
		fmt.Println("Error! Expected: go run main.go sample.txt result.txt")
		return
	}

	inputfilename := os.Args[1]
	outputfilename := os.Args[2]

	inputfile, err := os.Open(inputfilename)
	if err != nil { // handling error during opening of input file
		fmt.Println("Error: ", err)
	}
	defer inputfile.Close() // here the closing of the file is rescheduled to after its use

	outputfile, err := os.Create(outputfilename)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer outputfile.Close()

	scanner := bufio.NewScanner(inputfile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // breaking the text scanned to words separated by space in between

		// calling the functions for manipulation
		words = transformation.Vowels(words)
		words = transformation.Capitalize(words)
		words = transformation.Uppercase(words)
		words = transformation.Lowercase(words)
		words = transformation.Binary(words)
		words = transformation.Hexadecimal(words)
		words = transformation.Apostrophe(words)
		words = transformation.Punctuation(words)

		finalstr := strings.Join(words, " ")              // rejoining the words and putting space in between
		_, err := outputfile.WriteString(finalstr + "\n") // writting the result in the output file
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
