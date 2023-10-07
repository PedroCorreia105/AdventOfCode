package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

type Number struct {
	value int
	next  *Number
	prev  *Number
}

func main() {
	input := utils.StringsToInts(utils.ReadFile(2022, 20, "\n"))
	fmt.Println("2022 Day 20")
	fmt.Println("\tPart 1:", calculate(input, 1, 1))
	fmt.Println("\tPart 2:", calculate(input, 10, 811589153))
}

func calculate(input []int, mixTimes, decodeKey int) (int) {
	size, total := len(input), 0
	list := []*Number{}
	var zeroNumber *Number

	// load each number
	for i := 0; i < size; i++ {
		list = append(list, &Number{value: input[i] * decodeKey})
		if (input[i] == 0) { zeroNumber = list[i] }
	}

	// add prev and next pointers
	for i := 0; i < size; i++ {
		(*list[i]).next = list[utils.Mod(i + 1, size)]
		(*list[i]).prev = list[utils.Mod(i - 1, size)]
	}
	
	for z := 0; z < mixTimes; z++ {
		for i := 0; i < size; i++ {	
			currentNumber := list[i]
			// remove 1 since we are removing the element from the linked list
			leapDistance := utils.Mod(currentNumber.value, size - 1)

			if (leapDistance == 0) { continue }

			// remove element from linked list
			(*(*list[i]).prev).next = (*list[i]).next
			(*(*list[i]).next).prev = (*list[i]).prev

			// find element that precedes currentNumber new location
			for f := 0; f < leapDistance; f++ {
				currentNumber = (*currentNumber).next
			}

			// insert in linked list
			(*list[i]).prev = currentNumber
			(*list[i]).next = (*currentNumber).next
			(*(*currentNumber).next).prev = list[i]
			(*currentNumber).next = list[i]
		}
	}
	
	currentNumber := zeroNumber 
	for f := 0; f < size; f++ {
		switch f {
		case 1000 % size, 2000 % size, 3000 % size:
			total += (*currentNumber).value
		}
		currentNumber = (*currentNumber).next
	}
	
	return total
}