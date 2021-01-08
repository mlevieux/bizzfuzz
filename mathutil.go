package main

// we know that in ascending order, we have:
// [NOTE: we don't include 0 for this exercise]
// 9 numbers of 1 digit: 1, 2, 3, 4, 5, 6, 7, 8, 9
// 90 numbers of 2 digits: 10, 11, ..., 20, 21, ..., 30, ..., 40, ..., 98, 99
// 900 numbers of 3 digits: 100, 101, ..., 999
// ...
//
// This is precisely because in base 10, any number N will have M+1 digits, where
// M is such that 10^M < N < 10^(M+1)
//
// A good way to find the buffer size we need for a given range is:
// for i starting at M, until range is 0:
//   bsize += (range - 10^i + 1)(i+1)
//   range = 10^i - 1
//
// Also, notice that by construction, divisors will all be equally spread
// among all numbers, and since bsize (above) is the sum of all numbers of
// digits of all numbers to be printed, bsize / range is the average number
// of digits for the same range of numbers.

// numberOfDigits returns the number of digits needed to store a sequence
//  123...(n-1)n where n is an arbitrary integer
func numberOfDigits(n int) int {
	p, c := minimalPowerOf10(n)

	d := 0
	for n > 0 {
		d += (n - p + 1) * (c + 1)
		n = p - 1

		p /= 10
		c -= 1
	}

	return d
}

// minimalPowerOf10 returns the minimum power of 10 p
// such that p < n < p*10
// It returns p, as well as the number of characters needed to
// write p
func minimalPowerOf10(n int) (int, int) {
	p := 1
	c := 0
	for ; p*10 <= n; p, c = p*10, c+1 {
	}
	return p, c
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func lcm(a, b int) int {
	p := a * b
	if p < 0 {
		p = -p
	}
	return p / gcd(a, b)
}