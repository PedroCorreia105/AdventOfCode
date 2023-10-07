package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"sort"
)

var mapa = map[rune]byte{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func main() {
	input := utils.ReadFile(2021, 10, "\n")

	fmt.Println("2021 Day 10")
	fmt.Println("\tPart 1:", getSyntaxScore(input))
	fmt.Println("\tPart 2:", getMiddleScore(input))
}

func getSyntaxScore(input []string) int {
	total := 0
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, line := range input {
		lineSoFar := ""
	myloop:
		for _, char := range line {
			switch char {
			case '(', '{', '[', '<':
				lineSoFar += string(char)
			case ')', '}', ']', '>':
				if lineSoFar[len(lineSoFar)-1] == mapa[char] {
					lineSoFar = lineSoFar[:len(lineSoFar)-1]
				} else {
					total += scores[char]
					break myloop
				}
			}
		}
	}

	return total
}

func getMiddleScore(input []string) int {
	scores := []int{}
	scoreMap := map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	for _, line := range input {
		total := 0
		lineSoFar := ""
		brokenLine := false
	myloop:
		for _, char := range line {
			switch char {
			case '(', '{', '[', '<':
				lineSoFar += string(char)
			case ')', '}', ']', '>':
				if lineSoFar[len(lineSoFar)-1] == mapa[char] {
					lineSoFar = lineSoFar[:len(lineSoFar)-1]
				} else {
					brokenLine = true
					break myloop
				}
			}
		}
		if !brokenLine {
			for i := len(lineSoFar) - 1; i >= 0; i-- {
				total = (total * 5) + scoreMap[lineSoFar[i]]
			}
			scores = append(scores, total)
		}

	}
	sort.Ints(scores)

	return scores[len(scores)/2]
}
