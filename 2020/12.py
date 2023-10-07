from aocd import get_data

data = list(map(lambda e: (e[0], int(e[1:])), get_data(day=12, year=2020).splitlines()))
directions = {"N": (0, 1), "S": (0, -1), "W": (-1, 0), "E": (1, 0)}
angles = {90: (0, 1), 270: (0, -1), 180: (-1, 0), 0: (1, 0)}
facing = 0
coordinates = [0, 0]

for direction in data:
    if direction[0] in directions:
        coordinates[0] += directions[direction[0]][0] * direction[1]
        coordinates[1] += directions[direction[0]][1] * direction[1]
    elif direction[0] == "F":
        coordinates[0] += angles[facing][0] * direction[1]
        coordinates[1] += angles[facing][1] * direction[1]
    elif direction[0] == "L":
        facing = (facing + direction[1]) % 360
    elif direction[0] == "R":
        facing = (facing - direction[1]) % 360

print("2020 Day 12")
print("\tPart 1:", abs(coordinates[0]) + abs(coordinates[1]))

coordinates = [0, 0]
waypoint = [10, 1]

for direction in data:
    if direction[0] in directions:
        waypoint[0] += directions[direction[0]][0] * direction[1]
        waypoint[1] += directions[direction[0]][1] * direction[1]
    elif direction[0] == "F":
        coordinates[0] += waypoint[0] * direction[1]
        coordinates[1] += waypoint[1] * direction[1]
    else:
        angle = direction[1] * (1 if direction[0] == "L" else -1) % 360
        while angle > 0:
            waypoint = [-waypoint[1], waypoint[0]]
            angle -= 90

print("\tPart 2:", abs(coordinates[0]) + abs(coordinates[1]))
