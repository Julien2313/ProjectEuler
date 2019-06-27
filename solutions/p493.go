package solution

import (
	"fmt"
)

const NBRCOLORS = 3
const NBRPICKS = 4
const NBRBALLSPERCOLOR = 2

func P493() float64 {
	balls := []int{}
	for i := 0; i < NBRCOLORS; i++ {
		balls = append(balls, NBRBALLSPERCOLOR)
	}
	nbrColors, nbrCases := RecurP493(1, 0, balls)
	fmt.Println(nbrColors, nbrCases)
	return float64(nbrColors) / float64(nbrCases)
}

// nbrCouleurs/nbrLancés
func RecurP493(coef int, nbrPicks int, balls []int) (int, int) {
	nbrColor := 0
	if nbrPicks == NBRPICKS {
		for color := 0; color < NBRCOLORS; color++ {
			if balls[color] != NBRBALLSPERCOLOR {
				nbrColor++
			}
		}
		fmt.Println(balls)
		return nbrColor, 1
	}
	nbrPicks++
	totCoef := 0
	for color := 0; color < NBRCOLORS; color++ {
		if balls[color] == 0 {
			continue
		}
		balls[color]--
		newNbrColor, newCoef := RecurP493(1, nbrPicks, balls)
		balls[color]++

		totCoef += newCoef * balls[color]
		nbrColor += newNbrColor * balls[color]
	}
	coef *= totCoef
	return nbrColor, coef
}
