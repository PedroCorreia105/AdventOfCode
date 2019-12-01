from aocd import get_data
from intcode import run_program

program = list(map(int, get_data(day=13, year=2019).split(",")))
character = {0: " ", 1: "#", 2: "+", 3: "^", 4: "o"}
bar_x = ball_x = score = 0
screen, outputs = [], []


def add_pixel(x, y, value):
    global screen, ball_x, bar_x
    if y >= len(screen):
        screen += [[" "]] * (y - len(screen) + 1)
    if x >= len(screen[y]):
        screen[y] += [" "] * (x - len(screen[y]) + 1)
    screen[y][x] = character[value]
    if value == 4:
        ball_x = x
    elif value == 3:
        bar_x = x


def process_input(_):
    global ball_x, bar_x
    # print("\n".join(["".join(line) for line in screen]))
    return 1 if ball_x > bar_x else (-1 if ball_x < bar_x else 0)


def process_output(output, _):
    global outputs, score
    outputs.append(output)
    if len(outputs) == 3 and outputs[0] == -1 and outputs[1] == 0:
        score = outputs[2]
        outputs = []
    elif len(outputs) == 3:
        add_pixel(outputs[0], outputs[1], outputs[2])
        outputs = []


print("2019 Day 13")
run_program(program, process_input, process_output)
print("\tPart 1:", sum([line.count(character[2]) for line in screen]))

program[0] = 2
run_program(program, process_input, process_output)
print("\tPart 2:", score)
