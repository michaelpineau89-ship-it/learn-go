package main

import (
	"strings"
)

func toWeirdCase(str string) string {
	r := ""
	count := 0
	characters := strings.Split(str, "")
	for i := 0; i < len(characters); i++ {
		if characters[i] != " " {
			if count%2 == 0 {
				r += strings.ToUpper(characters[i])
			} else {
				r += strings.ToLower(characters[i])
			}
			count += 1
		} else {
			r += " "
			count = 0
		}
	}
	return r
}
