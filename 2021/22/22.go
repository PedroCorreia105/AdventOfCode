package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"sort"
)

type Instruction struct {
	status     bool
	x1, y1, z1 int
	x2, y2, z2 int
}

func main() {
	input := utils.ReadFile(2021, 22, "\n")

	fmt.Println("2021 Day 22")
	fmt.Println("\tPart 1:", calculate1(input))
	fmt.Println("\tPart 2:", calculate2(input))
}

func calculate1(input []string) int {
	var status string
	cubes := make(map[string]int)

	for _, line := range input {
		var i Instruction
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &status, &i.x1, &i.x2, &i.y1, &i.y2, &i.z1, &i.z2)
		i.x1 = utils.Max(i.x1, -50)
		i.x2 = utils.Min(i.x2, 50)
		i.y1 = utils.Max(i.y1, -50)
		i.y2 = utils.Min(i.y2, 50)
		i.z1 = utils.Max(i.z1, -50)
		i.z2 = utils.Min(i.z2, 50)

		for x := i.x1; x <= i.x2; x++ {
			for y := i.y1; y <= i.y2; y++ {
				for z := i.z1; z <= i.z2; z++ {
					key := fmt.Sprintf("%d %d %d", x, y, z)

					if status == "on" {
						cubes[key] = 1
					} else {
						delete(cubes, key)
					}
				}
			}
		}
	}
	return len(cubes)
}

func calculate2(input []string) uint64 {
	instructions := []Instruction{}

	xAxis := make([]int, 0, 2*len(input)+2)
	yAxis := make([]int, 0, 2*len(input)+2)
	zAxis := make([]int, 0, 2*len(input)+2)

	for _, line := range input {
		var status string
		var i Instruction
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &status, &i.x1, &i.x2, &i.y1, &i.y2, &i.z1, &i.z2)
		i.x2++
		i.y2++
		i.z2++
		if status == "on" {
			i.status = true
		} else {
			i.status = false
		}
		instructions = append(instructions, i)

		if !utils.ContainsInt(xAxis, i.x1) {
			xAxis = append(xAxis, i.x1)
		}
		if !utils.ContainsInt(xAxis, i.x2) {
			xAxis = append(xAxis, i.x2)
		}
		if !utils.ContainsInt(yAxis, i.y1) {
			yAxis = append(yAxis, i.y1)
		}
		if !utils.ContainsInt(yAxis, i.y2) {
			yAxis = append(yAxis, i.y2)
		}
		if !utils.ContainsInt(zAxis, i.z1) {
			zAxis = append(zAxis, i.z1)
		}
		if !utils.ContainsInt(zAxis, i.z2) {
			zAxis = append(zAxis, i.z2)
		}
	}

	sort.Ints(xAxis)
	sort.Ints(yAxis)
	sort.Ints(zAxis)
	xAxisLen := len(xAxis)
	yAxisLen := len(yAxis)

	cubeStatus := make([]bool, xAxisLen*len(yAxis)*len(zAxis))

	for _, i := range instructions {
		zIndex1 := utils.IndexOfInt(zAxis, i.z1)
		zIndex2 := utils.IndexOfInt(zAxis, i.z2)
		yIndex1 := utils.IndexOfInt(yAxis, i.y1)
		yIndex2 := utils.IndexOfInt(yAxis, i.y2)
		xIndex1 := utils.IndexOfInt(xAxis, i.x1)
		xIndex2 := utils.IndexOfInt(xAxis, i.x2)

		for z := zIndex1; z < zIndex2; z++ {
			offset := z * xAxisLen * yAxisLen
			for y := yIndex1; y < yIndex2; y++ {
				offset := y*xAxisLen + offset
				for x := xIndex1; x < xIndex2; x++ {
					cubeStatus[offset+x] = i.status
				}
			}
		}
	}

	var total uint64 = 0
	for z := 0; z < len(zAxis)-1; z++ {
		offset := z * xAxisLen * yAxisLen
		cubes := uint64(zAxis[z+1] - zAxis[z])
		for y := 0; y < yAxisLen-1; y++ {
			offset := y*xAxisLen + offset
			cubes := cubes * uint64(yAxis[y+1]-yAxis[y])
			for x := 0; x < xAxisLen-1; x++ {
				if cubeStatus[offset+x] {
					cubes := cubes * uint64(xAxis[x+1]-xAxis[x])
					total += cubes
				}
			}
		}
	}
	return total
}
