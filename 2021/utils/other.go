package utils

import (
	"sort"
	"strings"
)

func Sort(string1 string) string {
	slice := strings.Split(string1, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func NumberOfCommonCharacters(string1, string2 string) int {
	map1 := make([]int, 26)
	map2 := make([]int, 26)
	common := 0

	for _, c1 := range string1 {
		map1[c1-97]++
	}

	for _, c2 := range string2 {
		map2[c2-97]++
	}

	for i := 0; i < 26; i++ {
		if map1[i] > 0 && map2[i] > 0 {
			common++
		}
	}

	return common
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
