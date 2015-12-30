package mutest


func Matches(pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}
	for i := 0; i < len(text); i++ {
		if matchHere(pattern, text[i:]) {
			return true
		}
	}
	return false
}

func matchHere(pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}
	return matchSingle(pattern[0], text[0]) && matchHere(pattern[1:], text[1:])
}

func matchSingle(patternChar, textChar byte) bool {
	return patternChar == '.' || patternChar == textChar
}

