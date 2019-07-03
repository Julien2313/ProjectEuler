package helpers

import (
	"math"
)

func Factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func Sub2Factorials(n1, n2 uint64) uint64 {
	if n1 < n2 {
		n3 := n2
		n2 = n1
		n1 = n3
	}

	diff := n1
	for i := uint64(1); i < n1-n2; i++ {
		diff *= n1 - i
	}
	return (diff - 1) * Factorial(n2)
}

// x!/(x-y)!
func FactorialsXDivXMinusY(x, y uint64) uint64 {
	if x < y {
		panic("x < y")
	}
	mult := uint64(1)
	for i := uint64(0); i < y; i++ {
		mult *= x - i
	}
	return mult
}

// x!/(x-y)!/y!
func FactorialsXDivXMinusYDivY(x, y uint64) uint64 {
	if x < y {
		panic("x < y")
	}
	fact := uint64(1)
	xmy := x - y
	if x < y*2 {
		t := xmy
		xmy = y
		y = t
	}
	Fy := Factorial(y)
	for ; x > xmy; x-- {
		fact *= x
		if fact%Fy == 0 {
			fact /= Fy
			Fy = 1
		}
	}

	return fact / Fy
}

func IsPrimeSqrt(value uint64) bool {
	for i := uint64(2); i <= uint64(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == uint64(0) {
			return false
		}
	}
	return value > 1
}

func IsPrimeSqrtInt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func AllDivisor(n int) []int {
	divisors := []int{1, n}
	max := int(math.Sqrt(float64(n))) + 1
	for d := 2; d < max; d++ {
		if n%d == 0 {
			divisors = append(divisors, d, n/d)
		}
	}

	return divisors
}
func Gcd(x, y uint64) uint64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
