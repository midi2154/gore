package transformations

import "regexp"

func Quote(text string) string {
	reg := regexp.MustCompile(`'s+(.*?)\s+'`)
	text = reg.ReplaceAllString(text, `'$1'`)
	return text
}
