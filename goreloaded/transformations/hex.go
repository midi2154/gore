package transformations

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Hex(text string) string {
	fields := strings.Fields(text)

	for i := 0; i < len(fields); i++ {
		if fields[i] == "(hex)" && i > 0 {
			decimal, err := strconv.ParseInt(fields[i-1], 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			fields[i-1] = fmt.Sprint(decimal)
			fields = append(fields[:i], fields[i-1:]...)
			i--
		}

	}
	return strings.Join(fields, " ")
}
