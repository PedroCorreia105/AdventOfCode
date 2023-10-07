package main

import (
	"AdventOfCode/2021/utils"
	"strings"
)

func getEnergyCost2(input []string) int {
	states := make(map[string]State)
	input = append(input[:3], append(ADDITIONAL_LINES, input[3:]...)...)
	currentStateText := strings.Join(input, "\n")
	currentState := State{text: currentStateText, costSoFar: 0}
	states[currentStateText] = currentState
	queue := []State{currentState}

	for currentState.text != TARGET_TEXT2 {
		// calculate neightbour states, add them to the queue
		for _, newState := range generateNewStates2(currentState) {
			if _, present := states[newState.text]; !present {
				queue = append(queue, newState)
			}
		}

		// remove visited state from queue
		queue = queue[1:]

		// remove visited states from queue
		currentPosition := 0
		for currentPosition < len(queue) {
			if _, present := states[queue[currentPosition].text]; present {
				queue[currentPosition] = queue[len(queue)-1]
				queue = queue[:len(queue)-1]
			} else {
				currentPosition += 1
			}
		}

		// make the first element of the queue the one with the lowest cost and visit it next loop
		if len(queue) > 0 {
			currentStateIndex := 0
			for i, state := range queue {
				if state.costSoFar-state.similarity < queue[currentStateIndex].costSoFar-queue[currentStateIndex].similarity {
					currentStateIndex = i
				}
			}
			currentState = queue[currentStateIndex]
			states[currentState.text] = currentState
			queue[currentStateIndex] = queue[0]
			queue[0] = currentState
		}
	}
	return currentState.costSoFar
}

// generates a list of new states that are reachable moving 1 piece from the current state
func generateNewStates2(currentState State) []State {
	area := strings.Split(currentState.text, "\n")
	newStates := []State{}

	// create states where a piece is removed from a room
	for _, roomColumn := range ROOM_COLUMNS {
		hallway := []int{roomColumn, HALLWAY_ROW}
		roomLevel1 := []int{roomColumn, ROOM_LEVEL_1_ROW}
		roomLevel2 := []int{roomColumn, ROOM_LEVEL_2_ROW}
		roomLevel3 := []int{roomColumn, ROOM_LEVEL_3_ROW}
		roomLevel4 := []int{roomColumn, ROOM_LEVEL_4_ROW}

		// if the first 3 rooms are empty and the last one is incorrect, remove the last piece
		if isEmpty(area, roomLevel1) && isEmpty(area, roomLevel2) && isEmpty(area, roomLevel3) && !belongsInTheRoom2(area, roomLevel4) {
			newStates = append(newStates, generateNewStatesToHallway2(area, currentState, roomLevel4, hallway)...)
			// if the first 2 rooms are empty and one of the others is incorrect, remove the third piece
		} else if isEmpty(area, roomLevel1) && isEmpty(area, roomLevel2) && (!belongsInTheRoom2(area, roomLevel3) || !belongsInTheRoom2(area, roomLevel4)) {
			newStates = append(newStates, generateNewStatesToHallway2(area, currentState, roomLevel3, hallway)...)
			// if the first room is empty and one of the others is incorrect, remove the second piece
		} else if isEmpty(area, roomLevel1) && (!belongsInTheRoom2(area, roomLevel2) || !belongsInTheRoom2(area, roomLevel3) || !belongsInTheRoom2(area, roomLevel4)) {
			newStates = append(newStates, generateNewStatesToHallway2(area, currentState, roomLevel2, hallway)...)
			// if any of the pieces is incorrect, remove the first piece
		} else if !belongsInTheRoom2(area, roomLevel1) || !belongsInTheRoom2(area, roomLevel2) || !belongsInTheRoom2(area, roomLevel3) || !belongsInTheRoom2(area, roomLevel4) {
			newStates = append(newStates, generateNewStatesToHallway2(area, currentState, roomLevel1, hallway)...)
		}
	}

	// create states where a piece is removed from the hallway
	for _, hallwayColumn := range HALLWAY_COLUMNS {
		hallway := []int{hallwayColumn, HALLWAY_ROW}
		if !isEmpty(area, hallway) {
			roomLevel1 := []int{PIECE_COLUMNS[area[hallway[1]][hallway[0]]], ROOM_LEVEL_1_ROW}
			roomLevel2 := []int{PIECE_COLUMNS[area[hallway[1]][hallway[0]]], ROOM_LEVEL_2_ROW}
			roomLevel3 := []int{PIECE_COLUMNS[area[hallway[1]][hallway[0]]], ROOM_LEVEL_3_ROW}
			roomLevel4 := []int{PIECE_COLUMNS[area[hallway[1]][hallway[0]]], ROOM_LEVEL_4_ROW}

			// try to put the piece in the level 4 room
			if isEmpty(area, roomLevel4) {
				if pathIsClear(area, hallway, roomLevel4) {
					newStates = append(newStates, generateNewState2(area, currentState, hallway, roomLevel4))
				}
				// try to put the piece in the level 3 room
			} else if belongsInTheRoom2(area, roomLevel4) {
				if isEmpty(area, roomLevel3) {
					if pathIsClear(area, hallway, roomLevel3) {
						newStates = append(newStates, generateNewState2(area, currentState, hallway, roomLevel3))
					}
					// try to put the piece in the level 2 room
				} else if belongsInTheRoom2(area, roomLevel3) {
					if isEmpty(area, roomLevel2) {
						if pathIsClear(area, hallway, roomLevel2) {
							newStates = append(newStates, generateNewState2(area, currentState, hallway, roomLevel2))
						}
						// try to put the piece in the level 1 room
					} else if belongsInTheRoom2(area, roomLevel2) && isEmpty(area, roomLevel1) {
						if pathIsClear(area, hallway, roomLevel1) {
							newStates = append(newStates, generateNewState2(area, currentState, hallway, roomLevel1))
						}
					}
				}
			}
		}
	}

	return newStates
}

