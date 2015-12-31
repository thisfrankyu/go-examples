package matcher


func Matches(pattern, text string) bool {
	for i := 0; i == 0 || i < len(text); i++ {
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
	if len(pattern) > 1 && pattern[1] == '*' {
		return matchStar(pattern[0], pattern[2:], text)
	}
	if len(text) == 0 {
		return false
	}
	return matchSingle(pattern[0], text[0]) && matchHere(pattern[1:], text[1:])
}

func matchStar(starPattern byte, pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}
	return (matchSingle(starPattern, text[0]) && matchStar(starPattern, pattern, text[1:])) || matchHere(pattern, text)
}

func matchSingle(patternChar, textChar byte) bool {
	return patternChar == '.' || patternChar == textChar
}

