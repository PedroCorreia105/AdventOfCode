package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile(2021, 17, "\n")[0]
	var x1, x2, y1, y2 int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	highest := -1
	total := 0
	// #BruteForce ftw
	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			canReach, height := canReachTarget(x, y, x1, x2, y1, y2)
			if canReach {
				total += 1
				if highest == -1 || height > highest {
					highest = height
				}
			}
		}
	}

	fmt.Println("2021 Day 17")
	fmt.Println("\tPart 1:", highest)
	fmt.Println("\tPart 2:", total)
}

func canReachTarget(velocityX, velocityY, minX, maxX, minY, maxY int) (bool, int) {
	currentX, currentY := 0, 0
	highestY := -1
	for currentX <= maxX && currentY >= minY {
		
		currentX += velocityX
		currentY += velocityY
		
		if highestY == -1 || currentY > highestY {
			highestY = currentY
		}

		if currentX >= minX && currentX <= maxX && currentY >= minY && currentY <= maxY {
			return true, highestY
		}

		if velocityX > 0 {
			velocityX -=1
		} else if velocityX < 0 {
			velocityX +=1
		}
		velocityY -=1
	}
	return false, highestY
}