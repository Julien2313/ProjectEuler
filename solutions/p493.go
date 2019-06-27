package solution

import (
	"fmt"
)

const NBRCOLORS = 7
const NBRPICKS = 10
const NBRBALLSPERCOLOR = 10

func P493() float64 {
	balls := []uint64{}
	for i := 0; i < NBRCOLORS; i++ {
		balls = append(balls, NBRBALLSPERCOLOR)
	}
	nbrColors, nbrCases := RecurP493(1, 0, balls)
	fmt.Println(nbrColors, nbrCases)
	return float64(nbrColors) / float64(nbrCases)
}

// nbrCouleurs/nbrLancés
func RecurP493(coef uint64, nbrPicks int, balls []uint64) (uint64, uint64) {
	nbrColor := uint64(0)
	if nbrPicks == NBRPICKS {
		for color := 0; color < NBRCOLORS; color++ {
			if balls[color] != NBRBALLSPERCOLOR {
				nbrColor++
			}
		}
		// fmt.Println(balls)
		return nbrColor, 1
	}
	nbrPicks++
	totCoef := uint64(0)
	futurCoef := uint64(1)
	for color := 0; color < NBRCOLORS; color++ {
		if balls[color] == 0 {
			continue
		}
		if color+1 == NBRCOLORS {
			balls[color]--

			newNbrColor, newCoef := RecurP493(1, nbrPicks, balls)
			balls[color]++
			totCoef += newCoef * balls[color] * futurCoef
			nbrColor += newNbrColor * balls[color] * futurCoef
			futurCoef = 1.0
		} else {
			if balls[color+1] == balls[color] {
				futurCoef++
			} else {
				balls[color]--
				newNbrColor, newCoef := RecurP493(1, nbrPicks, balls)
				balls[color]++

				totCoef += newCoef * balls[color] * futurCoef
				nbrColor += newNbrColor * balls[color] * futurCoef
				futurCoef = 1.0
			}
		}
	}
	coef *= totCoef
	// gdc := helpers.Gcd(nbrColor, coef)
	// nbrColor /= gdc
	// coef /= gdc
	return nbrColor, coef
}
