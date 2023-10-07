from aocd import get_data

data = get_data(day=3, year=2020).splitlines()
height, width = len(data), len(data[0])


def find_trees(slope_x: int, slope_y: int):
    trees = current_x = current_y = 0

    while current_y < height:
        if data[current_y][current_x] == "#":
            trees += 1
        current_y += slope_y
        current_x = (current_x + slope_x) % width
    return trees


print("2020 Day 03")
print("\tPart 1:", find_trees(3, 1))
print(
    "\tPart 2:",
    find_trees(1, 1)
    * find_trees(3, 1)
    * find_trees(5, 1)
    * find_trees(7, 1)
    * find_trees(1, 2),
)
