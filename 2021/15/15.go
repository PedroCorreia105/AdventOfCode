package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

var vectors = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Node struct {
	x         int
	y         int
	cost      int
	costSoFar int
	visited   bool
}

func main() {
	input := utils.ReadFile(2021, 15, "\n")

	fmt.Println("2021 Day 15")
	fmt.Println("\tPart 1:", getLowestRisk(input, 1))
	fmt.Println("\tPart 2:", getLowestRisk(input, 2))
}

func getLowestRisk(input []string, version int) int {
	size := len(input)
	if version == 2 {
		size *= 5
	}
	nodes := make([][]Node, size)
	var cost int

	for y := 0; y < size; y++ {
		nodes[y] = make([]Node, size)
		for x := 0; x < size; x++ {
			if version == 1 {
				cost = utils.StringToInt(string(input[y][x]))
			} else {
				cost = utils.StringToInt(string(input[y%len(input)][x%len(input)])) + x/len(input) + y/len(input)
				if cost > 9 {
					cost -= 9
				}
			}

			nodes[y][x] = Node{
				x:         x,
				y:         y,
				cost:      cost,
				costSoFar: 0,
				visited:   false,
			}
		}
	}

	currentPosition := 0
	currentNode := nodes[0][0]
	currentNode.visited = true
	queue := []Node{currentNode}

	for currentNode.x < size-1 || currentNode.y < size-1 {

		// add neightbour nodes to queue and calculate their costs
		for _, vector := range vectors {
			nextX := currentNode.x + vector[0]
			nextY := currentNode.y + vector[1]
			if 0 <= nextX && nextX < size && 0 <= nextY && nextY < size && !nodes[nextY][nextX].visited {
				nodes[nextY][nextX].visited = true
				nodes[nextY][nextX].costSoFar = currentNode.costSoFar + nodes[nextY][nextX].cost
				queue = append(queue, nodes[nextY][nextX])
			}
		}

		// remove visited node from queue
		queue[currentPosition] = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		// choose best next node to visit
		currentPosition = -1
		for index, queueNode := range queue {
			if currentPosition == -1 || queueNode.costSoFar < currentNode.costSoFar {
				currentNode = queueNode
				currentPosition = index
			}
		}
	}

	return currentNode.costSoFar
}
