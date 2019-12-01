from aocd import get_data
from copy import deepcopy

data = get_data(day=20, year=2019).splitlines()
vectors = [(0, 1), (0, -1), (1, 0), (-1, 0)]
tiles, adjacency_matrix = {}, {}


class Tile:
    def __init__(self, coordinate: tuple):
        self.coordinate = coordinate
        self.neighbours = []
        self.outside = self.portal = self.visited = False
        self.portal_code = ""
        self.distance = 0


class State:
    def __init__(self, portal_code: str, level: int, outside: bool, distance: int):
        self.portal_code = portal_code
        self.level = level
        self.outside = outside
        self.distance = distance

    def __repr__(self):
        return self.portal_code + repr(self.level) + repr(self.outside)


# Parse input into the tiles list
for y in range(2, len(data) - 2):
    for x in range(2, len(data[0]) - 2):
        if data[y][x] == ".":
            tile = Tile((x, y))
            tiles[repr((x, y))] = tile

            for v in vectors:
                neighbour = (x + v[0], y + v[1])
                if repr(neighbour) in tiles:
                    tile.neighbours.append(tiles[repr(neighbour)])
                    tiles[repr(neighbour)].neighbours.append(tile)
                if data[neighbour[1]][neighbour[0]].isalpha():
                    tile.portal = True
                    a, b = (
                        data[neighbour[1]][neighbour[0]],
                        data[neighbour[1] + v[1]][neighbour[0] + v[0]],
                    )
                    tile.portal_code = a + b if a < b else b + a
                    tile.outside = neighbour[0] in [1, len(data[0]) - 2] or neighbour[
                        1
                    ] in [1, len(data) - 2]

# Small bfs to link portals and store distances in adjacency_matrix
# Its not necessary but improves performance
for tile1 in tiles.values():
    if tile1.portal:
        if tile1.portal_code not in adjacency_matrix:
            adjacency_matrix[tile1.portal_code] = {}
        to_visit = [tile1]

        while to_visit != []:
            tile2 = to_visit.pop(0)

            for neighbour in tile2.neighbours:
                if not neighbour.visited:
                    to_visit.append(neighbour)
                    neighbour.distance = tile2.distance + 1

            tile2.visited = True
            if tile2.portal and tile1 != tile2:
                adjacency_matrix[tile1.portal_code][tile2.portal_code] = (
                    tile2.distance,
                    tile1.outside,
                    tile2.outside,
                )

        for tile in tiles.values():
            tile.visited = False
            tile.distance = 0

initial_state = State("AA", 0, True, 0)
to_visit = [initial_state]
cache = [repr(initial_state)]


# Bfs between portals in several levels
def calculate(part: int, to_visit):
    to_visit = deepcopy(to_visit)
    while to_visit != []:
        state = to_visit.pop(0)

        for portal_code2 in adjacency_matrix[state.portal_code]:
            distance, source_outside, destiny_outside = adjacency_matrix[
                state.portal_code
            ][portal_code2]

            if (
                portal_code2 == "ZZ"
                and state.level == 0
                and source_outside == state.outside
            ):
                return state.distance + distance

            if part == 1:
                level_change = 0
            else:
                level_change = -1 if destiny_outside else 1

            if 0 <= state.level + level_change and source_outside == state.outside:
                state2 = State(
                    portal_code2,
                    state.level + level_change,
                    not destiny_outside,
                    state.distance + distance + 1,
                )
                if repr(state2) not in cache:
                    to_visit.append(state2)
                    cache.append(repr(state2))


print("2019 Day 20")
print("\tPart 1:", calculate(1, to_visit))
print("\tPart 2:", calculate(2, to_visit))
