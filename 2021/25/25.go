package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile(2021, 25, "\n")

	fmt.Println("2021 Day 25")
	fmt.Println("\tPart 1:", moveCucumber(input))
}

func moveCucumber(area []string) int {
	height, width := len(area), len(area[0])
	areaCopy := utils.CopyStringSlice(area)
	changeMade := true
	step := 0

	for changeMade {
		step++
		changeMade = false

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if area[y][x] == '>' && area[y][(x+1)%width] == '.' {
					changeMade = true
					areaCopy[y] = utils.ReplaceAtIndex(areaCopy[y], '.', x)
					areaCopy[y] = utils.ReplaceAtIndex(areaCopy[y], '>', (x+1)%width)
				}
			}
		}
		area = utils.CopyStringSlice(areaCopy)

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if area[y][x] == 'v' && area[(y+1)%height][x] == '.' {
					changeMade = true
					areaCopy[y] = utils.ReplaceAtIndex(areaCopy[y], '.', x)
					areaCopy[(y+1)%height] = utils.ReplaceAtIndex(areaCopy[(y+1)%height], 'v', x)
				}
			}
		}
		area = utils.CopyStringSlice(areaCopy)
	}
	return step
}
