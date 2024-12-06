package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fi, _ := os.Open("input.txt")
	defer fi.Close()

	data, _ := io.ReadAll(fi)
	input := string(data)

	// delimit by new line
	lines := strings.Split(input, "\n")

	// create matrix out of chars
	var matrix = make([][]rune, len(lines))
	for i := range matrix {
		matrix[i] = make([]rune, len(lines[0]))
	}
	for i, line := range lines {
		for j, char := range line {
			matrix[i][j] = char
		}
	}

	ROWS := len(matrix)
	COLS := len(matrix[0])

	word := "XMAS"

	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
		{-1, 1},
		{-1, -1},
		{1, 1},
		{1, -1},
	}

	totalCount := 0
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			for _, d := range directions {
				if dfs(r, c, word, matrix, d) {
					totalCount++
				}
			}

		}
	}

	fmt.Println(totalCount)
}

func dfs(r int, c int, word string, matrix [][]rune, d [2]int) bool {
	for i := 0; i < len(word); i++ {
		dr := r + d[0]*i
		dc := c + d[1]*i

		if dr < 0 || dc < 0 || dr >= len(matrix) || dc >= len(matrix[0]) || matrix[dr][dc] != rune(word[i]) {
			return false
		}
	}
	return true
}