func generateNewStatesToHallway2(area []string, currentState State, source, hallway []int) []State {
	newStates := []State{}
	localHallway := []int{hallway[0], hallway[1]}
	// go left
	for isEmpty(area, localHallway) {
		// cannot remain above room entrance
		if !utils.ContainsInt(ROOM_COLUMNS, localHallway[0]) {
			newStates = append(newStates, generateNewState2(area, currentState, source, localHallway))
		}
		localHallway[0] -= 1
	}
	localHallway = []int{hallway[0], hallway[1]}
	// go right
	for isEmpty(area, localHallway) {
		// cannot remain above room entrance
		if !utils.ContainsInt(ROOM_COLUMNS, localHallway[0]) {
			newStates = append(newStates, generateNewState2(area, currentState, source, localHallway))
		}
		localHallway[0] += 1
	}
	return newStates
}

func generateNewState2(area []string, currentState State, source, destination []int) State {
	startX := source[0]
	startY := source[1]
	endX := destination[0]
	endY := destination[1]
	distance := utils.Abs(endX-startX) + utils.Abs(endY-startY)
	areaCopy := utils.CopyStringSlice(area)
	startCharacter := areaCopy[startY][startX]
	areaCopy[endY] = areaCopy[endY][:endX] + string(startCharacter) + area[endY][endX+1:]
	areaCopy[startY] = areaCopy[startY][:startX] + "." + areaCopy[startY][startX+1:]
	newSimilarity := similarityToTarget2(areaCopy)
	return State{
		text:          strings.Join(areaCopy, "\n"),
		costSoFar:     currentState.costSoFar + distance*COSTS[startCharacter],
		similarity:    newSimilarity,
	}
}

func belongsInTheRoom2(area []string, room []int) bool {
	x := room[0]
	y := room[1]
	return area[y][x] == TARGET_STATE2[y][x]
}

// returns a value representing the similarity between two states
func similarityToTarget2(currentState []string) int {
	similarity := 0
	for _, roomColumn := range ROOM_COLUMNS {
		if currentState[ROOM_LEVEL_4_ROW][roomColumn] == TARGET_STATE2[ROOM_LEVEL_4_ROW][roomColumn] {
			similarity += 8000 * COSTS[currentState[ROOM_LEVEL_4_ROW][roomColumn]]

			if currentState[ROOM_LEVEL_3_ROW][roomColumn] == TARGET_STATE2[ROOM_LEVEL_3_ROW][roomColumn] {
				similarity += 400 * COSTS[currentState[ROOM_LEVEL_3_ROW][roomColumn]]

				if currentState[ROOM_LEVEL_2_ROW][roomColumn] == TARGET_STATE2[ROOM_LEVEL_2_ROW][roomColumn] {
					similarity += 20 * COSTS[currentState[ROOM_LEVEL_2_ROW][roomColumn]]

					if currentState[ROOM_LEVEL_1_ROW][roomColumn] == TARGET_STATE2[ROOM_LEVEL_1_ROW][roomColumn] {
						similarity += 1 * COSTS[currentState[ROOM_LEVEL_1_ROW][roomColumn]]
					}
				}
			}
		}
	}
	return similarity
}