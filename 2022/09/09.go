package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile(2022, 9, "\n")
	fmt.Println("2022 Day 09")
	fmt.Println("\tPart 1:", calculate(input, 2))
	fmt.Println("\tPart 2:", calculate(input, 10))
}

func calculate(input []string, knotNumber int) (int) {
	var direction string
	var displacement int
	var knotsLocation [10][2]int
	var knotsVector [10][2]int
	visitedPositions := make(map[string]int)
	mapa := map[byte][2]int{
		'U': [2]int{0, -1},
		'D': [2]int{0, 1},
		'L': [2]int{-1, 0},
		'R': [2]int{1, 0},
	}

	// process each line
	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		fmt.Sscanf(input[lineIndex], "%s %d", &direction, &displacement)
		for step := 0; step < displacement; step++ {
			// use direction provided in inpput as initial vector
			knotsLocation[0][0] += mapa[direction[0]][0]
			knotsLocation[0][1] += mapa[direction[0]][1]
			for knot := 0; knot < knotNumber-1; knot++ {
				// calculate next knot vector based on knot position
				knotsVector[knot+1][0] = knotsLocation[knot][0] - knotsLocation[knot+1][0] 
				knotsVector[knot+1][1] = knotsLocation[knot][1] - knotsLocation[knot+1][1]
				// if one of the vector components is at least 2, move the next knot
				if (utils.Abs(knotsVector[knot+1][0]) > 1 || utils.Abs(knotsVector[knot+1][1]) > 1) {
					knotsLocation[knot+1][0] += unitify(knotsVector[knot+1][0])
					knotsLocation[knot+1][1] += unitify(knotsVector[knot+1][1])
				}
			}
			utils.AddToMap(visitedPositions, knotsLocation[knotNumber-1][0], knotsLocation[knotNumber-1][1], 1)
		}
	}

	return len(visitedPositions)
}

// vector component with a 2/-2 turn into a 1/-1
func unitify(number int) (int) {
	if (number%2==0) { return number/2}
	return number
}