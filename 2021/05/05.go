package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2021, 5, "\n")

	fmt.Println("2021 Day 05")
	fmt.Println("\tPart 1:", findOverlaps(input))
	fmt.Println("\tPart 2:", findOverlapsWithDiagonals(input))
}

func findOverlaps(input []string) int {
	counter := make(map[string]int)
	total := 0

	for _, line := range input {
		points := strings.Split(line, " -> ")
		coordinate1 := utils.StringsToInts(strings.Split(points[0], ","))
		coordinate2 := utils.StringsToInts(strings.Split(points[1], ","))
		var minX, maxX, minY, maxY int

		if coordinate1[0] == coordinate2[0] {
			if coordinate1[1] < coordinate2[1] {
				minY = coordinate1[1]
				maxY = coordinate2[1]
			} else {
				minY = coordinate2[1]
				maxY = coordinate1[1]
			}

			for y := minY; y <= maxY; y++ {
				coordinate := fmt.Sprintf("%d, %d", coordinate1[0], y)
				if _, ok := counter[coordinate]; ok {
					counter[coordinate] += 1
				} else {
					counter[coordinate] = 1
				}
			}
		} else if coordinate1[1] == coordinate2[1] {
			if coordinate1[0] < coordinate2[0] {
				minX = coordinate1[0]
				maxX = coordinate2[0]
			} else {
				minX = coordinate2[0]
				maxX = coordinate1[0]

			}

			for x := minX; x <= maxX; x++ {
				coordinate := fmt.Sprintf("%d, %d", x, coordinate1[1])
				if _, ok := counter[coordinate]; ok {
					counter[coordinate] += 1
				} else {
					counter[coordinate] = 1
				}
			}
		}

	}

	for _, val := range counter {
		if val > 1 {
			total += 1
		}
	}

	return total
}

func findOverlapsWithDiagonals(input []string) int {
	counter := make(map[string]int)
	total := 0

	for _, line := range input {
		points := strings.Split(line, " -> ")
		coordinate1 := utils.StringsToInts(strings.Split(points[0], ","))
		coordinate2 := utils.StringsToInts(strings.Split(points[1], ","))
		var vectorX, vectorY int

		if vectorX = (coordinate2[0] - coordinate1[0]); vectorX > 0 {
			vectorX = 1
		} else if vectorX < 0 {
			vectorX = -1
		}

		if vectorY = (coordinate2[1] - coordinate1[1]); vectorY > 0 {
			vectorY = 1
		} else if vectorY < 0 {
			vectorY = -1
		}

		for {
			coordinate := fmt.Sprintf("%d, %d", coordinate1[0], coordinate1[1])
			if _, ok := counter[coordinate]; ok {
				counter[coordinate] += 1
			} else {
				counter[coordinate] = 1
			}

			if coordinate2[0] == coordinate1[0] && coordinate2[1] == coordinate1[1] {
				break
			}

			coordinate1[0] += vectorX
			coordinate1[1] += vectorY
		}
	}

	for _, val := range counter {
		if val > 1 {
			total += 1
		}
	}

	return total
}
