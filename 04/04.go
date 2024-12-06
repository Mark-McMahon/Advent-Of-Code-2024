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

	result1 := part1(matrix)
	fmt.Println(result1)

	result2 := part2(matrix)
	fmt.Println(result2)

}

func part2(matrix [][]rune) int {

	ROWS := len(matrix)
	COLS := len(matrix[0])

	totalCount := 0
	for r := 0; r < ROWS-1; r++ { // we can avoid edges here :)
		for c := 0; c < COLS-1; c++ {
			if matrix[r][c] == 'A' {
				if dfspart2(r, c, matrix) {
					totalCount++
				}
			}
		}
	}

	return totalCount

}

func dfspart2(r int, c int, matrix [][]rune) bool {

	if r-1 < 0 || r+1 >= len(matrix) || c-1 < 0 || c+1 >= len(matrix[0]) {
		return false
	}

	topLeft := matrix[r-1][c-1]
	bottomRight := matrix[r+1][c+1]
	topRight := matrix[r-1][c+1]
	bottomLeft := matrix[r+1][c-1]

	diag1Valid := (topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')
	diag2Valid := (topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')

	return diag1Valid && diag2Valid
}

func part1(matrix [][]rune) int {

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
				if dfspart1(r, c, word, matrix, d) {
					totalCount++
				}
			}

		}
	}

	return totalCount
}

func dfspart1(r int, c int, word string, matrix [][]rune, d [2]int) bool {
	for i := 0; i < len(word); i++ {
		dr := r + d[0]*i
		dc := c + d[1]*i

		if dr < 0 || dc < 0 || dr >= len(matrix) || dc >= len(matrix[0]) || matrix[dr][dc] != rune(word[i]) {
			return false
		}
	}
	return true
}
