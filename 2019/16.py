from aocd import get_data

data = get_data(day=16, year=2019)
digit_list = list(map(int, data))
starting_pattern = [0, 1, 0, -1]

for _ in range(100):
	digit_list_copy = []
	for line in range(len(digit_list)):
		pattern = []
		for number in starting_pattern:
			pattern += (line+1) * [number]

		total = 0
		for index, value in enumerate(digit_list):
			total += value * pattern[(index+1) % len(pattern)]
		digit_list_copy.append(abs(total)%10)
	digit_list = digit_list_copy


print("2019 Day 16")
print("\tPart 1:", "".join(map(str, digit_list[:8])))

data = data * 10000
digit_list = map(int, data[int(data[:7]):][::-1])

for _ in range(100):
	digit_list_copy = []
	total = 0
	for digit in digit_list:
		total += digit
		digit_list_copy.append(total % 10)
	digit_list = digit_list_copy

print("\tPart 2:", "".join(map(str,digit_list[::-1][:8])))