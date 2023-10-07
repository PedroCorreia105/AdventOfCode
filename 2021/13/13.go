package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadFile(2021, 13, "\n\n")
	points := strings.Split(input[0], "\n")
	folds := strings.Split(input[1], "\n")

	fmt.Println("2021 Day 13")
	fmt.Println("\tPart 1:", getPoints(points, folds))
	fmt.Println("\tPart 2:")
	printCode(points, folds)
}

func getPoints(points, folds []string) int {
	area := [][]int{}

	for _, point := range points {
		coordinates := strings.Split(point, ",")
		area = append(area, utils.StringsToInts(coordinates))
	}

	fold := folds[0]
	assignement := strings.Split(fold, " ")[2]
	assignementValues := strings.Split(assignement, "=")
	value := utils.StringToInt(assignementValues[1])
	for i, point := range area {
		if assignementValues[0] == "x" && point[0] > value {
			area[i][0] = 2*value - point[0]
		} else if assignementValues[0] == "y" && point[1] > value {
			area[i][1] = 2*value - point[1]
		}
	}

	duplicates := 0
	for i1, point1 := range area {
		for i2, point2 := range area {
			if i1 != i2 && point1[0] == point2[0] && point1[1] == point2[1] {
				duplicates++
			}
		}
	}

	return len(area) - duplicates/2
}

func printCode(points, folds []string) {
	area := [][]int{}

	for _, point := range points {
		coordinates := strings.Split(point, ",")
		area = append(area, utils.StringsToInts(coordinates))
	}

	for _, fold := range folds {
		assignement := strings.Split(fold, " ")[2]
		assignementValues := strings.Split(assignement, "=")
		value := utils.StringToInt(assignementValues[1])
		for i, point := range area {
			if assignementValues[0] == "x" && point[0] > value {
				area[i][0] = 2*value - point[0]
			} else if assignementValues[0] == "y" && point[1] > value {
				area[i][1] = 2*value - point[1]
			}
		}
	}

	graph := make(map[string]string)
	maxX, maxY := -1, -1
	for _, point := range area {
		if maxX == -1 || point[0] > maxX {
			maxX = point[0]
		}

		if maxY == -1 || point[1] > maxY {
			maxY = point[1]
		}

		utils.AddToStringMap(graph, point[0], point[1], "#")
	}

	for y := 0; y <= maxY; y++ {
		line := ""
		for x := 0; x <= maxX; x++ {
			if utils.IsInStringMap(graph, x, y) {
				line += utils.GetFromStringMap(graph, x, y)
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}
