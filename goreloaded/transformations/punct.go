package transformations

import (
	"regexp"
)

func Punct(word string)string{
	Checkpunct := regexp.MustCompile(`\s+([.,:;?!])`)
	word = Checkpunct.ReplaceAllString(word, "$1")

	checkpunct2 := regexp.MustCompile(`([.,:;?!])(\s*)(\w)`)
	word = checkpunct2.ReplaceAllString(word, "$1 $3")

	return word
}
