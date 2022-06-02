package nombre

func firstLetter(s string) rune {
	if len(s) == 0 {
		return '\n'
	} else {
		return rune(s[0])
	}
}
