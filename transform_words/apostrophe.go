package transformation

func Apostrophe(words []string) []string {
	x := 0 // helps the program know the first apostrophe and how to handle it
	for i := 0; i < len(words); i++ {
		if x == 0 && words[i] == "'" {
			words[i+1] = "'" + words[i+1]
			words[i] = ""
			words = append(words[:i], words[i+1:]...)
			x = 1 // after handling the first apostrophe, we change value of x to 1 to tell the program that we're done with the first apostrophe and the next to be ecountered is handled differently
		} else if x == 1 && words[i] == "'" {
			words[i-1] += "'"
			words = append(words[:i], words[i+1:]...)
			x = 0 // after handling the second apostrophe, we change value of x to 0 to tell the program that the next apostrophe is to be handled like the first apostrophe
		}
	}
	return words
}
