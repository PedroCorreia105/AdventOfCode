package utils

import (
	"strconv"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}

func BinToInt(str string) int {
	number, err := strconv.ParseInt(str, 2, 64)

	if err != nil {
		panic(err)
	}

	return int(number)
}

func StringsToInts(array []string) []int {
	var result = []int{}

	for _, i := range array {
		result = append(result, StringToInt(i))
	}

	return result
}
