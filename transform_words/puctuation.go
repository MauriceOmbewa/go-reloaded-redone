package transformation

func Punctuation(words []string) []string {
	for i := 0; i < len(words); i++ {
		if i > 0 && words[i][0] == '.' || words[i][0] == ',' || words[i][0] == '?' || words[i][0] == ':' || words[i][0] == ';' || words[i][0] == '!' {
			special := true // variable special is true when the whole word consists of special characters only
			for j := 0; j < len(words[i]); j++ {
				if words[i][j] == '.' || words[i][j] == ',' || words[i][j] == '!' || words[i][j] == ':' || words[i][j] == ';' || words[i][j] == '?' {
					words[i-1] += string(words[i][j])
				} else {
					special = false // variable special is false when a part of the word is made up of special characters and another part made up of letters
					words[i] = words[i][j:]
					break
				}
			}
			if special {
				words = append(words[:i], words[i+1:]...)
				i-- // indicates a decrement in i
			}
		}
	}
	return words
}
