package solution

import (
	"math"
	"math/big"

	"github.com/Julien2313/ProjectEuler/helpers"
)

func P387() uint64 {
	sum := uint64(0)
	for d := uint64(1); d < 10; d++ {
		sum += recurP387(d)
	}
	return sum
}

func recurP387(strongHarshad uint64) uint64 {
	sum := uint64(0)
	oldSumDigits := helpers.ComputeSumDigits(strongHarshad)

	if int(math.Log10(float64(strongHarshad)))+1 >= 14 {
		return 0
	}

	for d := uint64(0); d < 10; d++ {
		newStrongHarshad := strongHarshad*10 + d
		newSumDigits := helpers.ComputeSumDigits(newStrongHarshad)

		if big.NewInt(int64(newStrongHarshad)).ProbablyPrime(0) {
			if big.NewInt(int64(strongHarshad / oldSumDigits)).ProbablyPrime(0) {
				sum += newStrongHarshad
			}
		}
		if newStrongHarshad%newSumDigits == uint64(0) {
			sum += recurP387(newStrongHarshad)
		}
	}
	return sum
}
