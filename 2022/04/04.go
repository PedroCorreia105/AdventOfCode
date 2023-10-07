package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2022, 4, "\n")
	fmt.Println("2022 Day 04")
	fmt.Println("\tPart 1:", calculate(input, 1))
	fmt.Println("\tPart 2:", calculate(input, 2))
}

func calculate(input []string, part int) (int) {
	total := 0
	
	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		stringSlice := strings.Split(input[lineIndex], ",")
		elf1 := utils.StringsToInts(strings.Split(stringSlice[0], "-"))
		elf2 := utils.StringsToInts(strings.Split(stringSlice[1], "-"))
		
		if (part == 1 && (elf1[0] <= elf2[0] && elf1[1] >= elf2[1] || elf2[0] <= elf1[0] && elf2[1] >= elf1[1])) {
			total++
		} else if (part == 2 && (elf1[0] <= elf2[0] && elf1[1] >= elf2[0] || elf2[0] <= elf1[0] && elf2[1] >= elf1[0])) {
			total++
		}
	}
	
	return total
}