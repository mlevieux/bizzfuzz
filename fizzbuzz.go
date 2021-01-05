package main

import (
	"math"
	"strings"
)

var (
	numSpace = map[int]int{
		10: 9,
		100: 180,
		1000: 2700,
		10000: 36000,
	}
)

func makeFizzBuzzBuffer(d1, d2, limit int, str1, str2 string) []byte {
	lenS1, lenS2 := len(str1), len(str2)
	lenLcm := lenS1 + lenS2
	lcm := lcm(d1, d2)
	oS1, oS2, oLcm := limit / d1, limit / d2, limit / lcm
	D := numberOfDigits(limit)
	A := float64(D) / float64(limit)

	bufferSizeS1 := float64(oS1 -oLcm) * (float64(lenS1) - A)
	bufferSizeS2 := float64(oS2 -oLcm) * (float64(lenS2) - A)
	bufferSizeLcm := float64(oLcm) * (float64(lenLcm) - A)
	nbCommas := limit - 1

	bufferSize := D + nbCommas + int(math.Ceil(bufferSizeS1 + bufferSizeS2 + bufferSizeLcm))
	return make([]byte, bufferSize)
}

func fizzBuzz(d1, d2 int, limit int, str1, str2 string) string {

	var (
		result string
	)

	buf := make([]byte, 11)
	g := newGen()
	for i := 1 ; i <= limit ; i++ {
		if i % d1 == 0 && i % d2 == 0 {
			result += str1+str2
		} else if i % d1 == 0 {
			result += str1
		} else if i % d2 == 0 {
			result += str2
		} else {
			n := g.fillNext(buf)
			result += string(buf[:n])
		}

		result += ","
		g.inc()
	}

	return strings.TrimRight(result, ",")
}