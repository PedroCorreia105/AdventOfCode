package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
	"encoding/json"
	"sort"
)

func main() {
	input := utils.ReadFile(2022, 13, "\n\n")
	fmt.Println("2022 Day 13")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) (int) {
	total := 0
	var leftSide, rightSide any
	
	for pairIndex, pair := range input {
		pairSlice := strings.Split(pair, "\n")
		json.Unmarshal([]byte(pairSlice[0]), &leftSide)
		json.Unmarshal([]byte(pairSlice[1]), &rightSide)
		if (compare(leftSide, rightSide) < 0) {
			total += 1 + pairIndex
		}
	}
	
	return total
}

func part2(input []string) (int) {
	var leftSide, rightSide any
	list := []any{}
	total := 1
	marker1 := []byte("[[2]]")
	marker2 := []byte("[[6]]")

	for _, pair := range input {
		pairSlice := strings.Split(pair, "\n")
		json.Unmarshal([]byte(pairSlice[0]), &leftSide)
		json.Unmarshal([]byte(pairSlice[1]), &rightSide)
		list = append(list, leftSide, rightSide)
	}
	json.Unmarshal(marker1, &leftSide)
	json.Unmarshal(marker2, &rightSide)
	list = append(list, leftSide, rightSide)

	sort.Slice(list, func(i, j int) bool {
		return compare(list[i], list[j]) < 0
	})

	for i := 0; i < len(list); i++ {
		if compare(list[i], leftSide) == 0 || compare(list[i], rightSide) == 0 {
			total *= i + 1
		}
	}
	return total
}

func compare(leftSide, rightSide any) int {
	leftNumber, leftList, leftIsNumber := typeOf(leftSide)
	rightNumber, rightList, rightIsNumber := typeOf(rightSide)

	if (leftIsNumber && rightIsNumber) {
		return compareNumbers(leftNumber, rightNumber)
	} else if (leftIsNumber) {
		return compare([]any{leftNumber}, rightSide)
	} else 	if (rightIsNumber) {
		return compare(leftSide, []any{rightNumber})
	} else {
		for i := 0; i < utils.Min(len(leftList), len(rightList)); i++ {
			comparisonResult := compare(leftList[i], rightList[i])
			if (comparisonResult != 0) {
				return comparisonResult
			}
		}

		return compareNumbers(len(leftList), len(rightList))
	}
}

func typeOf(value any) (int, []any, bool) {
	switch valueType := value.(type) {
	case []any:
		return 0, valueType, false
	case int:
		return valueType, nil, true
	case float64:
		return int(valueType), nil, true
	}
	panic("Unable to continue")
}

func compareNumbers(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}