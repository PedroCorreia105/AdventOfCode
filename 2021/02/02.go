package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadFile(2021, 2, "\n")

	fmt.Println("2021 Day 02")
	fmt.Println("\tPart 1:", findPosition(input))
	fmt.Println("\tPart 2:", findPosition2(input))
}

func findPosition(input []string) int {
	depth, hPosition := 0, 0

	for _, line := range input {
		commands := strings.Split(line, " ")
		number, err := strconv.Atoi(commands[1])

		if err != nil {
			panic(err)
		}

		switch commands[0] {
		case "forward":
			hPosition += number
		case "up":
			depth -= number
		case "down":
			depth += number
		}
	}

	return depth * hPosition
}

func findPosition2(input []string) int {
	depth, hPosition, aim := 0, 0, 0

	for _, line := range input {
		commands := strings.Split(line, " ")
		number, err := strconv.Atoi(commands[1])

		if err != nil {
			panic(err)
		}

		switch commands[0] {
		case "forward":
			hPosition += number
			depth += aim * number
		case "up":
			aim -= number
		case "down":
			aim += number
		}
	}

	return depth * hPosition
}
