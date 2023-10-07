from ast import literal_eval
from copy import deepcopy
from aocd import get_data

data = get_data(day=24, year=2020).splitlines()
vectors = {
    "e": (2, 0),
    "ne": (1, -1),
    "nw": (-1, -1),
    "w": (-2, 0),
    "sw": (-1, 1),
    "se": (1, 1),
}
floor = {}

for line in data:
    direction = ""
    current_location = [0, 0]
    for char in line:
        if direction + char in vectors:
            vector = vectors[direction + char]
            current_location[0] += vector[0]
            current_location[1] += vector[1]
            direction = ""
            if repr(current_location) not in floor:
                floor[repr(current_location)] = False
        else:
            direction += char

    floor[repr(current_location)] = not floor[repr(current_location)]


print("2020 Day 24")
print("\tPart 1:", sum(value for value in floor.values()))

for day in range(101):
    floor_copy = deepcopy(floor)
    for tile in floor:
        x, y = literal_eval(tile)
        adjacent_coordinates = [
            repr([x + vector[0], y + vector[1]]) for vector in vectors.values()
        ]

        adjacents_tiles = []
        for coordinate in adjacent_coordinates:
            if coordinate not in floor:
                floor_copy[coordinate] = False
            else:
                adjacents_tiles.append(floor[coordinate])

        if day > 0 and floor[tile] and not 0 < adjacents_tiles.count(True) < 3:
            floor_copy[tile] = False
        elif day > 0 and (not floor[tile]) and adjacents_tiles.count(True) == 2:
            floor_copy[tile] = True

    floor = floor_copy


print("\tPart 2:", sum(value for value in floor.values()))
