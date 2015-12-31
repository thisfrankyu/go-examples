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
	matcher, rest := nextMatcher(pattern)
	if len(rest) > 0 && rest[0] == '*' {
		return matchStar(matcher, rest[1:], text)
	}
	if len(text) == 0 {
		return false
	}
	return matcher.match(text[0]) && matchHere(rest, text[1:])
}

func matchStar(matcher charMatcher, pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}
	return (matcher.match(text[0]) && matchStar(matcher, pattern, text[1:])) || matchHere(pattern, text)
}

type singleMatcher struct {
	c byte
}

func (matcher *singleMatcher) match(c byte) bool {
	return matcher.c == '.' || matcher.c == c
}

type bracketMatcher struct {
	chars string
}

func (matcher *bracketMatcher) match(c byte) bool {
	for i := 0; i < len(matcher.chars); i++ {
		if c == matcher.chars[i] {
			return true
		}
	}
	return false
}

type charMatcher interface {
	match(c byte) bool
}

func nextMatcher(p string) (charMatcher, string) {
	if p[0] == '[' {
		i := 0
		for i < len(p) && p[i] != ']' {
			i++
		}
		rest := p[i+1:]
		return &bracketMatcher{p[1:i]}, rest
	}
	return &singleMatcher{p[0]}, p[1:]
}
