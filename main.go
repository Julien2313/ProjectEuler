package main

import (
	"fmt"
	"time"

	"github.com/Julien2313/ProjectEuler/solutions"
)

func main() {

	start := time.Now()
	solution.P357()
	fmt.Println("time took", time.Since(start))

}
