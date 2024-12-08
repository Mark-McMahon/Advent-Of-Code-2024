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

	// delimit by new line
	lines := strings.Split(input, "\n")

	ordering := [][]int{}
	orderingSublist := []int{}

	updates := [][]int{}
	updatesSublist := []int{}

	updatesBool := false
	for _, line := range lines {
		if line == "" {
			updatesBool = true
			continue
		}
		if !updatesBool {
			for i := range strings.Split(line, "|") {
				convertedNum, _ := strconv.Atoi(strings.Split(line, "|")[i])
				orderingSublist = append(orderingSublist, convertedNum)
			}
			ordering = append(ordering, orderingSublist)
			orderingSublist = []int{}

		} else {
			for i := range strings.Split(line, ",") {
				convertedNum, _ := strconv.Atoi(strings.Split(line, ",")[i])
				updatesSublist = append(updatesSublist, convertedNum)
			}
			updates = append(updates, updatesSublist)
			updatesSublist = []int{}
		}
	}

	graph := make(map[int][]int)
	for i := range len(ordering) {
		graph[ordering[i][0]] = append(graph[ordering[i][0]], ordering[i][1])
	}

	validUpdates, invalidUpdates := isUpdateValid(graph, updates)

	total := 0

	for i := range validUpdates {
		middleIndex := (len(validUpdates[i]) / 2)
		total += validUpdates[i][middleIndex]
	}

	invalidSortedUpdates := [][]int{}
	for i := range invalidUpdates {
		subgraph := buildSubGraph(graph, invalidUpdates[i])
		sortedUpdate := topologicalSort(subgraph, invalidUpdates[i])
		invalidSortedUpdates = append(invalidSortedUpdates, sortedUpdate)
	}

	invalidTotal := 0
	for i := range invalidSortedUpdates {
		middleIndex := (len(invalidSortedUpdates[i]) / 2)
		invalidTotal += invalidSortedUpdates[i][middleIndex]
	}

	fmt.Println(invalidTotal)

}

func topologicalSort(subgraph map[int][]int, nodes []int) []int {
	indegree := map[int]int{}
	for _, node := range nodes {
		indegree[node] = 0
	}

	for _, outgoings := range subgraph {
		for _, outgoing := range outgoings {
			indegree[outgoing]++
		}
	}

	q := []int{}
	for node := range indegree {
		if indegree[node] == 0 {
			q = append(q, node)
		}
	}

	sorted := []int{}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		sorted = append(sorted, curr)

		for _, outgoing := range subgraph[curr] {
			indegree[outgoing]--
			if indegree[outgoing] == 0 {
				q = append(q, outgoing)
			}
		}
	}

	return sorted

}

func buildSubGraph(graph map[int][]int, invalidUpdates []int) map[int][]int {

	subGraph := map[int][]int{}
	for _, node := range invalidUpdates {
		for _, neighbor := range graph[node] {
			if contains(invalidUpdates, neighbor) {
				subGraph[node] = append(subGraph[node], neighbor)
			}
		}
	}
	return subGraph
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func isUpdateValid(graph map[int][]int, updates [][]int) ([][]int, [][]int) {
	validUpdates := [][]int{}
	invalidUpdates := [][]int{}
	for i := range updates {
		isValid := true
		for j := 0; j < len(updates[i])-1; j++ {
			if !isCurrUpdateValid(graph, updates[i][j], updates[i][j+1]) {
				isValid = false
				break
			}
		}
		if isValid {
			validUpdates = append(validUpdates, updates[i])
		} else {
			invalidUpdates = append(invalidUpdates, updates[i])
		}
	}
	return validUpdates, invalidUpdates
}

func isCurrUpdateValid(graph map[int][]int, prereq int, req int) bool {
	for i := range graph[prereq] {
		if graph[prereq][i] == req {
			return true
		}
	}
	return false
}
