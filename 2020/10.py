from aocd import get_data

data = sorted(list(map(int, get_data(day=10, year=2020).splitlines())))
data.append(data[-1] + 3)
data.insert(0, 0)
ones = threes = 0

for i in range(1, len(data)):
    if data[i] - data[i - 1] == 1:
        ones += 1
    elif data[i] - data[i - 1] == 3:
        threes += 1


# Find number of possible ways to go through the interval, including first and last position
# [0, 1] = 1
# [0, 1, 2] = 2 = [0, 2] + [0, 1, 2]
# [0, 1, 2, 3] = 4 = [0, 1, 2, 3] + [0, 1, 3] + [0, 2, 3] + [0, 3]
# [0, 1, 2, 4] = 3 = [0, 1, 4] + [0, 2, 4] + [0, 1, 2, 4]
def ways(interval: list):
    if len(interval) == 2:
        return 1
    if interval[2] - interval[0] <= 3:
        return ways(interval[1:]) + ways([interval[0]] + interval[2:])
    return ways(interval[1:])


total = 1
mandatory = 0
for i in range(1, len(data) - 1):
    # Find mandatory numbers that must be in all permutations
    if data[i + 1] - data[i - 1] > 3:
        # Calculate ways to travel between mandatory numbers and multiply possibilities
        total *= ways(data[mandatory : i + 1])
        mandatory = i


print("2020 Day 10")
print("\tPart 1:", ones * threes)
print("\tPart 2:", total)
