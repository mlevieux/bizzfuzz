package main

import (
	"math"
)

func approximateFizzBuzzBufferLen(d1, d2, limit int, str1, str2 string) int {
	lenS1, lenS2 := len(str1), len(str2)
	lenLcm := lenS1 + lenS2
	lcm := lcm(d1, d2)
	oS1, oS2, oLcm := limit/d1, limit/d2, limit/lcm
	D := numberOfDigits(limit)

	mpS1, _ := minimalPowerOf10(d1)
	mpS2, _ := minimalPowerOf10(d2)
	mpLcm, _ := minimalPowerOf10(lcm)

	AS1 := float64(D-numberOfDigits(lenS1-1)) / float64(limit-(mpS1-1))
	AS2 := float64(D-numberOfDigits(lenS2-1)) / float64(limit-(mpS2-1))
	ALcm := float64(D-numberOfDigits(lenLcm-1)) / float64(limit-(mpLcm-1))

	bufferSizeS1 := float64(oS1-oLcm) * (float64(lenS1) - AS1)
	bufferSizeS2 := float64(oS2-oLcm) * (float64(lenS2) - AS2)
	bufferSizeLcm := float64(oLcm) * (float64(lenLcm) - ALcm)
	nbCommas := limit - 1

	bufferSize := D + nbCommas + int(math.Ceil(bufferSizeS1+bufferSizeS2+bufferSizeLcm))
	return bufferSize
}

func fizzBuzz(d1, d2 int, limit int, str1, str2 string) []byte {

	buf := make([]byte, approximateFizzBuzzBufferLen(d1, d2, limit, str1, str2) + 100)
	n := 0
	len1, len2, len3 := len(str1), len(str2), len(str1) + len(str2)
	g := newGen()
	for i := 1; i <= limit; i++ {
		if i%d1 == 0 && i%d2 == 0 {
			copy(buf[n:], str1 + str2)
			n += len3
		} else if i%d1 == 0 {
			copy(buf[n:], str1)
			n += len1
		} else if i%d2 == 0 {
			copy(buf[n:], str2)
			n += len2
		} else {
			n += g.fillNext(buf[n:])
		}

		buf[n] = ','
		n++
		g.inc()
	}

	return buf[:n-1]
}
