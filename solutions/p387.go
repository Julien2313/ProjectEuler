package solution

import (
	"fmt"
	"math"
	"math/big"

	"github.com/Julien2313/ProjectEuler/helpers"
	"github.com/kavehmz/prime"
)

func isHarshad(harshad uint64) bool {

	for ; harshad > 10; harshad /= 10 {
		sumDigits := helpers.ComputeSumDigit(harshad)
		if harshad%sumDigits != uint64(0) {
			return false
		}
	}
	return true
}

func isStrongRightTruncatable(harshad uint64) bool {

	for ; harshad > 10; harshad /= 10 {
		sumDigits := helpers.ComputeSumDigit(harshad)
		if harshad%sumDigits != uint64(0) {
			return false
		}
	}
	return true
}

func P387() uint64 {
	sum := uint64(0)
	for d := uint64(1); d < 10; d++ {
		sum += recurP387(d)
	}
	// fmt.Println(recurP387(2011))
	return sum
}
func recurP387(harshad uint64) uint64 {
	if int(math.Log10(float64(harshad))) >= 4 {
		return 0
	}
	sum := uint64(0)
	sumDigits := helpers.ComputeSumDigit(harshad / 10)
	if sumDigits != 0 && harshad/10%sumDigits == uint64(0) {
		for d := uint64(0); d < 10; d++ {
			sum += recurP387(harshad*10 + d)
		}
	}
	if big.NewInt(int64(harshad)).ProbablyPrime(0) {
		return sum + harshad
	}
	return sum
}

func P387_() {
	sum := uint64(0)
	primes := prime.Primes(1000000000)
	harshads := [15][]uint64{}
	for _, prime := range primes {
		harshad := prime / 10
		sumDigits := helpers.ComputeSumDigit(harshad)
		if sumDigits == 0 {
			continue
		}
		if harshad%sumDigits != uint64(0) {
			continue
		}
		isOk := true
		for ; harshad > 10; harshad /= 10 {
			sumDigits := helpers.ComputeSumDigit(harshad)
			nbrDigits := int(int(math.Log10(float64(prime))))
			if helpers.BinarySearch(harshads[nbrDigits], harshad) {
				break
			}
			if harshad%sumDigits != uint64(0) {
				isOk = false
				break
			}
		}
		if !big.NewInt(int64(prime / 10 / sumDigits)).ProbablyPrime(0) {
			continue
		}

		if isOk {
			nbrDigits := int(int(math.Log10(float64(prime))))
			harshads[nbrDigits] = append(harshads[nbrDigits], prime)
			sum += prime
		}
	}
	fmt.Println(sum)
}
