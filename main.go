package main

import (
	"fmt"
	"time"

	solution "github.com/Julien2313/ProjectEuler/solutions"
)

func main() {

	start := time.Now()
	fmt.Println(solution.P345())
	fmt.Println("time took", time.Since(start))

}
