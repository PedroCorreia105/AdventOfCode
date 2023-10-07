package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

var NUMBER_OF_DIGITS = 14

func main() {
	input := utils.ReadFile(2021, 24, "\n")

	fmt.Println("2021 Day 24")
	fmt.Println("\tPart 1:", calculate(input, 9_999_999, 1_111_111, -1))
	fmt.Println("\tPart 2:", calculate(input, 1_111_111, 9_999_999, 1))
}

func calculate(input []string, min, max, iterator int) int {
	processedInput := processInput(input)
	for number := min; number != max; number += iterator {
		numberSlice := utils.IntToSlice(number)
		if !utils.ContainsInt(numberSlice, 0) {
			result := testDigits(processedInput, numberSlice)
			if len(result) > 0 {
				return utils.SliceToInt(result)
			}
		}
	}
	return 0
}

func processInput(input []string) []int {
	var values = make([]int, NUMBER_OF_DIGITS)
	// there is 1 block of instructions for each digit of the model number
	for block := 0; block < NUMBER_OF_DIGITS; block += 1 {
		// each block has 18 instructions
		if strings.Contains(input[block*18+5], "-") {
			// if the 5th instruction has a negative number register its value
			values[block] = utils.StringToInt(strings.Split(input[block*18+5], " ")[2])
		} else {
			// else register the value in the 15th line
			values[block] = utils.StringToInt(strings.Split(input[block*18+15], " ")[2])
		}
	}
	return values
}

func testDigits(processedInput, possibleNumber []int) []int {
	var result = make([]int, NUMBER_OF_DIGITS)
	zValue := 0
	digitIndex := 0

	for digit := 0; digit < NUMBER_OF_DIGITS; digit += 1 {
		value := processedInput[digit]

		if value < 0 {
			result[digit] = ((zValue % 26) + value)
			zValue /= 26
			if result[digit] < 1 || 9 < result[digit] {
				return []int{}
			}
		} else {
			zValue = zValue*26 + possibleNumber[digitIndex] + value
			result[digit] = possibleNumber[digitIndex]
			digitIndex += 1
		}
	}

	return result
}
