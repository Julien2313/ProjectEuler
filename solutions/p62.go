package solution

import (
	"fmt"
	"math"
)

const maxNbr = 100000

func checkSameDigits(digits [10]int, n1, n2 uint64) bool {
	// if int(math.Log10(float64(n1))) != int(math.Log10(float64(n2))) {
	// 	return false
	// }

	for n1 != 0 {
		digits[n1%10] += 1
		digits[n2%10] -= 1
		n1 /= 10
		n2 /= 10
	}
	for _, value := range digits {
		if value != 0 {
			return false
		}
	}
	return true
}

func P62() {
	var digits [10]int
	for n1 := 0; n1 < maxNbr; n1++ {
		cube1 := uint64(n1 * n1 * n1)
		for n2 := n1 + 1; n2 < maxNbr; n2++ {
			cube2 := uint64(n2 * n2 * n2)
			if int(math.Log10(float64(cube2))) != int(math.Log10(float64(cube1))) {
				break
			}
			if !checkSameDigits(digits, cube2, cube1) {
				continue
			}
			for n3 := n2 + 1; n3 < maxNbr; n3++ {
				cube3 := uint64(n3 * n3 * n3)
				if int(math.Log10(float64(cube3))) != int(math.Log10(float64(cube2))) {
					break
				}
				if !checkSameDigits(digits, cube3, cube1) {
					continue
				}
				if !checkSameDigits(digits, cube3, cube2) {
					continue
				}
				for n4 := n3 + 1; n4 < maxNbr; n4++ {
					cube4 := uint64(n4 * n4 * n4)
					if int(math.Log10(float64(cube4))) != int(math.Log10(float64(cube3))) {
						break
					}
					if !checkSameDigits(digits, cube4, cube1) {
						continue
					}
					if !checkSameDigits(digits, cube4, cube2) {
						continue
					}
					if !checkSameDigits(digits, cube4, cube3) {
						continue
					}
					for n5 := n4 + 1; n5 < maxNbr; n5++ {
						cube5 := uint64(n5 * n5 * n5)
						if int(math.Log10(float64(cube5))) != int(math.Log10(float64(cube4))) {
							break
						}
						if !checkSameDigits(digits, cube5, cube1) {
							continue
						}
						if !checkSameDigits(digits, cube5, cube2) {
							continue
						}
						if !checkSameDigits(digits, cube5, cube3) {
							continue
						}
						if !checkSameDigits(digits, cube5, cube4) {
							continue
						}
						fmt.Println(cube1)
						return
					}
				}
			}
		}
	}
}
