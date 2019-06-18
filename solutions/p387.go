package solution

import (
	"fmt"
	"math/big"

	"github.com/Julien2313/ProjectEuler/helpers"
	"github.com/kavehmz/prime"
)

func P387() {
	sum := uint64(90619)
	// primes := prime.Primes(100000000000000)
	primes := prime.Primes(1000000000)
	var harshads []uint64
	for x := uint64(1000); x < 10000; x++ {
		sumDigits := helpers.ComputeSumDigit(x)
		if x%sumDigits == uint64(0) {
			harshads = append(harshads, x)
		}
	}

	for _, prime := range primes[1229:] {
		if !helpers.ContainsUInt64(harshads, helpers.GetHigherDigits(prime, 4)) {
			continue
		}

		harshad := prime / 10
		sumDigits := helpers.ComputeSumDigit(harshad)
		if harshad%sumDigits != uint64(0) {
			continue
		}
		if !big.NewInt(int64(harshad / sumDigits)).ProbablyPrime(0) {
			continue
		}
		isOk := true
		for ; harshad > 10; harshad /= 10 {
			sumDigits := helpers.ComputeSumDigit(harshad)
			if harshad%sumDigits != uint64(0) {
				isOk = false
				break
			}
		}
		if isOk {
			sum += prime
		}
	}
	fmt.Println(sum)
}
