package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
	"errors"
)

var operations = make(map[string][]string)
var values = make(map[string]int)

func main() {
	input := utils.ReadFile(2022, 21, "\n")
	fmt.Println("2022 Day 21")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) (int) {
	for _, line := range(input) {
		parsedLine := strings.Split(line, " ")
		if (len(parsedLine) > 2) {
			operations[parsedLine[0][:4]] = parsedLine[1:]
		} else {
			values[parsedLine[0][:4]] = utils.StringToInt(parsedLine[1])
		}
	}
	result, _ := recursiveFind("root")
	return result
}

func part2(input []string) (int) {
	limits := []int{0, 0}
	half := 0

	for _, line := range(input) {
		parsedLine := strings.Split(line, " ")
		if (len(parsedLine) > 2) {
			operations[parsedLine[0][:4]] = parsedLine[1:]
		} else if (parsedLine[0] != "humn:") {
			values[parsedLine[0][:4]] = utils.StringToInt(parsedLine[1])
		}
	}
	
	//lets assume the first half requires the humn code
	provisionalHalf, _ := recursiveFind(operations["root"][0])
	goodHalf, err2 := recursiveFind(operations["root"][2])

	// in case of wrong assumption
	if (err2 != nil) { goodHalf, half, provisionalHalf = provisionalHalf, 2, 0 }
	
	limits[1] = goodHalf

	// log(n) brute force assigns a value to humn and checks if it's right.
	for limits[0] < limits[1]-1 {
		values["humn"] = (limits[0] + limits[1])/2
		provisionalHalf, _ = recursiveFind(operations["root"][half])
		if (provisionalHalf > goodHalf) {
			limits[0] = values["humn"]
		} else {
			limits[1] = values["humn"]
		}
	}
	return values["humn"]
}

func recursiveFind(key string) (int, error) {
	if _, ok := values[key]; ok {
		return values[key], nil
	}
	if _, ok := operations[key]; !ok {
		return 0, errors.New("Error")
	}

	op1, err1 := recursiveFind(operations[key][0])
	op2, err2 := recursiveFind(operations[key][2])

	if (err1 != nil || err2 != nil) { return 0, errors.New("Error") }

	switch operations[key][1] {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		return op1 / op2, nil
	}
	return 0, errors.New("Error")
}