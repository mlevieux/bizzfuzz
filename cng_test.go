package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestContiguousNumberGenerator_Next(t *testing.T) {

	for n := next() ; n != "30" ; n = next() {
		fmt.Println(n)
		inc()
	}
}

func BenchmarkContiguousNumberGenerator_Next(b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		for j := 0 ; j <= 50000 ; j++ {
			_ = next()
			inc()
		}
		reset()
	}
}

func BenchmarkStrconv_Itoa(b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		for j := 1 ; j <= 50000 ; j++ {
			x := strconv.Itoa(j)
			if x[0] < 10 {
				return
			}
		}
	}
}