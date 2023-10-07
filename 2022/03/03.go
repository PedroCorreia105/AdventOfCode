package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile(2022, 3, "\n")
	fmt.Println("2022 Day 03")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) int {
	points := 0

	for lineIndex := 0; lineIndex < len(input); lineIndex++ {
		char := findCommonChar(input[lineIndex])
		if char >= 'a' {
			char -= 96
		} else {
			char -= 38
		}
		points += int(char)
	}

	return points
}

func part2(input []string) int {
	points := 0

	for lineIndex := 0; lineIndex < len(input); lineIndex += 3 {
		char := findCommonChar2(input[lineIndex], input[lineIndex+1], input[lineIndex+2])
		if char >= 'a' {
			char -= 96
		} else {
			char -= 38
		}
		points += int(char)
	}

	return points
}

func findCommonChar(word string) byte {
	for c1 := 0; c1 < len(word)/2; c1++ {
		for c2 := len(word) / 2; c2 < len(word); c2++ {
			if word[c1] == word[c2] {
				return word[c1]
			}
		}
	}
	return 'v'
}

func findCommonChar2(word1 string, word2 string, word3 string) byte {
	for c1 := 0; c1 < len(word1); c1++ {
		for c2 := 0; c2 < len(word2); c2++ {
			if word1[c1] == word2[c2] {
				for c3 := 0; c3 < len(word3); c3++ {
					if word2[c2] == word3[c3] {
						return word1[c1]
					}
				}
			}
		}
	}
	return 'v'
}
