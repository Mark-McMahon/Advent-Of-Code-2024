package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	safe := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var record []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				continue
			}
			record = append(record, num)
		}

		if isMonotonic(record) && isSafeDiff(record) {
			safe += 1
			continue
		}

		for i := 0; i < len(record); i++ {
			removal := append([]int{}, record[:i]...)
			if i != len(record)-1 {
				removal = append(removal, record[i+1:]...)
			}
			if isMonotonic(removal) && isSafeDiff(removal) {
				safe++
				break
			}

		}

	}

	fmt.Println(safe)

}

func isMonotonic(record []int) bool {

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(record); i++ {
		if record[i] < record[i-1] {
			isIncreasing = false
		}
		if record[i] > record[i-1] {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing

}

func isSafeDiff(record []int) bool {

	isSafeDiff := true

	for i := 1; i < len(record); i++ {

		diff := record[i] - record[i-1]
		absdiff := int(math.Abs(float64(diff)))
		if absdiff < 1 || absdiff > 3 {
			isSafeDiff = false
		}
	}
	return isSafeDiff
}
