from copy import deepcopy
from itertools import product
from aocd import get_data

data = list(map(list, get_data(day=17, year=2020).splitlines()))
ACTIVE, INACTIVE = "#", "."

def calc_surroundings1(x, y, z):
    vectors = list(product([1, 0, -1], repeat=3))
    vectors.remove((0, 0, 0))
    surroundings = []
    for vector in vectors:
        surrounding = repr((x + vector[0], y + vector[1], z + vector[2]))
        if surrounding not in space:
            space[surrounding] = INACTIVE
        surroundings.append(space[surrounding])
    return surroundings

def calc_surroundings2(x, y, z, w):
    vectors = list(product([1, 0, -1], repeat=4))
    vectors.remove((0, 0, 0, 0))
    surroundings = []
    for vector in vectors:
        surrounding = repr((x + vector[0], y + vector[1], z + vector[2], w + vector[3]))
        if surrounding not in space:
            space[surrounding] = INACTIVE
        surroundings.append(space[surrounding])
    return surroundings

size = len(data)
space = {
    repr((x - (size // 2), y - (size // 2), 0)): data[y][x]
    for x in range(size)
    for y in range(size)
}
space_copy = deepcopy(space)

for _ in range(6):
    size += 2
    for z in range(-(size // 2), size // 2 + 1):
        for y in range(-(size // 2), size // 2 + 1):
            for x in range(-(size // 2), size // 2 + 1):
                cube = repr((x, y, z))
                if cube not in space:
                    space[cube] = INACTIVE
                surroundings = calc_surroundings1(x, y, z)
                if space[cube] == ACTIVE and surroundings.count(ACTIVE) not in [
                    2,
                    3,
                ]:
                    space_copy[cube] = INACTIVE
                elif space[cube] == INACTIVE and surroundings.count(ACTIVE) == 3:
                    space_copy[cube] = ACTIVE
    space = deepcopy(space_copy)

print("2020 Day 17")
print("\tPart 1:", list(space.values()).count(ACTIVE))

size = len(data)
space = {
    repr((x - (size // 2), y - (size // 2), 0, 0)): data[y][x]
    for x in range(size)
    for y in range(size)
}
space_copy = deepcopy(space)

for _ in range(6):
    size += 2
    for w in range(-(size // 2), size // 2 + 1):
        for z in range(-(size // 2), size // 2 + 1):
            for y in range(-(size // 2), size // 2 + 1):
                for x in range(-(size // 2), size // 2 + 1):
                    cube = repr((x, y, z, w))
                    if cube not in space:
                        space[cube] = INACTIVE
                    surroundings = calc_surroundings2(x, y, z, w)
                    if space[cube] == ACTIVE and surroundings.count(ACTIVE) not in [
                        2,
                        3,
                    ]:
                        space_copy[cube] = INACTIVE
                    elif space[cube] == INACTIVE and surroundings.count(ACTIVE) == 3:
                        space_copy[cube] = ACTIVE
    space = deepcopy(space_copy)


print("\tPart 2:", list(space.values()).count(ACTIVE))
