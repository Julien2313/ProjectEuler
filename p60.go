package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"

	"github.com/julien2313/ProjectEuler/helpers"
	"github.com/kavehmz/prime"
)

const NBRPRIMES = 5

type P60 struct {
	Sum    uint64
	primes []uint64
}

func concatNumber(n1, n2 uint64) uint64 {
	return n1*uint64(math.Pow10(int(math.Log10(float64(n2))+1))) + n2
}

func test(num, nbr, deep, lenPrime int, combinations *[]P60, primes []uint64, primeToStore []uint64) {
	deep++
	for p := num + 1; p < lenPrime; p++ {
		primeToStore[deep] = primes[p]
		if deep == NBRPRIMES-1 {
			sum := uint64(0)
			for _, num := range primeToStore {
				sum += num
			}
			(*combinations)[nbr].primes = primeToStore
			(*combinations)[nbr].Sum = sum
			nbr++
		} else {
			test(p, nbr, deep, lenPrime, combinations, primes, primeToStore)
		}
	}
}

func generateAllCombinations(primes []uint64) []P60 {
	nbrPrimes := int64(len(primes))
	combinations := make([]P60, helpers.FactorialsXDivXMinusY(nbrPrimes, NBRPRIMES)/helpers.Factorial(NBRPRIMES))
	// test(-1, 0, 0, len(primes), &combinations, primes, make([]uint64, NBRPRIMES))
	nbr := 0
	for p1 := 0; p1 < len(primes); p1++ {
		for p2 := p1 + 1; p2 < len(primes); p2++ {
			for p3 := p2 + 1; p3 < len(primes); p3++ {
				for p4 := p3 + 1; p4 < len(primes); p4++ {
					for p5 := p4 + 1; p5 < len(primes); p5++ {
						combinations[nbr].primes = []uint64{primes[p1], primes[p2], primes[p3], primes[p4], primes[p5]}
						combinations[nbr].Sum = uint64(primes[p1] + primes[p2] + primes[p3] + primes[p4] + primes[p5])
						nbr++
					}
				}
			}
		}
	}
	return combinations
}

func checkCombinations(combinations []P60) {
	for i, _ := range combinations {

		isOk := true

	Loop:
		for p1 := 0; p1 < NBRPRIMES; p1++ {
			for p2 := p1 + 1; p2 < NBRPRIMES; p2++ {
				// fmt.Println(combinations[i])
				primeToCheck := concatNumber(combinations[i].primes[p1], combinations[i].primes[p2])

				if !big.NewInt(int64(primeToCheck)).ProbablyPrime(1) {
					isOk = false
					break Loop
				}
				primeToCheck = concatNumber(combinations[i].primes[p2], combinations[i].primes[p1])

				if !big.NewInt(int64(primeToCheck)).ProbablyPrime(1) {
					isOk = false
					break Loop
				}
			}
		}
		if isOk {
			fmt.Print("Found : ")
			fmt.Println(combinations[i].Sum)
			return
		}
	}
}

func p60() {
	primes := prime.Primes(700)
	combinations := generateAllCombinations(primes)

	sort.Slice(combinations, func(i, j int) bool {
		return combinations[i].Sum < combinations[j].Sum
	})
	fmt.Println("starting checking concat...")
	checkCombinations(combinations)
}
