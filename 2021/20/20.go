package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

var vectors = [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}


func main() {
	input := utils.ReadFile(2021, 20, "\n\n")
	area := strings.Split(input[1], "\n")

	fmt.Println("2021 Day 20")
	fmt.Println("\tPart 1:", decode(input[0], area, 2))
	fmt.Println("\tPart 2:", decode(input[0], area, 50))

}

func decode(cypher string, area []string, times int) (int) {
	borderSize := 2
	newArea := make([]string, len(area) + 2*borderSize)

	for y := -borderSize; y < len(area) + borderSize; y++ {
		for x := -borderSize; x < len(area[0]) + borderSize; x++ {
			generatedString := ""
			for _, vector := range vectors {
				nextX := x + vector[0]
				nextY := y + vector[1]
				if 0 <= nextX && nextX < len(area[0]) && 0 <= nextY && nextY < len(area){
					generatedString += string(area[nextY][nextX])
				} else if times%2 == 0 {
					generatedString += "."
				} else {
					generatedString += "#"
				}
			}
			newArea[y + borderSize] += string(cypher[decodeString(generatedString)])
		}
	}


	if times > 1 {
		return decode(cypher, newArea, times -1)
	} else {
		total := 0
		for _, line := range newArea[borderSize : len(newArea) - borderSize] {
			for _, char := range line[borderSize : len(line) - borderSize] {
				if char == '#' {
					total += 1
				}
			}
		}
		return total
	}
}


func decodeString(code string) (int) {
	code = strings.ReplaceAll(code, "#", "1")
	code = strings.ReplaceAll(code, ".", "0")
	return utils.BinToInt(code)
}