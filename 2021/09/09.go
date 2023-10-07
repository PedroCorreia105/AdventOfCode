package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"sort"
)

var maxX int
var maxY int
var visited = []string{}
var input = []string{}

func main() {
	input = utils.ReadFile(2021, 9, "\n")
	maxY = len(input)
	maxX = len(input[0])

	fmt.Println("2021 Day 09")
	fmt.Println("\tPart 1:", getRiskLevel())
	fmt.Println("\tPart 2:", getBasins())
}

func getRiskLevel() int {
	total := 0

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			biggerThanUp := y == 0 || input[y-1][x] > input[y][x]
			biggerThanDown := y == maxY-1 || input[y+1][x] > input[y][x]
			biggerThanLeft := x == 0 || input[y][x-1] > input[y][x]
			biggerThanRight := x == maxX-1 || input[y][x+1] > input[y][x]

			if biggerThanUp && biggerThanDown && biggerThanLeft && biggerThanRight {
				total += int(input[y][x]-'0') + 1
			}
		}
	}

	return total
}

func getBasins() int {
	basins := []int{}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			basins = append(basins, visit(x, y))
		}
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func visit(x, y int) int {
	if 0 <= y && y <= maxY-1 && 0 <= x && x <= maxX-1 && input[y][x]-'0' != 9 {
		if !utils.ContainsString(visited, fmt.Sprintf("%d, %d", x, y)) {
			visited = append(visited, fmt.Sprintf("%d, %d", x, y))

			return visit(x+1, y) + visit(x-1, y) + visit(x, y+1) + visit(x, y-1) + 1
		}
	}
	return 0
}
