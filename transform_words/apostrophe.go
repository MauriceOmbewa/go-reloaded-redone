package transformation

func Apostrophe(words []string) []string {
	x := 0
	for i := 0; i < len(words); i++ {
		if x == 0 && words[i] == "'" {
			words[i+1] = "'" + words[i+1]
			words[i] = ""
			words = append(words[:i], words[i+1:]...)
			x = 1
		} else if x == 1 && words[i] == "'" {
			words[i-1] += "'"
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
