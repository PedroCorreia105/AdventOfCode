package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2022, 5, "\n\n")
	crates := strings.Split(input[0], "\n")
	instructions := strings.Split(input[1], "\n")
	fmt.Println("2022 Day 05")
	fmt.Println("\tPart 1:", calculate(crates, instructions, 1))
	fmt.Println("\tPart 2:", calculate(crates, instructions, 2))
}

func calculate(crates []string, instructions []string, part int) string {
	total := ""
	crateMap := make([][]string, len(crates[0])/4+1)
	var c1, c2, c3 int

	for lineIndex := 0; lineIndex < len(crates); lineIndex++ {
		for columnIndex := 0; columnIndex < len(crates[lineIndex]); columnIndex++ {
			if crates[lineIndex][columnIndex] >= 'A' && crates[lineIndex][columnIndex] <= 'Z' {
				crateMap[columnIndex/4] = append(crateMap[columnIndex/4], string(crates[lineIndex][columnIndex]))
			}
		}
	}

	for lineIndex := 0; lineIndex < len(instructions); lineIndex++ {
		fmt.Sscanf(instructions[lineIndex], "move %d from %d to %d", &c1, &c2, &c3)
		if part == 1 {
			for box := 0; box < c1; box++ {
				crateMap[c3-1] = append([]string{crateMap[c2-1][0]}, crateMap[c3-1]...)
				crateMap[c2-1] = crateMap[c2-1][1:]
			}
		} else {
			crateMap[c3-1] = append(utils.CopyStringSlice(crateMap[c2-1][:c1]), crateMap[c3-1]...)
			crateMap[c2-1] = crateMap[c2-1][c1:]
		}
	}
	for _, i := range crateMap {
		if len(i) > 0 {
			total += i[0]
		}
	}
	return total
}
