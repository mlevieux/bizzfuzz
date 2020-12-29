package main

import "testing"

type fizzBuzzTestSet struct {
	d1, d2 int
	limit int
	s1, s2 string

	result string
}

func TestFizzBuzz(t *testing.T) {

	sets := []fizzBuzzTestSet{
		{
			d1: 1,
			d2: 3,
			limit: 15,
			s1: "fizz",
			s2: "buzz",
			result: "fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz",
		},
		{
			d1: 3,
			d2: 3,
			limit: 15,
			s1: "fizz",
			s2: "buzz",
			result: "1,2,fizzbuzz,4,5,fizzbuzz,7,8,fizzbuzz,10,11,fizzbuzz,13,14,fizzbuzz",
		},
		{
			d1: 7,
			d2: 3,
			limit: 15,
			s1: "fizz",
			s2: "buzz",
			result: "1,2,buzz,4,5,buzz,fizz,8,buzz,10,11,buzz,13,fizz,buzz",
		},
		{
			d1: 21,
			d2: 3,
			limit: 15,
			s1: "fizz",
			s2: "buzz",
			result: "1,2,buzz,4,5,buzz,7,8,buzz,10,11,buzz,13,14,buzz",
		},
		{
			d1: 16,
			d2: 16,
			limit: 15,
			s1: "fizz",
			s2: "buzz",
			result: "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15",
		},
	}

	for _, set := range sets {
		fizzBuzzResult := fizzBuzz(set.d1, set.d2, set.limit, set.s1, set.s2)
		if fizzBuzzResult != set.result {
			t.Logf("With values (d1: %d, d2: %d, limit: %d, s1: %s, s2: %s)\n\tresult should be:\n%s\n\tbut is:\n%s\n", set.d1, set.d2, set.limit, set.s1, set.s2, set.result, fizzBuzzResult)
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
		fizzBuzz(3, 5, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_5_5000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_50_5000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz30_5_5000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_5_50000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz300_500_50000_4_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 500, "fizz", "buzz")
	}
}

func BenchmarkFizzBuzz3_500_50000_40_4(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		fizzBuzz(3, 5, 500, "fizzfizzfizzfizzfizzfizzfizzfizz", "buzz")
	}
}
