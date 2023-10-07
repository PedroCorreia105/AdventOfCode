package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

var board = make([][]int, 10)
var total = 0

func main() {
	input := utils.ReadFile(2021, 11, "\n")

	fmt.Println("2021 Day 11")
	fmt.Println("\tPart 1:", getFlashes(input))
	fmt.Println("\tPart 2:", getFlashyStep(input))
}

func getFlashes(input []string) int {
	for y, line := range input {
		for _, char := range line {
			board[y] = append(board[y], int(char-'0'))
		}
	}

	for step := 0; step < 100; step++ {
		for y, line := range input {
			for x := range line {
				board[y][x]++
			}
		}
		for y, line := range input {
			for x := range line {
				visit(x, y, 0)
			}
		}
	}

	return total
}

func getFlashyStep(input []string) int {
	step := 0
	isFlashy := false
	board = make([][]int, 10)
	for y, line := range input {
		for _, char := range line {
			board[y] = append(board[y], int(char-'0'))
		}
	}

	for !isFlashy {
		step++
		isFlashy = true

		for y, line := range input {
			for x := range line {
				board[y][x]++
			}
		}
		for y, line := range input {
			for x := range line {
				visit(x, y, 0)
			}
		}

		for y, line := range input {
			for x := range line {
				isFlashy = isFlashy && board[y][x] == 0
			}
		}
	}

	return step
}

func visit(x, y, increase int) {
	if 0 <= y && y <= 9 && 0 <= x && x <= 9 && board[y][x] != 0 {
		board[y][x] += increase
		if board[y][x] > 9 {
			board[y][x] = 0
			total++
			visit(x+1, y+1, 1)
			visit(x+1, y, 1)
			visit(x+1, y-1, 1)
			visit(x, y+1, 1)
			visit(x, y, 1)
			visit(x, y-1, 1)
			visit(x-1, y+1, 1)
			visit(x-1, y, 1)
			visit(x-1, y-1, 1)

		}
	}
}
