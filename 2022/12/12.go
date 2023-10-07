package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

type Square struct {
	x        int
	y        int
	height   byte
	distance int
	visited  bool
}

func main() {
	input := utils.ReadFile(2022, 12, "\n")
	fmt.Println("2022 Day 12")
	fmt.Println("\tPart 1:", calculate(input, 1))
	fmt.Println("\tPart 2:", calculate(input, 2))
}

func calculate(input []string, part int) int {
	var vectors = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var starting, ending, lowest *Square
	condition := true
	mapa := make(map[int]*Square)
	height, width := len(input), len(input[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mapa[y*width+x] = &Square{
				x:        x,
				y:        y,
				height:   input[y][x],
				distance: -1,
				visited:  false,
			}
			if input[y][x] == 'S' {
				(*mapa[y*width+x]).height = 'a'
				starting = mapa[y*width+x]
			}
			if input[y][x] == 'E' {
				(*mapa[y*width+x]).height = 'z'
				ending = mapa[y*width+x]
			}
		}
	}

	if part == 2 { // in part 2 we start from the end
		starting = ending
	}

	(*starting).distance = 0
	queue := []*Square{starting}

	for condition {
		lowest = queue[0]
		lowestIndex := 0
		for i, queueElement := range queue {
			if (*queueElement).distance < (*lowest).distance {
				lowest = queueElement
				lowestIndex = i
			}
		}
		queue = remove(queue, lowestIndex)
		(*lowest).visited = true

		for _, vector := range vectors {
			x := (*lowest).x + vector[0]
			y := (*lowest).y + vector[1]
			// if next square is within the map and wasn't visited yet
			if 0 <= x && x < width && 0 <= y && y < height && !(*mapa[y*width+x]).visited {
				// if can climb or descend (depends on part)
				if part == 1 && (*mapa[y*width+x]).height <= (*lowest).height+1 || part == 2 && (*lowest).height <= (*mapa[y*width+x]).height+1 {
					// if this path is shorter than a previously found one
					if (*mapa[y*width+x]).distance == -1 || (*lowest).distance+1 < (*mapa[y*width+x]).distance {
						(*mapa[y*width+x]).distance = (*lowest).distance + 1
						queue = append(queue, mapa[y*width+x])
					}
				}
			}
		}

		if part == 1 {
			condition = (*ending).distance == -1
		} else {
			condition = (*lowest).height != 'a'
		}
	}

	if part == 1 {
		return (*ending).distance
	} else {
		return (*lowest).distance
	}
}

func remove(s []*Square, i int) []*Square {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
