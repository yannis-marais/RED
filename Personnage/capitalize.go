package ProjetRED

func Capitalize(s string) string {
	result := []byte(s)
	newWord := true

	for i := 0; i < len(result); i++ {
		c := result[i]

		if c >= 'a' && c <= 'z' {
			if newWord {
				result[i] = c - 32
			}
			newWord = false
		} else if c >= 'A' && c <= 'Z' {
			if !newWord {
				result[i] = c + 32
			}
			newWord = false
		} else if c >= '0' && c <= '9' {
			newWord = false
		} else {
			newWord = true
		}
	}

	return string(result)
}
