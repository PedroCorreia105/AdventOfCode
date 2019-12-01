from aocd import get_data

data = get_data(day=24, year=2019).splitlines()
height, width = len(data), len(data[0])
vectors = [(0, -1), (1, 0), (0, 1), (-1, 0)]
BUG, EMPTY = "#", "."


# Returns the value in the position x, y in the last iteration area
def get_value_in_position(x: int, y: int):
    global areas
    return areas[-1][y * width + x]


# Returns true if the last area is duplicated
def is_dublicated(areas):
    if len(areas) < 2:
        return False

    return areas[-1] in areas[:-1]


areas = ["".join(data)]

while not is_dublicated(areas):
    newArea = ["" for _ in range(height)]
    for y in range(height):
        for x in range(width):
            bugs_in_the_neighbourhood = 0
            for v in vectors:
                neighbour = (x + v[0], y + v[1])
                if 0 <= neighbour[0] < width and 0 <= neighbour[1] < height:
                    if get_value_in_position(*neighbour) == BUG:
                        bugs_in_the_neighbourhood += 1

            current_position_value = get_value_in_position(x, y)

            if current_position_value == BUG and bugs_in_the_neighbourhood != 1:
                newArea[y] += EMPTY
            elif current_position_value == EMPTY and bugs_in_the_neighbourhood in [
                1,
                2,
            ]:
                newArea[y] += BUG
            else:
                newArea[y] += current_position_value
    areas += ["".join(newArea)]


total = 0
for y in range(height):
    for x in range(width):
        if areas[-1][y * width + x] == BUG:
            total += 2 ** (y * width + x)

print("2019 Day 24")
print("\tPart 1:", total)


def get_value_in_position(x: int, y: int, level: int):
    global areas
    if level not in areas:
        return EMPTY
    return areas[level][y][x]


def calculate_neighbours(x: int, y: int):
    neighbours = []
    for v in vectors:
        neighbour = (x + v[0], y + v[1])
        if (
            0 <= neighbour[0] < width
            and 0 <= neighbour[1] < height
            and neighbour != (2, 2)
        ):
            neighbours.append((*neighbour, 0))
    if x == 0:
        neighbours.append((1, 2, -1))
    elif x == 4:
        neighbours.append((3, 2, -1))
    if y == 0:
        neighbours.append((2, 1, -1))
    elif y == 4:
        neighbours.append((2, 3, -1))
    if (x, y) == (1, 2):
        for i in range(height):
            neighbours.append((0, i, 1))
    elif (x, y) == (2, 1):
        for i in range(height):
            neighbours.append((i, 0, 1))
    elif (x, y) == (2, 3):
        for i in range(height):
            neighbours.append((i, 4, 1))
    elif (x, y) == (3, 2):
        for i in range(height):
            neighbours.append((4, i, 1))
    return neighbours


# Precalculate the neighbour of each position
neighbours = []
for y in range(height):
    neighbours.append([])
    for x in range(width):
        neighbours[y].append(calculate_neighbours(x, y))

areas = {0: data}

for step in range(200):
    newAreas = {
        level: ["" for _ in range(height)] for level in range(-step - 1, step + 2)
    }
    for level in newAreas:
        for y in range(height):
            for x in range(width):
                bugs_in_the_neighbourhood = 0
                for neighbour in neighbours[y][x]:
                    neighbour_level = neighbour[2] + level
                    if neighbour_level in areas:
                        if (
                            get_value_in_position(
                                neighbour[0], neighbour[1], neighbour_level
                            )
                            == BUG
                        ):
                            bugs_in_the_neighbourhood += 1

                current_position_value = get_value_in_position(x, y, level)

                if (x, y) == (2, 2):
                    newAreas[level][y] += EMPTY
                elif current_position_value == BUG and bugs_in_the_neighbourhood != 1:
                    newAreas[level][y] += EMPTY
                elif current_position_value == EMPTY and bugs_in_the_neighbourhood in [
                    1,
                    2,
                ]:
                    newAreas[level][y] += BUG
                else:
                    newAreas[level][y] += current_position_value
    areas = newAreas

total = 0
for level in newAreas:
    for y in range(height):
        for x in range(width):
            if get_value_in_position(x, y, level) == BUG:
                total += 1

print("\tPart 2:", total)
