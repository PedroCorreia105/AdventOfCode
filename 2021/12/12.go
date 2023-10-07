package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := utils.ReadFile(2021, 12, "\n")

	fmt.Println("2021 Day 12")
	fmt.Println("\tPart 1:", getPaths(input, 1))
	fmt.Println("\tPart 2:", getPaths(input, 2))
}

func getPaths(input []string, version int) int {
	graph := make(map[string][]string)
	paths := [][]string{{"start"}}

	for _, line := range input {
		parts := strings.Split(line, "-")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
	}

	pathIndex := 0
	for pathIndex < len(paths) {
		path := paths[pathIndex]
		lastNode := path[len(path)-1]
		if lastNode != "end" {
			paths[pathIndex] = paths[len(paths)-1]
			paths = paths[:len(paths)-1]
			nextNodes := graph[lastNode]
			for _, nextNode := range nextNodes {
				pathCopy := make([]string, len(path))
				copy(pathCopy, path)
				if nextNode != "start" && (nextNode == "end" || unicode.IsUpper(rune(nextNode[0])) || canVisitLowerCase(&pathCopy, nextNode, version)) {
					paths = append(paths, append(pathCopy, nextNode))
				}
			}
		} else {
			pathIndex++
		}
	}

	return pathIndex
}

func canVisitLowerCase(array *[]string, element string, version int) bool {
	total := 0
	for _, elem := range *array {
		if elem == element {
			total += 1
		}
	}
	if version == 1 {
		return total == 0
	} else if total == 0 {
		return true
		// use the first position to "encode" when a lower case node was visited
	} else if total == 1 && (*array)[0] == "start" {
		(*array)[0] = "Start"
		return true
	}
	return false
}
