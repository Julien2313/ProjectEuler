package solution

import (
	"fmt"
	"math"

	"github.com/Julien2313/ProjectEuler/helpers"
)

const mod = 1000000007

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
	return sum % mod
}

func Bn(n uint64) uint64 {
	B := uint64(1)
	for k := uint64(1); k < n/2; k++ {
		B *= helpers.FactorialsXDivXMinusYDivY(n, k)
	}
	r := helpers.FactorialsXDivXMinusYDivY(n, (n / 2))

	if n%2 == 0 {
		return B * B * r
	}
	return B * B * r * r
}

func Dn(n uint64) uint64 {
	fmt.Println(Bn(n))
	return SumDivisorP650(Bn(n)) % mod
}

func Sn(n uint64) uint64 {
	sum := uint64(0)
	for k := uint64(1); k <= n; k++ {
		sum += Dn(k)
		sum %= mod
	}
	return sum
}

func P650() uint64 {

	// x := uint64(50)
	// y := uint64(24)

	// fmt.Println(helpers.Factorial(x) / helpers.Factorial(x-y) / helpers.Factorial(y))
	// fmt.Println(helpers.FactorialsXDivXMinusY(x, y) / helpers.Factorial(y))
	// fmt.Println(helpers.FactorialsXDivXMinusYDivY(x, y))

	// return 0
	return Sn(13)
}
