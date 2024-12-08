package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fi, _ := os.Open("input.txt")
	defer fi.Close()

	data, _ := io.ReadAll(fi)
	input := string(data)

	lines := strings.Split(input, "\n")

	total := 0
	for i := range lines {
		parts := strings.Split(lines[i], ":")

		fields := strings.Fields(parts[1])
		target, _ := strconv.Atoi(parts[0])

		nums := make([]int, 0, len(fields)+1)

		for _, strVals := range fields {
			num, _ := strconv.Atoi(strVals)
			nums = append(nums, num)
		}

		if isPossible(target, nums[0], 1, nums) {
			total += target
		}

	}
	fmt.Println(total)

}

func isPossible(target, curVal, i int, nums []int) bool {
	if i == len(nums) {
		if curVal == target {
			return true
		} else {
			return false
		}
	}

	return isPossible(target, (curVal*nums[i]), i+1, nums) ||
		isPossible(target, (curVal+nums[i]), i+1, nums) ||
		isPossible(target, concatenate(curVal, nums[i]), i+1, nums)

}

func concatenate(x, y int) int {
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)

	concat := strX + strY

	intConcat, _ := strconv.Atoi(concat)

	return intConcat

}
