package main

var (
	// num is a variable declared to speed up gen initialization
	num = [11]byte{
		48,
		48,
		48,
		48,
		48,
		48,
		48,
		48,
		48,
		48,
		49,
	}
)

// gen wraps data and behaviours for generation of contiguous
// numbers. It is used to take profit of the fact that in
// fizzbuzz exercise, all numbers are at least ranged over, even
// if a certain number of them are not "generated", as in output.
// It is designed to outperform strconv.Itoa in this specific
// contiguous number generation task (see benchmarks in cng_test.go).
type gen struct {
	num []byte
	length int
}

// newGen returns a new, freshly allocated gen
func newGen() gen {
	g := gen{
		num:    make([]byte, 11),
		length: 1,
	}
	copy(g.num, num[:])

	return g
}

// reset resets a gen
func (g *gen) reset() {
	g.length = 1
	copy(g.num, num[:])
}

// inc increments the gen counter so the next call
// to fillNext will fill the given buffer with the
// representation of the next number in ascending order
func (g *gen) inc() {
	j := 10
	for ; j >= 11 - g.length ; j-- {
		if g.num[j] == 57 {
			g.num[j] = 48
		} else {
			g.num[j]++
			return
		}
	}
	g.length++
	g.num[j]++
}

// fillNext takes a buffer and fills it with the representation
// of the current number gen holds. It does not perform any check
// on buf length so it's the caller's responsibility to ensure
// there is enough space in it.
// It returns the length of the number that has been written to buf
// in terms of number of bytes.
func (g gen) fillNext(buf []byte) int {
	copy(buf, g.num[11-g.length:])
	return g.length
}
