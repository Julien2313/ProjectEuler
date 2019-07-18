package solution

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func P345() uint64 {

	file, err := os.Open("solutions/p345input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	board := make([][]uint64, len(lines))
	for i, _ := range board {
		board[i] = make([]uint64, len(lines))
	}

	splitFn := func(r rune) bool {
		return r == ' '
	}

	for y, line := range lines {
		for x, num := range strings.FieldsFunc(line, splitFn) {
			t, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			board[x][y] = uint64(t)
		}
	}

	return recurP345(0, 0, board)
}

func removeLineColumn(index int, board [][]uint64) [][]uint64 {
	tBoard := make([][]uint64, len(board)-1)
	for j, _ := range board {
		if j == 0 {
			continue
		}
		tBoard[j-1] = make([]uint64, len(board)-1)
		gap := 0
		for k, value := range board[j] {
			if k == index {
				gap = 1
				continue
			}
			tBoard[j-1][k-gap] = value
		}
	}
	return tBoard
}

func recurP345(max, currentValue uint64, board [][]uint64) uint64 {
	left := len(board)
	if left == 0 {
		return currentValue
	}
	if max > currentValue+uint64(left*1000) {
		return currentValue
	}

	for i, _ := range board {
		tBoard := removeLineColumn(i, board)
		localMax := recurP345(max, currentValue+board[0][i], tBoard)
		if localMax > max {
			max = localMax
		}
	}
	return max
}
