package utils

func ContainsInt(slice []int, element int) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}

func CopyIntSlice(slice []int) []int {
	copy := make([]int, len(slice))
	for i := range slice {
		copy[i] = slice[i]
	}
	return copy
}

func CopyStringSlice(slice []string) []string {
	copy := make([]string, len(slice))
	for i := range slice {
		copy[i] = slice[i]
	}
	return copy
}

func IndexOfInt(slice []int, element int) int {
	for index, i := range slice {
		if i == element {
			return index
		}
	}
	return -1
}

func SumIntSlice(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

func InsertIfLarger(slice []int, element int) {
	next := element
	for i, value := range slice {
		if next > value {
			slice[i], next = next, value
		}
	}
}

func IntToSlice(num int) []int {
	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append(digits, digit)
		num /= 10
	}
	// Reverse the digits slice
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

func SliceToInt(digits []int) int {
	num := 0
	for _, digit := range digits {
		num = num*10 + digit
	}
	return num
}
