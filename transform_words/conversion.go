package transformation

import (
	"fmt"
	"strconv"
)

func Hexadecimal(words []string) []string {
	for i := 1; i < len(words); i++ {
		word := words[i]
		if word == "(hex)" {
			temp := words[i-1]
			convtemp, err := strconv.ParseInt(temp, 16, 64)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			words[i-1] = strconv.FormatInt(convtemp, 10)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func Binary(words []string) []string {
	for i := 1; i < len(words); i++ {
		word := words[i]
		if word == "(bin)" {
			temp := words[i-1]
			convtemp, err := strconv.ParseInt(temp, 2, 64)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			words[i-1] = strconv.FormatInt(convtemp, 10)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
