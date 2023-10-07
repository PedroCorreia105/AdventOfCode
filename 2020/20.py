from re import finditer
from aocd import get_data

data = get_data(day=20, year=2020).split("\n\n")
tiles = {}


def flip_rotate(field: list, times: int):
    # Flip
    if times > 3:
        field = field[::-1]

    # Rotate
    for _ in range(times % 4):
        field = [
            "".join([line[index] for line in field][::-1])
            for index in range(len(field))
        ]
    return field


class Tile:
    def __init__(self, number, field):
        self.number = number
        self.field = field
        self.conections = set()
        self.configuration = []
        self.x = 0
        self.y = 0
        north_edge = field[0]
        east_edge = "".join(line[-1] for line in field)
        south_edge = field[-1]
        west_edge = "".join(line[0] for line in field)

        self.configurations = [
            [north_edge, east_edge, south_edge, west_edge],
            [west_edge[::-1], north_edge, east_edge[::-1], south_edge],
            [south_edge[::-1], west_edge[::-1], north_edge[::-1], east_edge[::-1]],
            [east_edge, south_edge[::-1], west_edge, north_edge[::-1]],
            [south_edge, east_edge[::-1], north_edge, west_edge[::-1]],
            [west_edge, south_edge, east_edge, north_edge],
            [north_edge[::-1], west_edge, south_edge[::-1], east_edge],
            [east_edge[::-1], north_edge[::-1], west_edge[::-1], south_edge[::-1]],
        ]

    def change_field(self, index):
        self.configuration = self.configurations[index]

        self.field = flip_rotate(self.field, index)

        # Remove edges
        self.field = [line[1:-1] for line in self.field[1:-1]]


# Create Tiles
for tile in data:
    tiles_list = tile.splitlines()
    tile_number = tiles_list[0][5:-1]
    tiles[tile_number] = Tile(tile_number, tiles_list[1:])

# Find conections
for tile1 in tiles.values():
    for configuration1 in tile1.configurations:
        for edge_position in range(4):
            for tile2 in tiles.values():
                if tile1 != tile2:
                    for configuration2 in tile2.configurations:
                        if (
                            configuration1[edge_position]
                            == configuration2[(edge_position + 2) % 4]
                        ):
                            tile2.conections.add(tile1)
                            tile1.conections.add(tile2)

# Multiplicate corner values
multiplication = 1
for tile in tiles.values():
    if len(tile.conections) == 2:
        multiplication *= int(tile.number)
        corner = tile


print("2020 Day 20")
print("\tPart 1:", multiplication)


def match(tile_1: Tile, tile_2: Tile):
    for configuration_index in range(8):
        for edge in range(4):
            if (
                tile_1.configuration[edge]
                == tile_2.configurations[configuration_index][(edge + 2) % 4]
            ):
                vector = [(0, -1), (1, 0), (0, 1), (-1, 0)][edge]
                tile_2.change_field(configuration_index)
                return tile_1.x + vector[0], tile_1.y + vector[1]
    return False


def determine_coordinates(tile_1: Tile):
    for tile_2 in tile_1.conections:
        if tile_2 != corner and tile_2.x == 0 and tile_2.y == 0:
            tile_2.x, tile_2.y = match(tile_1, tile_2)
            determine_coordinates(tile_2)


corner.change_field(0)
determine_coordinates(corner)

small_x = small_y = big_x = big_y = 0
tile_coordinates = {}
for tile in tiles.values():
    small_x, small_y = min(tile.x, small_x), min(tile.y, small_y)
    big_x, big_y = max(tile.x, big_x), max(tile.y, big_y)
    tile_coordinates[repr((tile.x, tile.y))] = tile

final_image = []
for y in range(small_y, big_y + 1):
    for line in range(len(corner.field)):
        final_image.append(
            "".join(
                tile_coordinates[repr((x, y))].field[line]
                for x in range(small_x, big_x + 1)
            )
        )

# Find the correct final image
for times in range(8):
    final_image = flip_rotate(final_image, times)

    # Find number of monsters
    monsters = 0
    for line in range(len((final_image))):
        ns1 = [
            match.start(0)
            for match in finditer(r"(?=#....##....##....###)", final_image[line])
        ]
        for n1 in ns1:
            if n1 and line > 0:
                ns2 = [
                    match.start(0)
                    for match in finditer(
                        r"(?=.#..#..#..#..#..#)", final_image[line + 1]
                    )
                ]
                if n1 in ns2:
                    if final_image[line - 1][n1 + 18] == "#":
                        monsters += 1
    if monsters:
        print("\tPart 2:", sum(line.count("#") for line in final_image) - 15 * monsters)
        break
