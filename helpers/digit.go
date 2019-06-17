package helpers

func ComputeSumDigit(n uint64) uint64 {
	sum := uint64(0)
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	return sum
}

func ContainsUInt64(s []uint64, v uint64) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}
