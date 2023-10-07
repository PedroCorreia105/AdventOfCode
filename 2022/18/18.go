package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

var vectors = [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

type Coordinate struct {
	x       int
	y       int
	z       int
	visited bool
	blocked bool
}

func main() {
	input := utils.ReadFile(2022, 18, "\n")
	fmt.Println("2022 Day 18")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) (int) {
	var x, y, z int
	coordinates := make(map[string]Coordinate)
	adjacentCubes := 0
	
	for _, line := range(input) {
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		coordinates[line] = Coordinate{x: x, y: y, z: z}
	}

	for _, coordinate := range(coordinates) {
		for _, vector := range(vectors) {
			adjacentCoordinate := fmt.Sprintf("%d,%d,%d", coordinate.x + vector[0], coordinate.y + vector[1], coordinate.z + vector[2])
			if _, ok := coordinates[adjacentCoordinate]; ok {
				adjacentCubes++
			}
		}
	}
	return len(coordinates) * 6 - adjacentCubes
}

func part2(input []string) (int) {
	coordinates := make(map[string]Coordinate)
	queue := []string{"0,0,0"}
	total, min, max := 0, -1, 25
	
	for x := min; x <= max; x++ {
		for y := min; y <= max; y++ {
			for z := min; z <= max; z++ {
				coordinates[fmt.Sprintf("%d,%d,%d", x, y, z)] = Coordinate{x: x, y: y, z: z, visited: false, blocked: false}
			}
		}
	}

	for _, line := range(input) {
		// goland does not allow changing field of mapped object, so ...
		if entry, ok := coordinates[line]; ok {
			entry.blocked = true
			coordinates[line] = entry
		}
	}

	// bfs that when enconters a blocked cell, increases the total
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		if (!coordinates[current].visited) {
			if entry, ok := coordinates[current]; ok {
				entry.visited = true
				coordinates[current] = entry
			}

			for _, vector := range vectors {
				x := coordinates[current].x + vector[0]
				y := coordinates[current].y + vector[1]
				z := coordinates[current].z + vector[2]
				if (min <= x && x <= max && min <= y && y <= max && min <= z && z <= max) {
					nextNode := fmt.Sprintf("%d,%d,%d", x, y, z)
					if (coordinates[nextNode].blocked) {
						total++
					} else if (!coordinates[nextNode].visited) {
						queue = append(queue, nextNode)
					}
				}
			}
		}
	}
	return total
}