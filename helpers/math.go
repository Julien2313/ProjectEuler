package helpers

func Factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func Sub2Factorials(n1, n2 int64) int64 {
	if n1 < n2 {
		n3 := n2
		n2 = n1
		n1 = n3
	}

	diff := n1
	for i := int64(1); i < n1-n2; i++ {
		diff *= n1 - i
	}
	return (diff - 1) * Factorial(n2)
}

// x!/(x-y)!
func FactorialsXDivXMinusY(x, y int64) int64 {
	if x < y {
		panic("x < y")
	}

	mult := int64(1)
	for i := int64(0); i < y; i++ {
		mult *= x - i
	}
	return mult
}
