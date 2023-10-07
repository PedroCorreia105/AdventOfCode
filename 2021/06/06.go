package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.StringsToInts(utils.ReadFile(2021, 6, ","))

	fmt.Println("2021 Day 06")
	fmt.Println("\tPart 1:", countFish(input, 80))
	fmt.Println("\tPart 2:", countFish(input, 256))
}

func countFish(input []int, days int) int {
	counter := make([]int, 9)
	total := 0

	for _, val := range input {
		counter[val] += 1
	}

	for day := 1; day <= days; day++ {
		breedingToday := counter[0]
		counter = append(counter[1:], breedingToday)
		counter[6] += breedingToday
	}

	for _, val := range counter {
		total += val
	}

	return total
}
