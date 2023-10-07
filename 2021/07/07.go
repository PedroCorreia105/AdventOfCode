package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.StringsToInts(utils.ReadFile(2021, 7, ","))

	fmt.Println("2021 Day 07")
	fmt.Println("\tPart 1:", calculateFuel(input, 1))
	fmt.Println("\tPart 2:", calculateFuel(input, 2))
}

func calculateFuel(input []int, version int) int {
	counter := make(map[int]int)
	min, max, minGas := -1, -1, -1

	for _, val := range input {
		if _, ok := counter[val]; ok {
			counter[val] += 1
		} else {
			counter[val] = 1
		}
		if min == -1 || min > val {
			min = val
		}

		if max == -1 || max < val {
			max = val
		}
	}

	for destination := min; destination < max; destination++ {
		total := 0

		for origin, ammount := range counter {
			if version == 1 {
				total += ammount * utils.AbsDiff(destination, origin)
			} else {
				total += ammount * sumTo(utils.AbsDiff(destination, origin))
			}
		}

		if minGas == -1 || minGas > total {
			minGas = total
		}
	}

	return minGas
}

func sumTo(y int) int {
	return int(float64(y+1) * (float64(y) / 2))
}
