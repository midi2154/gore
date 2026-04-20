package transformations

import (
	"regexp"
)


func Punct(s string) string {
	p1 := regexp.MustCompile(`\s+([.,:;?!])`)
	s = p1.ReplaceAllString(s, "$1")

	p2 := regexp.MustCompile(`([.,:;?!])(\s*)(\w)`)
	s = p2.ReplaceAllString(s, "$1 $3")

	return s
}

