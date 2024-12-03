package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		num1, err1 := strconv.Atoi(fields[0])
		num2, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			continue
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	total := 0
	simScore := 0

	list2map := make(map[int]int)

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		absDiff := int(math.Abs(float64(diff)))
		total += absDiff

		list2map[list2[i]] = list2map[list2[i]] + 1
	}

	for i := 0; i < len(list1); i++ {
		simScore += list1[i] * list2map[list1[i]]
	}

	fmt.Println(total)
	fmt.Println(simScore)
}
