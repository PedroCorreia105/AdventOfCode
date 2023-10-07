package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

type Monkey struct {
	objects           []int
	divisibleBy       int
	operation         string
	nextMonkeyIfTrue  int
	nextMonkeyIfFalse int
	inspectedTimes    int
}

func main() {
	monkeys := utils.ReadFile(2022, 11, "\n\n")
	fmt.Println("2022 Day 11")
	fmt.Println("\tPart 1:", calculate(monkeys, 1, 20))
	fmt.Println("\tPart 2:", calculate(monkeys, 2, 10_000))
}

func calculate(input []string, part int, rounds int) int {
	var monkeyNumber, divisibleBy, nextMonkey, nextMonkeyIfTrue, nextMonkeyIfFalse, result int
	var startingItems, operation string
	monkeys := make([]Monkey, len(input))
	maximums := make([]int, 2)
	megaModulo := 1
	format := "Monkey%d:\nStartingitems:%s\nOperation:new=%s\nTest:divisibleby%d\nIftrue:throwtomonkey%d\nIffalse:throwtomonkey%d"

	for _, monkey := range input {
		// remove spaces because sscanf doesn't parse spaces
		spacelessMonkey := strings.ReplaceAll(monkey, " ", "")
		fmt.Sscanf(spacelessMonkey, format, &monkeyNumber, &startingItems, &operation, &divisibleBy, &nextMonkeyIfTrue, &nextMonkeyIfFalse)
		monkeys[monkeyNumber] = Monkey{
			objects:           utils.StringsToInts(strings.Split(startingItems, ",")),
			divisibleBy:       divisibleBy,
			operation:         operation,
			nextMonkeyIfTrue:  nextMonkeyIfTrue,
			nextMonkeyIfFalse: nextMonkeyIfFalse,
			inspectedTimes:    0,
		}
		megaModulo *= divisibleBy
	}

	for round := 0; round < rounds; round++ {
		for monkeyIndex := range monkeys {
			for _, object := range monkeys[monkeyIndex].objects {
				result = performOperation(monkeys[monkeyIndex].operation, object)

				if part == 1 {
					result /= 3
				} else {
					result %= megaModulo
				}

				if result%monkeys[monkeyIndex].divisibleBy == 0 {
					nextMonkey = monkeys[monkeyIndex].nextMonkeyIfTrue
				} else {
					nextMonkey = monkeys[monkeyIndex].nextMonkeyIfFalse
				}
				monkeys[nextMonkey].objects = append(monkeys[nextMonkey].objects, result)
				monkeys[monkeyIndex].inspectedTimes++
			}
			monkeys[monkeyIndex].objects = []int{}
		}
	}

	for _, monkey := range monkeys {
		utils.InsertIfLarger(maximums, monkey.inspectedTimes)
	}
	return maximums[0] * maximums[1]
}

func performOperation(operationString string, value int) int {
	var op1, op2 int
	newString := strings.ReplaceAll(operationString, "old", utils.IntToString(value))
	if strings.Contains(newString, "*") {
		fmt.Sscanf(newString, "%d*%d", &op1, &op2)
		return op1 * op2
	}
	fmt.Sscanf(newString, "%d+%d", &op1, &op2)
	return op1 + op2
}
