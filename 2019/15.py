from aocd import get_data
from intcode import run_program

program = list(map(int, get_data(day=15, year=2019).split(",")))
vectors = {1: (0, -1), 2: (0, 1), 3: (-1, 0), 4: (1, 0)}


class Tile:
    def __init__(self, coordinates, depth):
        self.coordinates = coordinates
        self.borders = [0, 0, 0, 0]
        self.wall = False
        self.destination = False
        self.visited = 0
        self.depth = depth
        self.depth2 = -1

    def is_destination(self):
        self.destination = True
        self.depth2 = 0

    def get_best_direction(self):
        for index, value in enumerate(self.borders):
            if value == 0:
                return index + 1
        return (
            self.borders.index(
                min(
                    [tile for tile in self.borders if not tile.wall],
                    key=lambda x: x.visited,
                )
            )
            + 1
        )

    def set_border(self, direction, value):
        if self.borders[direction - 1] == 0:
            self.borders[direction - 1] = value


next_tile = current_tile = Tile((0, 0), 0)
area = {repr((0, 0)): current_tile}


def process_input(_):
    global current_tile, next_tile
    direction = current_tile.get_best_direction()
    new_coordinates = (
        current_tile.coordinates[0] + vectors[direction][0],
        current_tile.coordinates[1] + vectors[direction][1],
    )

    if repr(new_coordinates) not in area:
        next_tile = Tile(new_coordinates, current_tile.depth + 1)
        area[repr(new_coordinates)] = next_tile
    else:
        next_tile = area[repr(new_coordinates)]

    if current_tile.depth2 >= 0:
        if next_tile.depth2 < 0:
            next_tile.depth2 = current_tile.depth2 + 1
        else:
            next_tile.depth2 = min(next_tile.depth2, current_tile.depth2 + 1)

    current_tile.set_border(direction, next_tile)
    current_tile.visited += 1
    return direction


def process_output(result, _):
    global current_tile, next_tile
    if result == 0:
        next_tile.wall = True
    elif result == 1:
        current_tile = next_tile
    else:
        current_tile = next_tile
        current_tile.is_destination()


def stoping_condition():
    return any(not tile.wall and tile.depth2 == -1 for tile in area.values())


run_program(program, process_input, process_output, stoping_condition)


print("2019 Day 15")
print("\tPart 1:", [tile.depth for tile in area.values() if tile.destination][0])
print("\tPart 2:", max(tile.depth2 for tile in area.values()))
