package helpers

import "math"

func ComputeSumDigits(n uint64) uint64 {
	sum := uint64(0)
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	return sum
}

func getHigherDigit(n uint64) uint64 {
	if int(math.Pow10(int(math.Log10(float64(n))))) <= 1 {
		return n
	}
	return n / uint64(math.Pow10(int(math.Log10(float64(n)))))
}

func GetHigherDigits(n uint64, nbrDigits int) uint64 {
	// if int(math.Log10(float64(n))) <= nbrDigits {
	// 	return n
	// }
	return n / uint64(math.Pow10(int(math.Log10(float64(n)))-nbrDigits+1))
}
