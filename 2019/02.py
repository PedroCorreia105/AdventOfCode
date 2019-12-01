from copy import deepcopy
from aocd import get_data

data = list(map(int, get_data(day=2, year=2019).split(",")))


def run_program(program: list, noun: int, verb: int):
    program[1] = noun
    program[2] = verb
    pointer = 0
    while program[pointer] != 99:
        if program[pointer] == 1:
            program[program[pointer + 3]] = (
                program[program[pointer + 1]] + program[program[pointer + 2]]
            )
        elif program[pointer] == 2:
            program[program[pointer + 3]] = (
                program[program[pointer + 1]] * program[program[pointer + 2]]
            )
        pointer += 4
    return program[0]


print("2019 Day 02")
print("\tPart 1:", run_program(deepcopy(data), 12, 2))

for n in range(100):
    for v in range(100):
        if run_program(deepcopy(data), n, v) == 19690720:
            print("\tPart 2:", n * 100 + v)
