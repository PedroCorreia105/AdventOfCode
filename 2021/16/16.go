package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ReadFile(2021, 16, "\n")[0]

	binary := ""
	for _, char := range input {
		result, _ := strconv.ParseUint(string(char), 16, 64) // hexadecimal string to uint64
		binary += fmt.Sprintf("%04b", result)                // uint64 to binary string
	}

	result, _, version := parseBinary(binary)

	fmt.Println("2021 Day 16")
	fmt.Println("\tPart 1:", version)
	fmt.Println("\tPart 2:", result)
}

func parseBinary(bits string) (int, int, int) {
	version := utils.BinToInt(bits[:3])
	typeId := utils.BinToInt(bits[3:6])

	results := []int{}
	var result int
	var subVersion int
	var bitsRead int
	var bitsReadSoFar int

	if typeId == 4 {
		final := ""
		for bitsReadSoFar = 6; bitsReadSoFar < len(bits); bitsReadSoFar += 5 {
			final += bits[bitsReadSoFar+1 : bitsReadSoFar+5]
			if bits[bitsReadSoFar] == '0' {
				bitsReadSoFar += 5
				break
			}
		}
		result = utils.BinToInt(final)
		return result, bitsReadSoFar, version
	}

	if bits[6] == '0' {
		numberOfBitsToRead := utils.BinToInt(bits[7:22])
		for bitsReadSoFar < numberOfBitsToRead {
			result, bitsRead, subVersion = parseBinary(bits[22+bitsReadSoFar : 22+numberOfBitsToRead])
			results = append(results, result)
			version += subVersion
			bitsReadSoFar += bitsRead
		}
		bitsReadSoFar += 22
	} else {
		numbersToRead := utils.BinToInt(bits[7:18])
		for i := 0; i < numbersToRead; i++ {
			result, bitsRead, subVersion = parseBinary(bits[18+bitsReadSoFar:])
			results = append(results, result)
			version += subVersion
			bitsReadSoFar += bitsRead
		}
		bitsReadSoFar += 18
	}

	switch typeId {
	case 0:
		result = 0
		for _, val := range results {
			result += val
		}
	case 1:
		result = 1
		for _, val := range results {
			result *= val
		}
	case 2:
		result = -1
		for _, val := range results {
			if result == -1 || val < result {
				result = val
			}
		}
	case 3:
		result = -1
		for _, val := range results {
			if result == -1 || val > result {
				result = val
			}
		}
	case 5:
		if results[0] > results[1] {
			result = 1
		} else {
			result = 0
		}
	case 6:
		if results[0] < results[1] {
			result = 1
		} else {
			result = 0
		}
	case 7:
		if results[0] == results[1] {
			result = 1
		} else {
			result = 0
		}
	}

	return result, bitsReadSoFar, version
}
