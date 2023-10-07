package utils

import (
	"fmt"
)

func IsInMap(graph map[string]int, x, y int) bool {
	_, ok := graph[fmt.Sprintf("%d, %d", x, y)]
	return ok
}

func AddToMap(graph map[string]int, x, y, value int) {
	graph[fmt.Sprintf("%d, %d", x, y)] = value
}

func GetFromMap(graph map[string]int, x, y int) int {
	return graph[fmt.Sprintf("%d, %d", x, y)]
}

func IsInStringMap(graph map[string]string, x, y int) bool {
	_, ok := graph[fmt.Sprintf("%d, %d", x, y)]
	return ok
}

func AddToStringMap(graph map[string]string, x, y int, value string) {
	graph[fmt.Sprintf("%d, %d", x, y)] = value
}

func GetFromStringMap(graph map[string]string, x, y int) string {
	return graph[fmt.Sprintf("%d, %d", x, y)]
}
