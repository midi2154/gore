import "regexp"

var quoteCleaner = regexp.MustCompile(`(['"])\s*(.*?)\s*\1`)

func FixQuotes(s string) string {
	return quoteCleaner.ReplaceAllString(s, `$1$2$1`)
}
