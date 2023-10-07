package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2021, 4, "\n\n")

	fmt.Println("2021 Day 04")
	fmt.Println("\tPart 1:", findWinningBoard(input))
	fmt.Println("\tPart 2:", findLastWinningBoard(input))
}

func findWinningBoard(input []string) int {
	numberSequence := utils.StringsToInts(strings.Split(input[0], ","))
	boards := make([][][]int, len(input[1:]))

	// transform input into an array of boards
	for boardIndex, board := range input[1:] {
		for _, strNumbers := range strings.Split(board, "\n") {
			boards[boardIndex] = append(boards[boardIndex], utils.StringsToInts(strings.Fields(strNumbers)))
		}
	}

	for _, number := range numberSequence {
		for boardIndex, board := range boards {
			// replace bingo number with -1
			for y, line := range board {
				for x, cell := range line {
					if cell == number {
						boards[boardIndex][y][x] = -1
					}
				}
			}

			if verifyBoard(board) {
				return calculateBoard(board) * number
			}
		}
	}
	return 0
}

func findLastWinningBoard(input []string) int {
	numberSequence := utils.StringsToInts(strings.Split(input[0], ","))
	boards := make([][][]int, len(input[1:]))
	winningBoards := make([]int, len(input[1:]))
	var lastWinningBoard int

	// transform input into an array of boards
	for boardIndex, board := range input[1:] {
		for _, strNumbers := range strings.Split(board, "\n") {
			boards[boardIndex] = append(boards[boardIndex], utils.StringsToInts(strings.Fields(strNumbers)))
		}
	}

	for _, number := range numberSequence {
		for boardIndex, board := range boards {
			// replace bingo number with -1
			for y, line := range board {
				for x, cell := range line {
					if cell == number {
						boards[boardIndex][y][x] = -1
					}
				}
			}

			if !utils.ContainsInt(winningBoards, boardIndex) && verifyBoard(board) {
				winningBoards = append(winningBoards, boardIndex)
				lastWinningBoard = calculateBoard(board) * number
			}
		}
	}
	return lastWinningBoard
}

func verifyBoard(board [][]int) bool {
	for a := 0; a < len(board); a++ {
		isBingoX := true
		isBingoY := true
		for b := 0; b < len(board); b++ {
			isBingoX = isBingoX && board[a][b] == -1
			isBingoY = isBingoY && board[b][a] == -1
		}

		if isBingoX || isBingoY {
			return true
		}
	}

	return false
}

func calculateBoard(board [][]int) int {
	total := 0

	for a := 0; a < len(board); a++ {
		for b := 0; b < len(board); b++ {
			if board[a][b] != -1 {
				total += board[a][b]
			}
		}
	}

	return total
}
