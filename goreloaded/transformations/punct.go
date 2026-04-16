func Punct(s string) string {
	var result strings.Builder

	for i := 0; i < len(s); i++ {
		if strings.ContainsRune(".,!?;:", rune(s[i])) && i > 0 && s[i-1] == ' ' {
			result.WriteByte(s[i])
			continue
		}
		result.WriteByte(s[i])
	}

	return result.String()
}

var spaceBeforePunct = regexp.MustCompile(`\s+([.,!?;:])`)

func FixPunct(s string) string {
	return spaceBeforePunct.ReplaceAllString(s, "$1")
}

