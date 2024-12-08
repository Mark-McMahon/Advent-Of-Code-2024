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

	res := part1(matrix)

	fmt.Println(res)

}

func part1(matrix [][]rune) int {

	x, y := findGaurd(matrix)

	directions := [][2]int{
		{-1, 0}, // up
		{0, 1},  //right
		{1, 0},  // down
		{0, -1}, //left
	}

	cycles := 0

	for i := range matrix {
		for j := range matrix[i] {
			if (i == x && j == y) || matrix[i][j] == '#' {
				continue
			}
			matrix[i][j] = '#'
			visited := make(map[[3]int]bool)
			if dfs(matrix, x, y, 0, visited, directions) {
				cycles++
			}
			matrix[i][j] = '.'
		}
	}

	return cycles

}

func dfs(matrix [][]rune, r, c, dirIndex int, visited map[[3]int]bool, d [][2]int) bool {
	if r < 0 || c < 0 || r >= len(matrix) || c >= len(matrix[0]) {
		return false
	}
	if visited[[3]int{r, c, dirIndex}] {
		return true
	}

	visited[[3]int{r, c, dirIndex}] = true

	dr := r + d[dirIndex][0]
	dc := c + d[dirIndex][1]

	if dr >= 0 && dr < len(matrix) && dc >= 0 && dc < len(matrix[0]) && matrix[dr][dc] == '#' {
		dirIndex = (dirIndex + 1) % 4
		return dfs(matrix, r, c, dirIndex, visited, d)
	} else {
		return dfs(matrix, dr, dc, dirIndex, visited, d)
	}

}

func findGaurd(matrix [][]rune) (int, int) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}
