package transformation

import (
	"fmt"
	"strconv"
	"strings"
)

func Lowercase(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if i > 0 && word == "(low," {
			temp := words[i+1][:len(words[i+1])-1]
			converted, err := strconv.Atoi(temp)
			if err != nil {
				fmt.Println("Error: ", err)
			} else if i-converted >= 0 {
				for j := 0; j < converted; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return words
}

func Uppercase(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if i > 0 && word == "(up," {
			temp := words[i+1][:len(words[i+1])-1]
			converted, err := strconv.Atoi(temp)
			if err != nil {
				fmt.Println("Error: ", err)
			} else if i-converted >= 0 {
				for j := 0; j < converted; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return words
}

func Capitalize(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if i > 0 && word == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if i > 0 && word == "(cap," {
			temp := words[i+1][:len(words[i+1])-1]
			converted, err := strconv.Atoi(temp)
			if err != nil {
				fmt.Println("Error: ", err)
			} else if i-converted >= 0 {
				for j := 1; j <= converted; j++ {
					words[i-j] = strings.Title(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				fmt.Println("Number of words is less than the number after '(cap,'")
			}
		}
	}
	return words
}
