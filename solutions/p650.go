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
	divisors := n + 1
	max := uint64(math.Sqrt(float64(n))) + 1
	for d := uint64(2); d < max; d++ {
		if n%d == 0 {
			if d != n/d {
				divisors += n / d
			}
			divisors += d
		}
	}
	return divisors
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
	}
	return sum
}

func P650() float64 {
	return float64(Sn(5))
}
