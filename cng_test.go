package main

import (
	"strconv"
	"testing"
)

func TestContiguousNumberGenerator_Next(t *testing.T) {

	buf := make([]byte, 51)
	g := newGen()
	n := 0
	for n != len(buf){
		n += g.fillNext(buf[n:])
		g.inc()
	}

	if string(buf) != "123456789101112131415161718192021222324252627282930" {
		t.Log(string(buf))
		t.Fail()
	}
}

func BenchmarkContiguousNumberGenerator_Next(b *testing.B) {
	g := newGen()
	buf := make([]byte, 11)
	for i := 0 ; i < b.N ; i++ {
		g.reset()
		for j := 0 ; j <= 50000 ; j++ {
			g.fillNext(buf)
			g.inc()
		}
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