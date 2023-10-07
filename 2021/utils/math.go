package utils

func Max(nums ...int) int {
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(nums ...int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

func SortNumbers(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func AbsDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Mod(a, b int) int {
	return (a%b + b) % b
}
