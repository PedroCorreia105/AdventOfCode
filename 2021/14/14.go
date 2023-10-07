package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)


func main() {
	input := utils.ReadFile(2021, 14, "\n\n")
	polymer := input[0]
	rules := strings.Split(input[1], "\n")

	fmt.Println("2021 Day 14")
	fmt.Println("\tPart 1:", getQuantities(polymer, rules, 10))
	fmt.Println("\tPart 2:", getQuantities(polymer, rules, 40))
}

func getQuantities(polymer string, rules []string, steps int) (int) {
	// {"AB":C}
	recipe := make(map[string]string)
	// count ocurrences of double chars {"AB":0}
	counter := make(map[string]int)
	
	// fill recipe map
	for _, rule := range rules {
		parts := strings.Split(rule, " -> ")
		recipe[parts[0]] = parts[1]
		counter[parts[0]] = 0
	}
	
	// fill initial state counter map
	for charIndex := 1; charIndex < len(polymer); charIndex++ {
		doubleChars := polymer[charIndex-1:charIndex+1]
		counter[doubleChars]++
	}

	for step := 0; step < steps; step++ {
		newCounter := make(map[string]int)
		for doubleChars, val := range counter {
			newCounter[string(doubleChars[0]) + recipe[doubleChars]] += val
			newCounter[recipe[doubleChars] + string(doubleChars[1])] += val
		}
		counter = newCounter
	}
	
	// count first character of each double
	charCounter := make(map[byte]int)
	for doubleChars, ammount := range counter {
		charCounter[doubleChars[0]] += ammount
	}
	// add last char
	charCounter[polymer[len(polymer) -1]] ++
	
	max, min := -1, -1
	for _, val := range charCounter {
		if min == -1 || min > val {
			min = val
		}
		
		if max == -1 || max < val {
			max = val
		}
	}

	return max - min
}