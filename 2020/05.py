from aocd import get_data

data = get_data(day=5, year=2020).splitlines()
seats = []

for code in data:
    ranges = [0, 127]
    for letter in code[:7]:
        if letter == "F":
            ranges[1] = ranges[0] + (ranges[1] - ranges[0]) // 2
        else:
            ranges[0] = ranges[0] + (ranges[1] - ranges[0]) // 2 + 1
    row = ranges[0]

    ranges = [0, 7]
    for letter in code[7:]:
        if letter == "L":
            ranges[1] = ranges[0] + (ranges[1] - ranges[0]) // 2
        else:
            ranges[0] = ranges[0] + (ranges[1] - ranges[0]) // 2 + 1
    column = ranges[0]

    seats += [row * 8 + column]

print("2020 Day 05")
print("\tPart 1:", max(seats))

for i in seats:
    if i + 1 not in seats and i + 2 in seats:
        print("\tPart 2:", i + 1)
