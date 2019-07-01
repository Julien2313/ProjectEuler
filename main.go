package main

import (
	"fmt"
	"time"

	solution "github.com/Julien2313/ProjectEuler/solutions"
)

func main() {

	start := time.Now()
	fmt.Println(solution.P650())
	fmt.Println("time took", time.Since(start))

}
