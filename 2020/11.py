from copy import deepcopy
from aocd import get_data

data = list(map(list, get_data(day=11, year=2020).splitlines()))
vectors = ((0, 1), (0, -1), (1, -1), (1, 0), (1, 1), (-1, -1), (-1, 0), (-1, 1))


def calc_surroundings1(y: int, x: int, data):
    surroundings = []
    for vector in vectors:
        if -1 < y + vector[0] < len(data) and -1 < x + vector[1] < len(data[y]):
            surroundings.append(data[y + vector[0]][x + vector[1]])
    return surroundings


def calc_surroundings2(y: int, x: int, data):
    surroundings = []
    for vector in vectors:
        found = False
        multiplier = 1
        while not found:
            surrounding_x = x + multiplier * vector[1]
            surrounding_y = y + multiplier * vector[0]

            if -1 < surrounding_y < len(data) and -1 < surrounding_x < len(data[y]):
                if data[surrounding_y][surrounding_x] != ".":
                    surroundings.append(data[surrounding_y][surrounding_x])
                    found = True
                else:
                    multiplier += 1
            else:
                found = True
    return surroundings


def calculate(data, part: int):
    data_copy = deepcopy(data)
    change_made = True
    while change_made:
        change_made = False
        for line_index, line in enumerate(data):
            for seat_index, seat in enumerate(line):
                if part == 1:
                    surroundings = calc_surroundings1(line_index, seat_index, data)
                else:
                    surroundings = calc_surroundings2(line_index, seat_index, data)
                if seat == "L" and "#" not in surroundings:
                    data_copy[line_index][seat_index] = "#"
                    change_made = True
                elif seat == "#" and surroundings.count("#") >= 3 + part:
                    data_copy[line_index][seat_index] = "L"
                    change_made = True
        data = deepcopy(data_copy)
    return data_copy


print("2020 Day 11")
print("\tPart 1:", sum(line.count("#") for line in calculate(data, 1)))
print("\tPart 2:", sum(line.count("#") for line in calculate(data, 2)))
