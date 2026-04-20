package transformations

import (
	"regexp"
	"strings"
)

func Quote(s string) string {
	re := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = re.ReplaceAllString(s, "'$1'")

	re2 := regexp.MustCompile(`"\s+(.*?)\s+"`)
	s = re2.ReplaceAllString(s, " \"$1\" ")

	//s = re2.ReplaceAllString(s, `"1$"`)
	//re3 := regexp.MustCompile(`(\w+)(\s+)`)
	//s = re3.ReplaceAllString(s, `$1 `)

	text := strings.Fields(s)
	return strings.Join(text, " ")
}


