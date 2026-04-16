package transformations

import (
	"strings"
)

func Up(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i-- 
		}
	}

	return strings.Join(words, " ")
}

