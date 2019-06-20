package solution

import (
	"fmt"
	"math"
)

const NBRBALLS = 3
const NBRPICKS = 2

func P493() uint64 {

	// 70 balles
	// 7 couleurs
	// 10 balles de chaque coulleur

	// 20 tirages sans remise
	/*
		1er, on tir forcément 1 nouvelle couleur
			[10, 10, 10, 10, 10, 10, 9]
		2eme, on a 2 possibilités, soit on tir la même couleur, soit on en tir une nouvelle
			[10, 10, 10, 10, 10, 10, 8] (9/69)

			[10, 10, 10, 10, 10, 9, 9] (60/69)
		3eme tirage, 3 possibilités
			[10, 10, 10, 10, 10, 10, 7] (8/68)

			[10, 10, 10, 10, 10, 9, 8] (10/68)

			[10, 10, 10, 10, 9, 9, 9] (50/68)
		4eme tirage, 4 possibilités
			[10, 10, 10, 10, 10, 10, 6] (/67)
			[10, 10, 10, 10, 10, 9, 7] (/67)

			[10, 10, 10, 10, 10, 8, 7] (/67)
			[10, 10, 10, 10, 10, 8, 7] (/67)
			[10, 10, 10, 8, 9, 9, 9] (/67)
	*/
	// fmt.Println(RecurP493(1.0, 0, []int{10, 10, 10, 10, 10, 10, 10, 10}))
	fmt.Println(float64(RecurP493(0, 0, []int{10, 10, 10})), math.Pow(NBRBALLS, NBRPICKS))
	return 0
}

// nbr couleur/nbrlLancés
func RecurP493(coef int, nbrPicks int, balls []int) int {
	nbrColor := 0
	if nbrPicks == NBRPICKS {
		// fmt.Println(balls)
		for color := 0; color < NBRBALLS; color++ {
			if balls[color] != 10 {
				nbrColor++
			}
		}
		fmt.Println("coef", nbrColor*coef)
		return nbrColor * coef
	}
	nbrPicks++
	futurCoef := 1
	for color := 0; color < NBRBALLS; color++ {
		if balls[color] == 0 {
			continue
		}
		if color+1 == NBRBALLS {
			balls[color]--

			newNbrColor := RecurP493(futurCoef, nbrPicks, balls)
			// nbrColor = (nbrColor*float64(coef) + newNbrColor*float64(newCoef)) / float64(newCoef+coef)
			nbrColor += newNbrColor

			balls[color]++
		} else {
			if balls[color+1] == balls[color] {
				futurCoef++
			} else {
				balls[color]--

				newNbrColor := RecurP493(futurCoef, nbrPicks, balls)
				// nbrColor = (nbrColor*float64(coef) + newNbrColor*float64(newCoef)) / float64(newCoef+coef)
				nbrColor += newNbrColor

				balls[color]++
				futurCoef = 1.0
			}
		}
	}
	return nbrColor
}
