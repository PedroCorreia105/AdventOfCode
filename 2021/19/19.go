package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

type Scanner struct {
	coordinates [][3]int
	location    [3]int
}

func main() {
	input := utils.ReadFile(2021, 19, "\n\n")

	fmt.Println("2021 Day 19")
	p1, p2 := calculate(input)
	fmt.Println("\tPart 1:", p1)
	fmt.Println("\tPart 2:", p2)
}

func calculate(input []string) (int, int) {
	var x, y, z int
	scanners := make([]*Scanner, len(input))

	for scannerIndex, scannerInput := range input {
		coordinateLines := strings.Split(scannerInput, "\n")
		scanners[scannerIndex] = &Scanner{coordinates: make([][3]int, len(coordinateLines)-1)}
		for lineIndex, line := range coordinateLines[1:] {
			fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
			scanners[scannerIndex].coordinates[lineIndex] = [3]int{x, y, z}
		}
	}

	// set to register the aligned coordinates
	beacons := make(map[string]bool)
	// assume the first scanner is aligned at 0, 0, 0
	scanners[0].location = [3]int{0, 0, 0}
	allAlignedScanners := []*Scanner{scanners[0]}
	queueAlignedScanners := []*Scanner{scanners[0]}
	unalignedScanners := scanners[1:]

	for len(unalignedScanners) > 0 {
		// pop the last scanner
		scanner1 := queueAlignedScanners[len(queueAlignedScanners)-1]
		queueAlignedScanners = queueAlignedScanners[:len(queueAlignedScanners)-1]
		unalignedScannersTemp := []*Scanner{}

		for _, scanner2 := range unalignedScanners {
			// try to align scanners
			alignedScanner := alignScanners(scanner1, scanner2)

			// if it wasn't possible
			if alignedScanner == nil {
				unalignedScannersTemp = append(unalignedScannersTemp, scanner2)
			} else {
				queueAlignedScanners = append(queueAlignedScanners, alignedScanner)
				allAlignedScanners = append(allAlignedScanners, alignedScanner)
			}
		}

		unalignedScanners = unalignedScannersTemp
	}

	furthestScanners := 0
	for _, scanner1 := range allAlignedScanners {
		for _, coordinate := range scanner1.coordinates {
			beacons[fmt.Sprintf("%d,%d,%d", coordinate[0], coordinate[1], coordinate[2])] = true
		}

		for _, scanner2 := range allAlignedScanners {
			distance := 0
			for i := 0; i < 3; i++ {
				distance += utils.Abs(scanner1.location[i] - scanner2.location[i])
			}

			if distance > furthestScanners {
				furthestScanners = distance
			}
		}
	}

	return len(beacons), furthestScanners
}

func alignScanners(scanner1, scanner2 *Scanner) *Scanner {
	var alignedTargetValues [3][]int
	var offsetCoordinates [3]int
	foundAxis := make(map[int]bool)
	axisValues1 := make([]int, len(scanner1.coordinates))
	axisValues2 := make([]int, len(scanner2.coordinates))

	for axis1 := 0; axis1 < 3; axis1++ {
		numSharedBeacons := -1
		offset := 0
		axis2 := 0

		for i, coordinate := range scanner1.coordinates {
			axisValues1[i] = coordinate[axis1]
		}

		for _, signedInt := range []int{1, -1} {
			for axis2 = 0; axis2 < 3; axis2++ {
				if _, found := foundAxis[axis2]; found {
					continue
				}

				for i, coordinate := range scanner2.coordinates {
					axisValues2[i] = coordinate[axis2] * signedInt
				}

				// calculate the difference between all possible combinations of values for each axis of each scanner 
				differences := []int{}
				for _, value1 := range axisValues1 {
					for _, value2 := range axisValues2 {
						differences = append(differences, value1-value2)
					}
				}

				// find the most common difference and see if it is at least 12
				offset, numSharedBeacons = mostCommon(differences)

				if numSharedBeacons >= 12 {
					break
				}
			}
			if numSharedBeacons >= 12 {
				break
			}
		}

		if numSharedBeacons < 12 {
			return nil
		}

		foundAxis[axis2] = true
		alignedTargetValues[axis1] = make([]int, len(axisValues2))
		offsetCoordinates[axis1] = offset

		// convert the coordinates based on the offset from the origin
		for i, targetVal := range axisValues2 {
			alignedTargetValues[axis1][i] = targetVal + offset
		}
	}
	return &Scanner{coordinates: transpose(alignedTargetValues), location: offsetCoordinates}
}

func mostCommon(slice []int) (int, int) {
	counts := make(map[int]int)
	for _, i := range slice {
		counts[i]++
	}

	maxCount := -1
	mostCommon := 0
	for i, count := range counts {
		if count > maxCount {
			maxCount = count
			mostCommon = i
		}
	}

	return mostCommon, maxCount
}

func transpose(matrix [3][]int) [][3]int {
	if len(matrix) == 0 {
		return nil
	}

	transposed := make([][3]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = [3]int{}
	}

	for i, row := range matrix {
		for j, val := range row {
			transposed[j][i] = val
		}
	}

	return transposed
}
