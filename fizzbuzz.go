package main

import (
	"strconv"
	"strings"
)

func fizzBuzz(d1, d2 int, limit int, s1, s2 string) string {

	var (
		result string
	)

	for i := 1 ; i <= limit ; i++ {
		if i % d1 == 0 && i % d2 == 0 {
			result += s1+s2
		} else if i % d1 == 0 {
			result += s1
		} else if i % d2 == 0 {
			result += s2
		} else {
			result += strconv.Itoa(i)
		}

		result += ","
	}

	return strings.TrimRight(result, ",")
}
