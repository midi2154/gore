package transformations

import (
	"strings"
)

func Articles(text string) string {
	word := strings.Fields(text)

	for i := 0; i < len(word); i++ {
		if i+1 < len(word) && len(word[i+1]) > 0 {
			isVowel := strings.ContainsRune("aeiouhAEIOUH", rune(word[i+1][0]))
			if word[i] == "a" && isVowel {
				word[i] = "an"
			} else if word[i] == "A" && isVowel {
				word[i] = "An"
			} else if word[i] == "an" && !isVowel {
				word[i] = "a"
			} else if word[i] == "An" && !isVowel {
				word[i] = "A"
			}
		}
	}
	return strings.Join(word, " ")
}
