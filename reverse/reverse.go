package reverse

import (
	"strings"
)

func ReverseSentence(sentence string) string {
	if len(sentence) == 0 {
		return ""
	}
	if strings.Index(sentence, " ") == -1 {
		return sentence
	}

	var reversed = ""
	var reversedSentence = ""

	for i := 0; i < len(sentence); i++ {
		reversed += string(sentence[len(sentence)-1-i])
	}

	var reversedWords []string

	for _, word := range strings.Split(reversed, " ") {
		var reversedWord string
		for i := 0; i < len(word); i++ {
			reversedWord += string(word[len(word)-1-i])
		}
		reversedWords = append(reversedWords, reversedWord)
	}
	reversedSentence = strings.Join(reversedWords, " ")

	return reversedSentence
}
