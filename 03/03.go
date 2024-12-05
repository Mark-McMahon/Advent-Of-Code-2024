package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	data, err := io.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	input := string(data)

	sum := 0
	i := 0

	for i < len(input) {
		if input[i] == 'm' && i+5 < len(input) && input[i:i+4] == "mul(" {
			i += 4
			startX := i

			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				i++
			}
			if input[i] == input[startX] || input[i] != ',' {
				continue
			}

			x := input[startX:i]

			i++ // skip comma

			startY := i

			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				i++
			}
			if input[i] == input[startY] || input[i] != ')' {
				continue
			}

			y := input[startY:i]

			intX, _ := strconv.Atoi(x)
			intY, _ := strconv.Atoi(y)

			sum += intX * intY

		}
		i++
	}
	fmt.Println(sum)
}
