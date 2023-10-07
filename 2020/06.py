from aocd import get_data

data = get_data(day=6, year=2020).split("\n\n")

sum1 = 0
for group in data:
    set1 = set(group)
    for individual in group.split("\n"):
        set1 = set1.intersection(individual)
    sum1 += len(set1)


print("2020 Day 06")
print("\tPart 1:", sum(len(set(line.replace("\n", ""))) for line in data))
print("\tPart 2:", sum1)
