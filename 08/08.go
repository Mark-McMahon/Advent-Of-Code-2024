package main

import (
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

}
