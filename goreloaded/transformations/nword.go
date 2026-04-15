package transformations

import (
	"fmt"
	"strconv"
	"strings"
)

func Nword(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		if (word[i] == "(up," || word[i] == "(low," || word[i] == "(cap,") && i+1 < len(word) {

			word[i+1] = strings.TrimSuffix(word[i+1], ")")
			n, err := strconv.Atoi(word[i+1])
			if err != nil {
				fmt.Println("Error converting")
				return ""
			}

			start := i - n
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				if word[i] == "(up," {
					word[j] = strings.ToUpper(word[j])
				}
				if word[i] == "(low," {
					word[j] = strings.ToLower(word[j])
				}
				if word[i] == "(cap," {
					if len(word[j]) > 0 {
						word[j] = strings.ToUpper(string(word[j][0])) + strings.ToLower(word[j][1:])
					}
				}
			}
			word = append(word[:i], word[i+2:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}


