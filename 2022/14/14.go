package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

const ROCK = 0
const SAND = 1

func main() {
	input := utils.ReadFile(2022, 14, "\n")
	fmt.Println("2022 Day 14")
	fmt.Println("\tPart 1:", calculate(input, 1))
	fmt.Println("\tPart 2:", calculate(input, 2))
}

func calculate(input []string, part int) (int) {
	total := 0
	sandLocation := []int{500, 0}
	graph, maxY := processInput(input)

	for (part == 1 && sandLocation[1] < maxY || part == 2 && utils.GetFromMap(graph, 500, 0) != SAND) {
		nextLocation := nextLocationAvailable(sandLocation, graph)
		if (nextLocation != nil && nextLocation[1] != maxY + 2) {
			sandLocation = nextLocation
		} else {
			utils.AddToMap(graph, sandLocation[0], sandLocation[1], SAND)
			sandLocation = []int{500, 0}
		}
	}

	for _, value := range graph {
		if (value == SAND) {
			total++
		}
	}

	return total
}

func processInput(input []string) (map[string]int, int) {
	maxY := 0
	graph := make(map[string]int)

	for _, line := range(input) {
		points := strings.Split(line, " -> ")
		// dont run for the first coordinate
		previous := utils.StringsToInts(strings.Split(points[0], ","))
		for _, pair := range(points[1:]) {
			pairNumbers := utils.StringsToInts(strings.Split(pair, ","))
			maxY = utils.Max(maxY, pairNumbers[1])
			if (len(previous) != 0) {
				// traveling vertically
				if (previous[0] == pairNumbers[0]) {
					a, b := utils.SortNumbers(previous[1], pairNumbers[1])
					for i := a; i <= b; i++ {
						utils.AddToMap(graph, previous[0], i, ROCK)
					}
				// traveling horizontally
				} else {
					a, b := utils.SortNumbers(previous[0], pairNumbers[0])
					for i := a; i <= b; i++ {
						utils.AddToMap(graph, i, previous[1], ROCK)
					}
				}
			}
			previous = utils.CopyIntSlice(pairNumbers)
		}
	}
	return graph, maxY
}

func nextLocationAvailable(coordinate []int, graph map[string]int) ([]int) {
	var vectors = [][]int{{0, 1}, {-1, 1}, {1, 1}}

	for _, vector := range vectors {
		x := coordinate[0] + vector[0]
		y := coordinate[1] + vector[1]
		if !utils.IsInMap(graph, x, y) {
			return []int{x, y}
		}
	}
	return nil
}
