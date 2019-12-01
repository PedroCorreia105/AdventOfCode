from aocd import get_data
from intcode import run_program

program = list(map(int, get_data(day=17, year=2019).split(",")))
input_index = dust = alignment_total = facing_direction = 0
vectors = [(0, -1), (1, 0), (0, 1), (-1, 0)]
area = ""
path = []


def process_input(_):
    global input_index
    value = input_list[input_index]
    input_index += 1
    return value


def process_output(output, _):
    global area, dust
    dust = output
    area += chr(output)


run_program(program, process_input, process_output)

# Calculate alignement total and current location
area = area.split("\n")[:-2]
for y in range(1, len(area) - 1):
    for x in range(1, len(area[0]) - 1):
        if (
            area[y][x]
            == area[y + 1][x]
            == area[y - 1][x]
            == area[y][x + 1]
            == area[y][x - 1]
            == "#"
        ):
            alignment_total += x * y
        elif area[y][x] == "^":
            current_location = (x, y)


print("2019 Day 17")
print("\tPart 1:", alignment_total)


def path_ahead(current_location, direction):
    y = current_location[1] + vectors[direction][1]
    x = current_location[0] + vectors[direction][0]
    return -1 < x < len(area[0]) and -1 < y < len(area) and area[y][x] == "#"


# Construct the path list by going through the # path
while True:
    if path_ahead(current_location, facing_direction):
        if len(path) > 0 and isinstance(path[-1], int):
            path[-1] += 1
        else:
            path.append(1)
        current_location = (
            current_location[0] + vectors[facing_direction][0],
            current_location[1] + vectors[facing_direction][1],
        )
    elif path_ahead(current_location, (facing_direction + 1) % 4):
        facing_direction = (facing_direction + 1) % 4
        path.append("R")
    elif path_ahead(current_location, (facing_direction - 1) % 4):
        facing_direction = (facing_direction - 1) % 4
        path.append("L")
    else:
        break


# Brute force dividing the path into 3 repeatable sections
def calculate():
    for A1 in range(2, 12, 2):
        for B0 in range(A1, len(path) - 4, 2):
            for B1 in range(B0 + 2, B0 + 10, 2):
                for C0 in range(B1, len(path) - 2, 2):
                    for C1 in range(C0 + 2, C0 + 10, 2):
                        path_withoutA = repr(path).replace(repr(path[:A1])[1:-1], "A")
                        path_withoutAB = path_withoutA.replace(
                            repr(path[B0:B1])[1:-1], "B"
                        )
                        path_withoutABC = path_withoutAB.replace(
                            repr(path[C0:C1])[1:-1], "C"
                        )

                        if all(i in "ABC, " for i in path_withoutABC[1:-1]):
                            return (
                                path_withoutABC[1:-1],
                                path[:A1],
                                path[B0:B1],
                                path[C0:C1],
                            )


main, A, B, C = calculate()
main = [ord(x) for x in main.replace(" ", "")] + [ord("\n")]
A = [ord(x) for x in ",".join(str(i) for i in A)] + [ord("\n")]
B = [ord(x) for x in ",".join(str(i) for i in B)] + [ord("\n")]
C = [ord(x) for x in ",".join(str(i) for i in C)] + [ord("\n"), ord("n"), ord("\n")]
input_list = main + A + B + C

program[0] = 2
run_program(program, process_input, process_output)
print("\tPart 2:", dust)
