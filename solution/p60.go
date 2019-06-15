package solution

import (
	"fmt"
	"math"

	"github.com/Julien2313/ProjectEuler/helpers"
	"github.com/kavehmz/prime"
)

const NBRPRIMES = 5

type t_p60 struct {
	Sum    int
	primes []int
}

func concatNumber(n1, n2 uint64) uint64 {
	return n1*uint64(math.Pow10(int(math.Log10(float64(n2))+1))) + n2
}

func isConcatPrime(primes []uint64) bool {
	if len(primes) == 1 {
		return true
	}
	primetoCheck, primes := primes[0], primes[1:]
	for _, p := range primes {

		if !helpers.IsPrimeSqrt(concatNumber(primetoCheck, p)) {
			return false
		}
		if !helpers.IsPrimeSqrt(concatNumber(p, primetoCheck)) {
			return false
		}
	}
	return isConcatPrime(primes)
}

func P60Recur(last, depth int, primesChoosen []uint64, primes []uint64) uint64 {
	for p := last + 1; p < len(primes); p++ {
		primesChoosen[depth] = primes[p]
		if !isConcatPrime(primesChoosen) {
			continue
		}
		if depth == NBRPRIMES-1 {
			fmt.Println(primesChoosen)
			return primesChoosen[depth]
		}
		primesChoosen = append(primesChoosen, uint64(0))
		sum := P60Recur(p, depth+1, primesChoosen, primes)
		if sum != 0 {
			return sum + primesChoosen[depth]
		}
		primesChoosen = primesChoosen[0 : len(primesChoosen)-1]
	}
	return 0
}

func P60Iter(primes []uint64) uint64 {
	for p1 := 0; p1 < len(primes); p1++ {
		for p2 := p1 + 1; p2 < len(primes); p2++ {
			if !isConcatPrime([]uint64{primes[p1], primes[p2]}) {
				continue
			}
			for p3 := p2 + 1; p3 < len(primes); p3++ {
				if !isConcatPrime([]uint64{primes[p1], primes[p2], primes[p3]}) {
					continue
				}
				for p4 := p3 + 1; p4 < len(primes); p4++ {
					if !isConcatPrime([]uint64{primes[p1], primes[p2], primes[p3], primes[p4]}) {
						continue
					}
					for p5 := p4 + 1; p5 < len(primes); p5++ {
						if !isConcatPrime([]uint64{primes[p1], primes[p2], primes[p3], primes[p4], primes[p5]}) {
							continue
						}
						fmt.Println(primes[p1], primes[p2], primes[p3], primes[p4], primes[p5])
						return primes[p1] + primes[p2] + primes[p3] + primes[p4] + primes[p5]
					}
				}
			}
		}
	}
	return 0
}

func P60() {
	primes := prime.Primes(10000)

	fmt.Println(P60Recur(-1, 0, make([]uint64, 1), primes))
	fmt.Println(P60Iter(primes))

}
