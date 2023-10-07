package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)


func main() {
	input := utils.ReadFile(2022, 10, "\n")
	fmt.Println("2022 Day 10")
	// fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:\n", part2(input))
}

func part1(input []string) (int) {
	x, cycle, strength :=  1, 0, 0
	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		if (input[lineIndex] == "noop") {
			cycle++
			strength += calculateStrength(cycle, x)
		} else {
			line := strings.Split(input[lineIndex], " ")
			cycle++
			strength += calculateStrength(cycle, x)
			cycle++
			strength += calculateStrength(cycle, x)
			x += utils.StringToInt(line[1])
		}
	}
	
	return strength
}

func calculateStrength(cycle, x int) (int) {
	if (cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220) {
		return cycle * x
	}
	return 0
}

// 13060

func part2(input []string) (string) {
	x, cycle :=  1, 0
	output, c := "", ""
	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		if (input[lineIndex] == "noop") {
			cycle, c = calculateStrength2(cycle, x)
			output += c
		} else {
			line := strings.Split(input[lineIndex], " ")
			cycle, c = calculateStrength2(cycle, x)
			output += c
			cycle, c = calculateStrength2(cycle, x)
			output += c
			x += utils.StringToInt(line[1])
		}
		if (len(output)%40 == 0) { output+= "\n"}
	}
	
	return output
}

func calculateStrength2(cycle, x int) (int, string) {
	if (cycle%40 == x - 1 || cycle%40 == x || cycle%40 == x + 1) {
		return cycle+1, "#"
	}
	return cycle+1, "."
}

// 13060