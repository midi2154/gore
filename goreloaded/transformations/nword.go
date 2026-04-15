package transformations

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func Nword(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(up," || word[i] == "(low," || word[i] == "(cap," && i+1 < len(word) {
			word[i+1] = strings.TrimSuffix(word[i+1], ")")
			n, err := strconv.Atoi(word[i+1])
			if err != nil {
				fmt.Println("Error converting")
				return ""
			}
			for j := i - n; j < i; j++ {
				if word[i] == "(up," {
					word[j] = strings.ToUpper(word[j])
				}
				if word[i] == "(low," {
					word[j] = strings.ToLower(word[j])
				}
				if word[i] == "(cap," {
					word[j] = strings.ToUpper(word[j][:1]) + strings.ToLower(word[j][1:])
				}
			}
			word = append(word[:i], word[i+2:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}
