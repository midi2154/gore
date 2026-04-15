package transformations

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Bin(text string) string {
	fields := strings.Fields(text)

	for i := 0; i < len(fields); i++ {
		if fields[i] == "(bin)" && i > 0 {
			decimal, err := strconv.ParseInt(fields[i-1], 2, 64)
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
