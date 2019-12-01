from aocd import get_data

data = get_data(day=3, year=2019).splitlines()


def calculate_path(directions):
    path = []
    location = [0, 0]
    for direction in directions:
        for _ in range(int(direction[1:])):
            if direction[0] == "R":
                location[0] += 1
            elif direction[0] == "L":
                location[0] -= 1
            elif direction[0] == "U":
                location[1] += 1
            elif direction[0] == "D":
                location[1] -= 1
            path.append(tuple(location))
    return path


wire_path1 = calculate_path(data[0].split(","))
wire_path2 = calculate_path(data[1].split(","))
intersections = set(wire_path1).intersection(wire_path2)


print("2019 Day 03")
print("\tPart 1:", min(map(lambda e: abs(e[0]) + abs(e[1]), intersections)))
print(
    "\tPart 2:",
    min(map(lambda e: 2 + wire_path1.index(e) + wire_path2.index(e), intersections)),
)
