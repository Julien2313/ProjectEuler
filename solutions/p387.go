package solution

import (
	"fmt"
	"math/big"

	"github.com/Julien2313/ProjectEuler/helpers"
	"github.com/kavehmz/prime"
)

func P387() {
	sum := uint64(0)
	// primes := prime.Primes(100000000000000)
	primes := prime.Primes(100000000)
	// var strongHarshadNumbers []uint64
	// strongHarshadNumbers := make(map[uint64]bool)
	var harshads []uint64
	for _, prime := range primes {

		harshad := prime / 10
		h := []uint64{}
		sumDigits := helpers.ComputeSumDigit(harshad)
		if sumDigits == 0 {
			continue
		}
		if !big.NewInt(int64(harshad / sumDigits)).ProbablyPrime(0) {
			continue
		}
		isOk := true
		for ; harshad > 0; harshad /= 10 {
			h = append(h, harshad)
			if helpers.ContainsUInt64(harshads, harshad) {
				break
			}
			sumDigits := helpers.ComputeSumDigit(harshad)
			if harshad%sumDigits != uint64(0) {
				isOk = false
				break
			}
		}
		if isOk {
			sum += prime
			harshads = append(harshads, h...)
		}
	}
	fmt.Println(sum)
}
