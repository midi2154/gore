package transformations

import (
	"strings"
)

func Cap(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
		}
		words = append(words[:1], words[i+1:]...)
		i--
	}
	return strings.Join(words, " ")
}
