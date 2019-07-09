package solution

import (
	"fmt"
	"math"

	"github.com/Julien2313/ProjectEuler/helpers"
	"github.com/kavehmz/prime"
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
	return sum
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
	return SumDivisorP650(Bn(n))
}

func Sn(n uint64) uint64 {
	sum := uint64(0)
	for k := uint64(1); k <= n; k++ {
		sum += Dn(k)
	}
	return sum
}

func primeFactorsToFactors(primeFactors map[uint64]int64, value uint64, factors *[]uint64) {
	if len(primeFactors) == 0 {
		*factors = append(*factors, value)
		return
	}

	primeFactor := getSomeKey(primeFactors)
	newPrimeFactors := make(map[uint64]int64)
	for k, v := range primeFactors {
		newPrimeFactors[k] = v
	}

	delete(newPrimeFactors, primeFactor)

	for x := int64(0); x <= primeFactors[primeFactor]; x++ {
		if x != 0 {
			value *= primeFactor
		}
		primeFactorsToFactors(newPrimeFactors, value, factors)
	}
}

func getSomeKey(m map[uint64]int64) uint64 {
	for k := range m {
		return k
	}
	panic("no key left")
}

func Sn_(n uint64) uint64 {
	sum := uint64(0)

	primes := prime.Primes(n)
	for x := uint64(1); x <= n; x++ {
		primeFactors := make(map[uint64]int64)
		if x%2 == 0 {
			for i := uint64(0); i < x/2-1; i++ {
				primeFactors[x-i] += int64((x/2 - 1 - i) * 2)
			}
			for i := uint64(2); i < x/2; i++ {
				primeFactors[i] -= int64((x/2 - i) * 2)
			}
			for key, _ := range primeFactors {
				if primeFactors[key] > 0 {
					primeFactors[key]++
				} else {
					primeFactors[key]--
				}
			}
			primeFactors[x/2] = -1
			primeFactors[x/2+1] = 1
		} else {
			for i := uint64(0); i < x/2; i++ {
				primeFactors[x-i] += int64(x/2-i) * 2
			}
			for i := uint64(2); i <= x/2; i++ {
				primeFactors[i] -= int64(x/2-i+1) * 2
			}
		}
		for factor, _ := range primeFactors {
			if helpers.IsPrimeSqrt(factor) {
				continue
			}
			f := factor
			for _, prime := range primes {
				if prime > f {
					break
				}
				for f%prime == 0 {
					primeFactors[prime] += primeFactors[factor]
					f /= prime
				}
			}
			delete(primeFactors, factor)
		}
		factors := []uint64{}
		fmt.Println(x, primeFactors)
		continue
		primeFactorsToFactors(primeFactors, 1, &factors)
		for _, f := range factors {
			sum += f
		}
		sum %= mod
	}
	return sum % mod
}

func P650() int64 {
	x := uint64(10)
	// fmt.Println(Sn(x))
	fmt.Println(Sn_(x))
	return 0
}

//  S(5)=5736
