package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile(2022, 2, "\n")
	fmt.Println("2022 Day 02")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) (int) {
	points := 0
	
	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		switch input[lineIndex] {
		case "A Y", "B Z", "C X": // configurations where we win
			points += 6
		case "A X", "B Y", "C Z": // configurations where we tie
			points += 3
		}
		
		switch input[lineIndex][2] { // our move
		case 'X': // rock
			points += 1
		case 'Y': // paper
			points += 2
		case 'Z': // scissors
			points += 3
		}
	}
	
	return points
}

func part2(input []string) (int) {
	points := 0
	
	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		switch input[lineIndex] {
		case "A Y", "B X", "C Z": // configurations where we play rock
			points += 1
		case "A Z", "B Y", "C X": // configurations where we play paper
			points += 2
		default:
			points += 3
		}
		
		switch input[lineIndex][2] { // desired outcome of the match
		case 'Y': // tie
			points += 3
		case 'Z': // win
			points += 6
		}
	}
	
	return points
}