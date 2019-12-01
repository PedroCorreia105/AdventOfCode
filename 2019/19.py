from aocd import get_data
from intcode import run_program

program = list(map(int, get_data(day=19, year=2019).split(",")))
total_pulled = counter = x_turn = corners_index = previous_output = found_beam = 0
corners = coordinate = [500, 0]
answer_not_found = True
size = 50


def process_input1(_):
    global counter, x_turn
    x_turn = not x_turn
    return counter // size if x_turn else counter % size


def process_output1(output, _):
    global total_pulled, counter
    counter += 1
    total_pulled += output


while counter < size * size:
    run_program(program, process_input1, process_output1)

print("2019 Day 19")
print("\tPart 1:", total_pulled)


def calculate_oposite_corner(coordinate):
    return [coordinate[0] + 99, coordinate[1] - 99] + coordinate


def process_input2(_):
    global corners, corners_index
    corners_index += 1
    return corners[corners_index - 1]


def process_output2(output, _):
    global previous_output, corners_index, coordinate, corners, answer_not_found, found_beam

    if corners_index == len(corners):
        if not found_beam and output == 0:
            corners = coordinate = [coordinate[0], coordinate[1] + 1]
        elif not found_beam and output == 1:
            found_beam = True
            corners = coordinate = [coordinate[0], coordinate[1] + 1]
        elif found_beam and output == 0:
            coordinate = [coordinate[0] + 1, coordinate[1] - 2]
            corners = calculate_oposite_corner(coordinate)
        elif found_beam and previous_output == 0 and output == 1:
            coordinate = [coordinate[0], coordinate[1] + 1]
            corners = calculate_oposite_corner(coordinate)
        elif found_beam and output == previous_output == 1:
            answer_not_found = False
        corners_index = previous_output = 0
    else:
        previous_output = output


while answer_not_found:
    run_program(program, process_input2, process_output2)

print("\tPart 2:", coordinate[0] * 10000 + coordinate[1] - 99)
