package solution

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func getHigherDigit(n int) int {
	if int(math.Pow10(int(math.Log10(float64(n))))) <= 1 {
		return n
	}
	return n / int(math.Pow10(int(math.Log10(float64(n)))))
}

func removeHigherDigit(n int) int {
	if int(math.Pow10(int(math.Log10(float64(n))))) <= 1 {
		return -1
	}
	return n % int(math.Pow10(int(math.Log10(float64(n)))))
}

func P79() {
	var passPhrase []int
	file, err := os.Open("solution/p79input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 4)

	for i := 0; i < 50; i++ {
		_, err := file.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		buf, err := strconv.Atoi(string(buf)[0 : len(buf)-1])
		if err != nil {
			fmt.Println(err)
			return
		}
		passPhrase = append(passPhrase, buf)
	}

	var before [10][10]bool
	var after [10][10]bool
	for _, pass := range passPhrase {
		d1 := pass / 100
		d2 := (pass / 10) % 10
		d3 := pass % 10
		after[d1][d2] = true
		after[d1][d3] = true

		before[d2][d1] = true
		after[d2][d3] = true

		before[d3][d1] = true
		before[d3][d2] = true

	}

	for i, b := range before {
		fmt.Print(i, ": ")
		for j, k := range b {
			if k {
				fmt.Print(j, ", ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	for i, a := range after {
		fmt.Print(i, ": ")
		for j, k := range a {
			if k {
				fmt.Print(j, ", ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {

			}
		}
	}
}
