package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile(2021, 3, "\n")

	fmt.Println("2021 Day 03")
	fmt.Println("\tPart 1:", findGammaEpsilon(input))
	fmt.Println("\tPart 2:", findOxygenAndCO2(input))
}

func findGammaEpsilon(input []string) int64 {
	positions := make([]int, len(input[0]))
	gamma, epsilon := "", ""
	for _, line := range input {
		for charIndex, char := range line {
			// if bit is zero
			if char == 48 {
				positions[charIndex] += 1
			}
		}
	}

	for _, freq := range positions {
		// if zeros are more frequent
		if freq >= len(input)/2 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaValue, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonValue, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	return gammaValue * epsilonValue
}

func findOxygenAndCO2(input []string) int64 {
	oxygen, co2 := "", ""

	for len(oxygen) < len(input[0]) {
		oxygen = getNextPattern(oxygen, input, "0", "1")
	}

	for len(co2) < len(input[0]) {
		co2 = getNextPattern(co2, input, "1", "0")
	}

	oxygenValue, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		panic(err)
	}

	co2Value, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		panic(err)
	}

	return oxygenValue * co2Value
}

func getNextPattern(pattern string, input []string, moreZeros string, moreOnes string) string {
	lastLine := ""
	freq, total := 0, 0
	for _, line := range input {
		if strings.HasPrefix(line, pattern) {
			total += 1
			lastLine = line

			// if bit is zero
			if line[len(pattern)] == 48 {
				freq += 1
			}
		}
	}

	if total == 1 {
		return lastLine
	}

	if freq > total/2 {
		return pattern + moreZeros
	}

	return pattern + moreOnes
}
