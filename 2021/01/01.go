package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.StringsToInts(utils.ReadFile(2021, 1, "\n"))
	fmt.Println("2021 Day 01")
	fmt.Println("\tPart 1:", findIncreases(input))
	fmt.Println("\tPart 2:", findIncreases2(input))

}

func findIncreases(input []int) int {
	increases := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases += 1
		}
	}
	return increases
}

func findIncreases2(input []int) int {
	increases := 0
	for i := 3; i < len(input); i++ {
		if input[i]+input[i-1]+input[i-2] > input[i-1]+input[i-2]+input[i-3] {
			increases += 1
		}
	}
	return increases
}
