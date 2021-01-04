package main

var (
	numbers = [...]byte{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
	}

	length = 1
	num = func () [11]byte { b := [11]byte{} ; for i := 0 ; i < 10 ; i++ {b[i] = 48} ; b[10] = 49 ; return b }()
)

func reset() {
	length = 1
}

func  inc() {
	for j := 10 ; j > 10 - length ; j-- {
		if num[j] == 57 {
			num[j] = 48
		} else {
			num[j]++
			return
		}
	}
	length++
}

func  next() string {
	return string(num[11-length:])
}
