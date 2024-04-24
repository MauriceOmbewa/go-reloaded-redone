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
			words[i-1] = strings.ToLower(words[i-1]) // transforming the whole word to lowercase
			words = append(words[:i], words[i+1:]...) // appending after transformation
		} else if i > 0 && word == "(low," { 
			temp := words[i+1][:len(words[i+1])-1]
			converted, err := strconv.Atoi(temp) // converting the number after "(low," to type int
			if err != nil {
				fmt.Println("Error: ", err)
			} else if i-converted >= 0 {
				for j := 1; j <= converted; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				fmt.Println("Number of words is less than the number after '(low,'") // handling an error where the number of words requested for transformation is larger than actual words present
			}
		}
	}
	return words
}

func Uppercase(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1]) // transforming the whole word to lowercase
			words = append(words[:i], words[i+1:]...) // appending after transformation
		} else if i > 0 && word == "(up," {
			temp := words[i+1][:len(words[i+1])-1]
			converted, err := strconv.Atoi(temp) // converting the number after "(up," to type int
			if err != nil {
				fmt.Println("Error: ", err)
			} else if i-converted >= 0 {
				for j := 0; j < converted; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				fmt.Println("Number of words is less than the number after '(up,'") // handling an error where the number of words requested for transformation is larger than actual words present
			}
		}
	}
	return words
}

func Capitalize(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		if i > 0 && word == "(cap)" {
			capitalizedfirst := strings.ToUpper(string(words[i-1][0])) + words[i-1][1:] // transforming only the first character in the word to uppercase using ToUpper
			words[i-1] = capitalizedfirst
			words = append(words[:i], words[i+1:]...) // appending after transformation
		} else if i > 0 && word == "(cap," {
			temp := words[i+1][:len(words[i+1])-1]
			converted, err := strconv.Atoi(temp) // converting the number after "(cap," to type int
			if err != nil {
				fmt.Println("Error: ", err)
			} else if i-converted >= 0 {
				for j := 1; j <= converted; j++ {
					capitalizedfirst := strings.ToUpper(string(words[i-j][0])) + words[i-j][1:]
					words[i-j] = capitalizedfirst
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				fmt.Println("Number of words is less than the number after '(cap,'") // handling an error where the number of words requested for transformation is larger than actual words present
			}
		}
	}
	return words
}
