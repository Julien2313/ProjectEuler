package solution

import (
	"fmt"
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

func P387() {
	sum := uint64(0)
	primes := prime.Primes(100000000)
	// good := make(map[uint64]bool)
	for i, j := 0, len(primes)-1; i < j; i, j = i+1, j-1 {
		primes[i], primes[j] = primes[j], primes[i]
	}
	for _, prime := range primes {
		harshad := prime / 10
		// if value, ok := good[harshad]; ok && !value {
		// 	continue
		// }
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
			if harshad%sumDigits != uint64(0) {
				isOk = false
				// good[harshad] = false
				break
			}
		}
		if !big.NewInt(int64(prime / 10 / sumDigits)).ProbablyPrime(0) {
			continue
		}

		if isOk {
			// for harshad := prime / 10; harshad > 10; harshad /= 10 {
			// 	good[harshad] = true
			// }
			sum += prime
		}
	}
	fmt.Println(sum)
}
