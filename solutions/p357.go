package solution

import (
	"fmt"
	"math"
	"math/big"

	"github.com/kavehmz/prime"
)

func P357() {
	sum := int64(0)
	primes := prime.Primes(100000000)
	for _, prime := range primes {
		n := int64(prime - 1)
		isOk := true
		max := int64(math.Sqrt(float64(n))) + 1
		for d := int64(2); d < max; d++ {
			if n%d == 0 {
				if !big.NewInt(d + n/d).ProbablyPrime(0) {
					isOk = false
					break
				}
			}
		}
		if isOk {
			sum += n
		}
	}
	fmt.Println(sum)
}
