package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2022, 1, "\n\n")
	fmt.Println("2022 Day 01")
	fmt.Println("\tPart 1:", sumMax(input, 1))
	fmt.Println("\tPart 2:", sumMax(input, 3))
}

// sums the [topAmount] biggest values of the input slice
func sumMax(input []string, topAmount int) int {
	maximumsSlice := make([]int, topAmount)

	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		stringSlice := strings.Split(input[lineIndex], "\n")
		intSlice := utils.StringsToInts(stringSlice)
		value := utils.SumIntSlice(intSlice)
		utils.InsertIfLarger(maximumsSlice, value)
	}

	return utils.SumIntSlice(maximumsSlice)
}
