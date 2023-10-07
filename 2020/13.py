from aocd import get_data

data = get_data(day=13, year=2020).splitlines()
time = int(data[0])
buses = data[1].split(",")
buses = [(int(bus), buses.index(bus)) for bus in buses if bus != "x"]

# Part 1
minimum = min([(bus - time % bus, bus) for bus, index in buses])

# Part 2
counter = solution = 1
for pair in buses:
    bus = pair[0]
    index = bus - pair[1]
    while solution % bus != index % bus:
        solution += counter
    counter = counter * bus


print("2020 Day 13")
print("\tPart 1:", minimum[0] * minimum[1])
print("\tPart 2:", solution)
