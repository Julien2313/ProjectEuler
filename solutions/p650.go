package solution

import (
	"math"
)

func FactP650(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return FactP650(n-1) * n
}

func SumDivisorP650(n uint64) uint64 {
	sum := uint64(1)
	for d := uint64(2); d <= uint64(math.Sqrt(float64(n))); d++ {

		dSum := uint64(1)
		dTerm := uint64(1)
		for n%d == 0 {
			n /= d

			dTerm *= d
			dSum += dTerm
		}
		sum *= dSum
	}

	if n >= 2 {
		sum *= (1 + n)
	}
	return sum
}

func Bn(n uint64) uint64 {
	B := uint64(1)
	for k := uint64(1); k < n/2; k++ {
		B *= FactP650(n) / FactP650(n-k) / FactP650(k)
	}
	r := FactP650(n) / FactP650(n-(n/2)) / FactP650(n/2)
	if n%2 == 0 {
		return B * B * r
	} else {
		return B * B * r * r
	}
}

func Dn(n uint64) uint64 {
	return SumDivisorP650(Bn(n))
}

func Sn(n uint64) uint64 {
	sum := uint64(0)
	for k := uint64(1); k <= n; k++ {
		sum += Dn(k)
		sum %= 1000000007
	}
	return sum
}

func P650() uint64 {
	return Sn(10)
}
