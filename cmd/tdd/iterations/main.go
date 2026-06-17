package iterations

import "strings"

func Repeat(character string, repititons int) string {
	var repeated strings.Builder
	for i := 0; i < repititons; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
