package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
	"strings"
)

var TARGET_STATE1 = []string{
	"#############",
	"#...........#",
	"###A#B#C#D###",
	"  #A#B#C#D#",
	"  #########",
}
var TARGET_STATE2 = []string{
	"#############",
	"#...........#",
	"###A#B#C#D###",
	"  #A#B#C#D#",
	"  #A#B#C#D#",
	"  #A#B#C#D#",
	"  #########",
}
var TARGET_TEXT1 = strings.Join(TARGET_STATE1, "\n")
var TARGET_TEXT2 = strings.Join(TARGET_STATE2, "\n")
var ADDITIONAL_LINES = []string{
	"  #D#C#B#A#",
	"  #D#B#A#C#",
}
var COSTS = map[byte]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}
var PIECE_COLUMNS = map[byte]int{
	'A': 3,
	'B': 5,
	'C': 7,
	'D': 9,
}
var HALLWAY_ROW = 1
var ROOM_LEVEL_1_ROW = 2
var ROOM_LEVEL_2_ROW = 3
var ROOM_LEVEL_3_ROW = 4
var ROOM_LEVEL_4_ROW = 5
var ROOM_COLUMNS = []int{3, 5, 7, 9}
var HALLWAY_COLUMNS = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

type State struct {
	text          string
	costSoFar     int
	similarity    int
	previousState *State
}

func main() {
	input := utils.ReadFile(2021, 23, "\n")
	fmt.Println("2021 Day 23")
	fmt.Println("\tPart 1:", getEnergyCost1(input))
	fmt.Println("\tPart 2:", getEnergyCost2(input))
}
