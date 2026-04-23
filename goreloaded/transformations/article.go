package transformations

import (
	"strings"
)

package transformations

import (
	"strings"
)

func Articles(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		if i+1 < len(word) {

			next := strings.Trim(word[i+1], `"'`)

			// find first letter only
			var ch byte
			for i := 0; i < len(next); i++ {
				if (next[i] >= 'a' && next[i] <= 'z') || (next[i] >= 'A' && next[i] <= 'Z') {
					ch = next[i]
					break
				}
			}

			isVowel := strings.ContainsAny(string(ch), "aeiouAEIOU")

			if word[i] == "a" && isVowel {
				word[i] = "an"
			}
			if word[i] == "A" && isVowel {
				word[i] = "An"
			}
			if word[i] == "an" && !isVowel && ch != 0 {
				word[i] = "a"
			}
			if word[i] == "An" && !isVowel && ch != 0 {
				word[i] = "A"
			}
		}
	}

	return strings.Join(word, " ")
}


or 

package transformations

import(
	"strings"
)
func Article(s string) string {
	f := strings.Fields(s)
	
	for i := 0; i < len(f); i++ {
		text := f[i]
		word := strings.ContainsAny(string(f[i][0]), "aeiouh'AEIOUH")
		if text == "an" && word {
			f[i] = "a"
		} else if text == "a" && word {
			f[i] = "an"
		} else if text == "An" && word {
			f[i] = "A"
		} else if text == "A" && word {
			f[i] = "An"
		}
	}
	return strings.Join(f, " ")
}


