from aocd import get_data

data = list(map(int, get_data(day=1, year=2019).splitlines()))

total1 = 0
for module in data:
    total1 += module // 3 - 2

total2 = 0
for module in data:
    fuel = module
    while fuel > 8:
        fuel = fuel // 3 - 2
        total2 += fuel

print("2019 Day 01")
print("\tPart 1:", total1)
print("\tPart 2:", total2)
