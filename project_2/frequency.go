package wordutils

import (
	"strings"
	"unicode"
)

func GetWordFrequency(text string) map[string]int {
	frequency := make(map[string]int)
	text = strings.ToLower(text)

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}