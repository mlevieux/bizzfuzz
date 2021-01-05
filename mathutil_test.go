package main

import "testing"

func testNumberOfDigits(n int, nbOfDigits int) bool {
	result := numberOfDigits(n)
	return result == nbOfDigits
}

type numberOfDigitsTestSet struct {
	number     int
	nbOfDigits int
}

func TestNumberOfDigits(t *testing.T) {
	sets := []numberOfDigitsTestSet{
		{
			number:     1,
			nbOfDigits: 1,
		},
		{
			number:     3,
			nbOfDigits: 3,
		},
		{
			number:     10,
			nbOfDigits: 11,
		},
		{
			number:     24,
			nbOfDigits: 39,
		},
		{
			number:     278,
			nbOfDigits: 726,
		},
		{
			number:     300,
			nbOfDigits: 792,
		},
	}

	for _, set := range sets {
		if !testNumberOfDigits(set.number, set.nbOfDigits) {
			t.Log(set.number, "-", set.nbOfDigits)
			t.Fail()
		}
	}
}

func testGcd(a, b, result int) bool {
	return gcd(a, b) == result
}

type gcdTestSet struct {
	a      int
	b      int
	result int
}

func TestGcd(t *testing.T) {
	sets := []gcdTestSet{
		{
			a:      15,
			b:      5,
			result: 5,
		},
		{
			a:      15,
			b:      85,
			result: 5,
		},
		{
			a:      153,
			b:      54,
			result: 9,
		},
		{
			a:      7,
			b:      89,
			result: 1,
		},
	}

	for _, set := range sets {
		if !testGcd(set.a, set.b, set.result) {
			t.Log(set.a, set.b, set.result)
			t.Fail()
		}
	}
}

func testLcm(a, b, result int) bool {
	return lcm(a, b) == result
}

type lcmTestSet struct {
	a int
	b int
	result int
}

func TestLcm(t *testing.T) {
	sets := []lcmTestSet{
		{
			a: 15,
			b: 5,
			result: 15,
		},
		{
			a: 21,
			b: 14,
			result: 42,
		},
		{
			a: 157,
			b: 59,
			result: 157*59,
		},
		{
			a: 9,
			b: 15,
			result: 45,
		},
	}

	for _, set := range sets {
		if !testLcm(set.a, set.b, set.result) {
			t.Log(set.a, set.b, set.result)
			t.Fail()
		}
	}
}