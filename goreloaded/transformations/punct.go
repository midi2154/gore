package transformations

import (
	"strings"
)

var new = strings.NewReplacer(
	" !!!", "!!!",
	" ...", "...",
	" !!", "!!",
	" ,", ",",
	" .", ".",
	" ?", "?",
	" ;", ";",
	" :", ";",
)

func Punct(s string) string {
	result := new.Replace(s)
	return result
}
