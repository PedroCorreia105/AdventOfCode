package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile(2022, 6, "\n")[0]
	fmt.Println("2022 Day 06")
	fmt.Println("\tPart 1:", calculate(input, 4))
	fmt.Println("\tPart 2:", calculate(input, 14))
}

func calculate(input string, size int) (int) {
	for charIndex := 0; charIndex < len(input); charIndex++ {
		mapa := make(map[byte]int)
		next := false
		for i := 0; i < size; i++ {
			if _, ok := mapa[input[charIndex + i]]; ok {
				next = true
			} else {
				mapa[input[charIndex + i]] = 1
			}
		}
		if (!next) {
			return charIndex + size
		} 
	} 
	return -1
}