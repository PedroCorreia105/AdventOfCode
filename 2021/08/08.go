package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2021, 8, "\n")

	fmt.Println("2021 Day 08")
	fmt.Println("\tPart 1:", decode(input, 1))
	fmt.Println("\tPart 2:", decode(input, 2))
}

func decode(input []string, version int) int {
	total := 0

	for _, line := range input {
		parts := strings.Split(line, " | ")
		part1 := strings.Split(parts[0], " ")
		part2 := strings.Split(parts[1], " ")
		decoder := make([]string, 10)

		// loop while it hasn't decoded all numbers
		for utils.ContainsString(decoder, "") {
			for _, code := range part1 {
				if len(code) == 3 {
					decoder[7] = utils.Sort(code)
				} else if len(code) == 4 {
					decoder[4] = utils.Sort(code)
				} else if len(code) == 2 {
					decoder[1] = utils.Sort(code)
				} else if len(code) == 7 {
					decoder[8] = utils.Sort(code)
				} else if len(code) == 6 && utils.NumberOfCommonCharacters(code, decoder[1]) == 1 {
					decoder[6] = utils.Sort(code)
				} else if len(code) == 6 && utils.NumberOfCommonCharacters(code, decoder[4]) == 4 {
					decoder[9] = utils.Sort(code)
				} else if len(code) == 6 {
					decoder[0] = utils.Sort(code)
				} else if len(code) == 5 && utils.NumberOfCommonCharacters(code, decoder[1]) == 2 {
					decoder[3] = utils.Sort(code)
				} else if len(code) == 5 && utils.NumberOfCommonCharacters(code, decoder[6]) == 4 {
					decoder[2] = utils.Sort(code)
				} else if len(code) == 5 && utils.NumberOfCommonCharacters(code, decoder[6]) == 5 {
					decoder[5] = utils.Sort(code)
				}
			}
		}

		if version == 1 {
			for _, code := range part2 {
				switch utils.Sort(code) {
				case decoder[1], decoder[4], decoder[7], decoder[8]:
					total++
				}
			}
		} else {
			n := 0
			for _, code := range part2 {
				code = utils.Sort(code)
				for i := 0; i < 10; i++ {
					if code == decoder[i] {
						n = 10*n + i
					}
				}
			}
			total += n
		}
	}

	return total
}
