package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func trivialFizzBuzz(d1, d2, limit int, s1, s2 string) string {
	result := ""
	for i := 1 ; i <= limit ; i++ {
		if i % d1 == 0 && i % d2 == 0 {
			result += s1 + s2
		} else if i % d1 == 0 {
			result += s1
		} else if i % d2 == 0 {
			result += s2
		} else {
			result += strconv.Itoa(i)
		}

		result += ","
	}
	result = strings.Trim(result, ",")
	return result
}

func testApproximateFizzBuzzBufferLen(d1, d2, limit int, s1, s2 string) bool {
	bufLen := approximateFizzBuzzBufferLen(d1, d2, limit, s1, s2)
	tfb := trivialFizzBuzz(d1, d2, limit, s1, s2)
	if bufLen < len(tfb) - 10 || bufLen > len(tfb) + 10 {
		fmt.Println(d1, d2, limit, s1, s2, "should produce", len(tfb), "but produced", bufLen)
		return false
	}
	return true
}

func randStr(rg *rand.Rand, min, max int) string {
	l := rg.Intn(max-min)+min
	return strings.Repeat("a", l)
}

func TestApproximateFizzBuzzBufferLen(t *testing.T) {
	sets := make([]fizzBuzzTestSet, 0)
	rg := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0 ; i < 10 ; i++ {
		sets = append(sets, fizzBuzzTestSet{
			d1:     rg.Intn(500) + 10,
			d2:     rg.Intn(500) + 10,
			limit:  rg.Intn(15000) + 10,
			str1:   randStr(rg, 1, 10),
			str2:   randStr(rg, 1, 10),
		})
	}

	for _, set := range sets {
		if !testApproximateFizzBuzzBufferLen(set.d1, set.d2, set.limit, set.str1, set.str2) {
			t.Fail()
		}
	}
}

type fizzBuzzTestSet struct {
	d1, d2 int
	limit int
	str1, str2 string

	result string
}

func TestFizzBuzz(t *testing.T) {

	sets := []fizzBuzzTestSet{
		{
			d1: 1,
			d2: 3,
			limit: 15,
			str1: "fizz",
			str2: "buzz",
			result: "fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz",
		},
		{
			d1: 3,
			d2: 3,
			limit: 15,
			str1: "fizz",
			str2: "buzz",
			result: "1,2,fizzbuzz,4,5,fizzbuzz,7,8,fizzbuzz,10,11,fizzbuzz,13,14,fizzbuzz",
		},
		{
			d1: 7,
			d2: 3,
			limit: 15,
			str1: "fizz",
			str2: "buzz",
			result: "1,2,buzz,4,5,buzz,fizz,8,buzz,10,11,buzz,13,fizz,buzz",
		},
		{
			d1: 21,
			d2: 3,
			limit: 15,
			str1: "fizz",
			str2: "buzz",
			result: "1,2,buzz,4,5,buzz,7,8,buzz,10,11,buzz,13,14,buzz",
		},
		{
			d1: 16,
			d2: 16,
			limit: 15,
			str1: "fizz",
			str2: "buzz",
			result: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15",
		},
	}

	for _, set := range sets {
		fizzBuzzResult := fizzBuzz(set.d1, set.d2, set.limit, set.str1, set.str2)
		if string(fizzBuzzResult) != set.result {
			t.Logf("With values (d1: %d, d2: %d, limit: %d, str1: %s, str2: %s)\n\tresult should be:\n%s\n\tbut is:\n%s\n", set.d1, set.d2, set.limit, set.str1, set.str2, set.result, fizzBuzzResult)
			t.Fail()
		}
	}
}

func BenchmarkFizzBuzz3_5_500_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_50_500_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 50, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_5_5000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 5000, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_50_5000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 50, 5000, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz30_5_5000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(30, 5, 5000, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_5_50000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 50000, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz300_500_50000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(300, 500, 50000, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_500_50000_40_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 500, 50000, "fizzfizzfizzfizzfizzfizzfizzfizz", "buzz")
	}
}