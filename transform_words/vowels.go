package transformation

func Vowels(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		if (word == "a") && (words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'i' || words[i+1][0] == 'o' || words[i+1][0] == 'u' || words[i+1][0] == 'h') {
			words[i] = "an"
		} else if (word == "A") && (words[i+1][0] == 'A' || words[i+1][0] == 'E' || words[i+1][0] == 'I' || words[i+1][0] == 'O' || words[i+1][0] == 'U' || words[i+1][0] == 'H') {
			words[i] = "An"
		}
	}
	return words
}
