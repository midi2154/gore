package transformations

import (
	"regexp"
)

func Quote(text string) string {
	s := regexp.MustCompile(`\s*([']+)\s*`)
	return s.ReplaceAllString(text, "$1")
}
