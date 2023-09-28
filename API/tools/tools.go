package tools

// FormatName lower case letter and replace spaces with _
func FormatName(str string) string {
	var tmpStp string

	for _, letter := range str {
		if letter >= 'A' && letter <= 'Z' {
			tmpStp += string(letter + 32)
		} else if letter == ' ' {
			tmpStp += "_"
		} else {
			tmpStp += string(letter)
		}
	}

	return tmpStp
}
