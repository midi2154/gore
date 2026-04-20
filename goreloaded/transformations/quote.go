package transformations

import (
	"regexp"
)


func Quote(s string) string {
	q1 := regexp.MustCompile(`'\s*(.*?)\s*'`)
	s = q1.ReplaceAllString(s, " '$1' ")

	q2 := regexp.MustCompile(`"\s*(.*?)s*"`)
	s = q2.ReplaceAllString(s, " \"$1\" ")

	//q3 := regexp.MustCompile((`(\w+)(\s+)`))
	//s = q3.ReplaceAllString(s, `$1 `)

	text := strings.Fields(s)
	return strings.Join(text, " ")

}



	


	return s
}
