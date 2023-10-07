package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

type Tree struct {
	x       int
	y       int
	visible bool
	height  int
}

func (tree *Tree) TravelDown(mapa [][]*Tree, largestSoFar int) {
	if tree.y == 0 || largestSoFar < tree.height {
		tree.visible = true
	}
	if tree.y < len(mapa)-1 {
		(*mapa[tree.y+1][tree.x]).TravelDown(mapa, utils.Max(largestSoFar, tree.height))
	}
}

func (tree *Tree) TravelUp(mapa [][]*Tree, largestSoFar int) {
	if tree.y == len(mapa)-1 || largestSoFar < tree.height {
		tree.visible = true
	}
	if tree.y > 0 {
		(*mapa[tree.y-1][tree.x]).TravelUp(mapa, utils.Max(largestSoFar, tree.height))
	}
}

func (tree *Tree) TravelLeft(mapa [][]*Tree, largestSoFar int) {
	if tree.x == len(mapa)-1 || largestSoFar < tree.height {
		tree.visible = true
	}
	if tree.x > 0 {
		(*mapa[tree.y][tree.x-1]).TravelLeft(mapa, utils.Max(largestSoFar, tree.height))
	}
}

func (tree *Tree) TravelRight(mapa [][]*Tree, largestSoFar int) {
	if tree.x == 0 || largestSoFar < tree.height {
		tree.visible = true
	}
	if tree.x < len(mapa)-1 {
		(*mapa[tree.y][tree.x+1]).TravelRight(mapa, utils.Max(largestSoFar, tree.height))
	}
}

func (tree *Tree) SizeDown(mapa [][]*Tree, originalHeight int, total int) int {
	if tree.y == len(mapa)-1 || originalHeight <= tree.height {
		return 1 + total
	} else {
		return (*mapa[tree.y+1][tree.x]).SizeDown(mapa, originalHeight, total+1)
	}
}

func (tree *Tree) SizeUp(mapa [][]*Tree, originalHeight int, total int) int {
	if tree.y == 0 || originalHeight <= tree.height {
		return 1 + total
	} else {
		return (*mapa[tree.y-1][tree.x]).SizeUp(mapa, originalHeight, total+1)
	}
}

func (tree *Tree) SizeLeft(mapa [][]*Tree, originalHeight int, total int) int {
	if tree.x == 0 || originalHeight <= tree.height {
		return 1 + total
	} else {
		return (*mapa[tree.y][tree.x-1]).SizeLeft(mapa, originalHeight, total+1)
	}
}

func (tree *Tree) SizeRight(mapa [][]*Tree, originalHeight int, total int) int {
	if tree.x == len(mapa)-1 || originalHeight <= tree.height {
		return 1 + total
	} else {
		return (*mapa[tree.y][tree.x+1]).SizeRight(mapa, originalHeight, total+1)
	}
}

func main() {
	input := utils.ReadFile(2022, 8, "\n")
	fmt.Println("2022 Day 08")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) int {
	mapa := make([][]*Tree, len(input))
	total := 0
	for y := 0; y < len(input); y++ {
		mapa[y] = make([]*Tree, len(input))
		for x := 0; x < len(input); x++ {
			mapa[y][x] = &Tree{
				x:       x,
				y:       y,
				visible: false,
				height:  utils.StringToInt(string(input[y][x])),
			}
		}
	}

	for i := 0; i < len(input); i++ {
		(*mapa[0][i]).TravelDown(mapa, -1)
		(*mapa[len(mapa)-1][i]).TravelUp(mapa, -1)
		(*mapa[i][0]).TravelRight(mapa, -1)
		(*mapa[i][len(mapa)-1]).TravelLeft(mapa, -1)
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			if (*mapa[y][x]).visible {
				total++
			}
		}
	}

	return total
}

func part2(input []string) int {
	mapa := make([][]*Tree, len(input))
	max := 0
	for y := 0; y < len(input); y++ {
		mapa[y] = make([]*Tree, len(input))
		for x := 0; x < len(input); x++ {
			mapa[y][x] = &Tree{
				x:       x,
				y:       y,
				visible: false,
				height:  utils.StringToInt(string(input[y][x])),
			}
		}
	}

	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input)-1; x++ {
			tree := (*mapa[y][x])
			a := (*mapa[y+1][x]).SizeDown(mapa, tree.height, 0)
			b := (*mapa[y-1][x]).SizeUp(mapa, tree.height, 0)
			c := (*mapa[y][x+1]).SizeRight(mapa, tree.height, 0)
			d := (*mapa[y][x-1]).SizeLeft(mapa, tree.height, 0)
			max = utils.Max(max, a*b*c*d)
		}
	}

	return max
}
