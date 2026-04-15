package transformations

import (
	"strings"
)

func Articles(text string) string {
	word := strings.Fields(text)

	for i := 0; i < len(word); i++ {
		if i+1 < len(word) && strings.ContainsAny(word[i+1][:1], "aeiouhAEIOUH") {
			if word[i] == "a" {
				word[i] = "an"
			} else if word[i] == "A" {
				word[i] = "An"
			}
		}
	}
	return strings.Join(word, " ")
}
