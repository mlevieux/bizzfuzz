package main

import (
	"strconv"
	"strings"
)

func fizzBuzz(d1, d2 int, limit int, str1, str2 string) string {

	var (
		result string
	)

	for i := 1 ; i <= limit ; i++ {
		if i % d1 == 0 && i % d2 == 0 {
			result += str1+str2
		} else if i % d1 == 0 {
			result += str1
		} else if i % d2 == 0 {
			result += str2
		} else {
			result += strconv.Itoa(i)
		}

		result += ","
	}

	return strings.TrimRight(result, ",")
}
