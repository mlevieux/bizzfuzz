package main

var (
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

type gen struct {
	num []byte
	length int
}

func newGen() gen {
	g := gen{
		num:    make([]byte, 11),
		length: 1,
	}
	copy(g.num, num[:])

	return g
}

func (g *gen) reset() {
	g.length = 1
	copy(g.num, num[:])
}

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

func (g gen) fillNext(buf []byte) int {
	copy(buf, g.num[11-g.length:])
	return g.length
}
