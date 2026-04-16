package transformations

import (
	"regexp"
	"strings"
)

func Punct(text string) string {
	s := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	return strings.TrimSpace(s.ReplaceAllString(text, "$1 "))
}
