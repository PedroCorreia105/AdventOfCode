package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

type Coordinate struct {
	x   int
	y   int
}

type Area struct {
	sensor    Coordinate
	beacon    Coordinate
	distance  int
}

func main() {
	input := utils.ReadFile(2022, 15, "\n")
	fmt.Println("2022 Day 15")
	fmt.Println("\tPart 1:", part1(input))
	fmt.Println("\tPart 2:", part2(input))
}

func part1(input []string) (int) {
	areas := []Area{}
	total := 0
	lineY := 2000000
	
	for _, line := range(input) {
		sensor := Coordinate{}
		beacon := Coordinate{}
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		distance := calculateManhattanDistance(sensor, beacon)
		areas = append(areas, Area {
			sensor: sensor,
			beacon: beacon,
			distance: distance,
		})
	}

	for pointX := -1000000; pointX <= 10000000; pointX++ {
		cannnotBeBeacon := false
		for _, area := range(areas) {
			// cannot be a beacon if it is closer to a sensor than its beacon
			if (calculateManhattanDistance(area.sensor, Coordinate{x: pointX, y: lineY}) <= area.distance) {
				cannnotBeBeacon = true
			}
			// remove coordinates that are beacons
			if (area.beacon.x == pointX && area.beacon.y == lineY) {
				cannnotBeBeacon = false
				break
			}
		}
		if (cannnotBeBeacon) {
			total++
		}
	}
	return total
}

func part2(input []string) (int) {
	areas := []Area{}
	borderCoordinates := []Coordinate{}
	
	for _, line := range(input) {
		sensor := Coordinate{}
		beacon := Coordinate{}
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		distance := calculateManhattanDistance(sensor, beacon)
		areas = append(areas, Area {
			sensor: sensor,
			beacon: beacon,
			distance: distance,
		})
		// the solution beacon is +1 from the border of an area
		distance++

		// top left diamond borders
		for y, x := sensor.y, sensor.x - distance; y > sensor.y - distance; y, x = y-1, x+1 {
			borderCoordinates = addCoordinate(borderCoordinates, x, y)
		}
		// top right diamond borders
		for y, x := sensor.y - distance, sensor.x; x < sensor.x + distance; y, x = y+1, x+1 {
			borderCoordinates = addCoordinate(borderCoordinates, x, y)
		}
		// bottom right diamond borders
		for y, x := sensor.y, sensor.x + distance; y < sensor.y + distance; y, x = y+1, x-1 {
			borderCoordinates = addCoordinate(borderCoordinates, x, y)
		}
		// bottom left diamond borders
		for y, x := sensor.y + distance, sensor.x; x > sensor.x - distance; y, x = y-1, x-1 {
			borderCoordinates = addCoordinate(borderCoordinates, x, y)
		}
	}

	for _, possibleSolution := range(borderCoordinates) {
		cannnotBeBeacon := false
		for _, area := range(areas) {
			// cannot be a beacon if it is closer to a sensor than its beacon
			if (calculateManhattanDistance(area.sensor, possibleSolution) <= area.distance) {
				cannnotBeBeacon = true
			}
			// remove coordinates that are beacons
			if (area.beacon.x == possibleSolution.x && area.beacon.y == possibleSolution.y) {
				cannnotBeBeacon = false
				break
			}
		}
		if (!cannnotBeBeacon) {
			return possibleSolution.x * 4000000 + possibleSolution.y
		}
	}
	panic("No solution found")
}

func calculateManhattanDistance(c1, c2 Coordinate) int {
	return utils.Abs(c1.x - c2.x) + utils.Abs(c1.y - c2.y)
}

func addCoordinate(borderCoordinates []Coordinate, x, y int) []Coordinate {
	if (x >= 0 && x <= 4000000 && y >= 0 && y <= 4000000) {
		borderCoordinates = append(borderCoordinates, Coordinate {x: x, y: y})
	}
	return borderCoordinates
}