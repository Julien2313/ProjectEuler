package solution

import (
	"fmt"
	"math"
)

const NBRCOLORS = 7
const NBRPICKS = 20
const NBRBALLSPERCOLOR = 10

func P493() float64 {
	fmt.Println(P493Fast())
	balls := []uint64{}
	for i := 0; i < NBRCOLORS; i++ {
		balls = append(balls, NBRBALLSPERCOLOR)
	}
	nbrColors, nbrCases, impr := RecurP493(0, balls)
	fmt.Println(nbrColors, nbrCases, impr)
	return float64(nbrColors) / float64(nbrCases)
}

func P493Fast() float64 {
	frac := float64(1)
	for x := 0; x < NBRPICKS; x++ {
		frac *= float64(NBRBALLSPERCOLOR*(NBRCOLORS-1)-x) / float64(NBRBALLSPERCOLOR*NBRCOLORS-x)
	}
	return float64(7 * (1 - frac))
}

// nbrCouleurs/nbrLancés/div
func RecurP493(nbrPicks int, balls []uint64) (uint64, uint64, int) {
	div := 0
	nbrColor := uint64(0)
	if nbrPicks == NBRPICKS {
		for color := 0; color < NBRCOLORS; color++ {
			if balls[color] != NBRBALLSPERCOLOR {
				nbrColor++
			}
		}
		return nbrColor, 1, 0
	}
	nbrPicks++
	coef := uint64(0)
	futurCoef := uint64(1)

	for color := 0; color < NBRCOLORS; color++ {
		if balls[color] == 0 {
			continue
		}
		if color+1 < NBRCOLORS && balls[color+1] == balls[color] {
			futurCoef++
		} else {
			balls[color]--
			newNbrColor, newCoef, newDiv := RecurP493(nbrPicks, balls)
			balls[color]++

			if newDiv > div {
				coef /= uint64(math.Pow(float64(newDiv-div), 2))
				nbrColor /= uint64(math.Pow(float64(newDiv-div), 2))
				div = newDiv
			} else if newDiv < div {
				newCoef /= uint64(math.Pow(float64(div-newDiv), 2))
				newNbrColor /= uint64(math.Pow(float64(div-newDiv), 2))
			}

			coef += newCoef * balls[color] * futurCoef
			nbrColor += newNbrColor * balls[color] * futurCoef
			futurCoef = 1.0
		}
	}

	for nbrColor > 100000000000000000 || coef > 100000000000000000 {
		nbrColor /= 2
		coef /= 2
		div++
	}
	return nbrColor, coef, div
}
