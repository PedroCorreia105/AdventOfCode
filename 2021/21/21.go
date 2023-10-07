package main

import (
	"AdventOfCode/2021/utils"
	"fmt"
)

type Universe struct {
	playerPosition []int
	playerPoints   []int
	amount         uint64
}

func main() {
	input := utils.ReadFile(2021, 21, "\n")
	playersPosition := []int{4, 8}
	fmt.Sscanf(input[0], "Player 1 starting position: %d", &playersPosition[0])
	fmt.Sscanf(input[1], "Player 2 starting position: %d", &playersPosition[1])
	playersPosition[0]--
	playersPosition[1]--

	fmt.Println("2021 Day 21")
	fmt.Println("\tPart 1:", playGame(playersPosition))
	fmt.Println("\tPart 2:", playGameHarder(playersPosition))
}

func playGame(positions []int) int {
	playersPoints := []int{0, 0}
	playersPosition := utils.CopyIntSlice(positions)
	playerTurn := 0
	dieValue := 0
	for playersPoints[0] < 1000 && playersPoints[1] < 1000 {
		playersPosition[playerTurn] = (playersPosition[playerTurn] + 3*dieValue + 6) % 10
		playersPoints[playerTurn] += playersPosition[playerTurn] + 1
		playerTurn = (playerTurn + 1) % 2
		dieValue += 3
	}
	return playersPoints[playerTurn] * dieValue
}

func playGameHarder(positions []int) uint64 {
	// Combinations of ways to reach the sum of the 3 die
	universesPerValue := map[int]uint64{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}
	universes := make(map[string]Universe)
	wonBy := []uint64{0, 0}
	playerTurn := 0

	firstUniverse := Universe{
		playerPosition: positions,
		playerPoints:   []int{0, 0},
		amount:         1,
	}
	universes[getKey(firstUniverse)] = firstUniverse

	for len(universes) > 0 {
		newUniverses := make(map[string]Universe)

		for _, universe := range universes {
			for diceValue, universesCreated := range universesPerValue {
				u := copy(universe)
				u.playerPosition[playerTurn] = (u.playerPosition[playerTurn] + diceValue) % 10
				u.playerPoints[playerTurn] += u.playerPosition[playerTurn] + 1

				if u.playerPoints[playerTurn] >= 21 {
					wonBy[playerTurn] += u.amount * universesCreated
				} else {
					if _, ok := newUniverses[getKey(u)]; ok {
						u.amount = newUniverses[getKey(u)].amount + u.amount*universesCreated
					} else {
						u.amount *= universesCreated
					}

					newUniverses[getKey(u)] = u
				}
			}
		}

		universes = newUniverses
		playerTurn = (playerTurn + 1) % 2
	}

	if wonBy[0] > wonBy[1] {
		return wonBy[0]
	}
	return wonBy[1]
}

func getKey(universe Universe) string {
	return fmt.Sprintf("%d %d %d %d", universe.playerPosition[0], universe.playerPosition[1], universe.playerPoints[0], universe.playerPoints[1])
}

func copy(universe Universe) Universe {
	return Universe{
		playerPosition: utils.CopyIntSlice(universe.playerPosition),
		playerPoints:   utils.CopyIntSlice(universe.playerPoints),
		amount:         universe.amount,
	}
}
