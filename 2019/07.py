from itertools import permutations
from copy import deepcopy
from aocd import get_data

amplifier = [list(map(int, get_data(day=7, year=2019).split(","))), 0]


def run_amplifier(amplifier, *inputs):
    program = amplifier[0]
    pointer = amplifier[1]
    input_index = 0
    while program[pointer] != 99:
        instruction = program[pointer] % 100
        param1 = pointer + 1 if program[pointer] // 100 % 10 else program[pointer + 1]
        param2 = pointer + 2 if program[pointer] // 1000 else program[pointer + 2]
        if pointer + 3 < len(program):
            param3 = program[pointer + 3]

        if instruction == 1:
            program[param3] = program[param1] + program[param2]
            pointer += 4
        elif instruction == 2:
            program[param3] = program[param1] * program[param2]
            pointer += 4
        elif instruction == 3:
            program[param1] = inputs[input_index]
            input_index += 1
            pointer += 2
        elif instruction == 4:
            amplifier[1] = pointer + 2
            return True, program[param1]
        elif instruction == 5:
            if program[param1]:
                pointer = program[param2]
            else:
                pointer += 3
        elif instruction == 6:
            if not program[param1]:
                pointer = program[param2]
            else:
                pointer += 3
        elif instruction == 7:
            if program[param1] < program[param2]:
                program[param3] = 1
            else:
                program[param3] = 0
            pointer += 4
        elif instruction == 8:
            if program[param1] == program[param2]:
                program[param3] = 1
            else:
                program[param3] = 0
            pointer += 4
    return False, inputs[input_index]


maximum1 = 0
for permutation in list(permutations([0, 1, 2, 3, 4])):
    output = 0
    for phase in permutation:
        boolean, output = run_amplifier(deepcopy(amplifier), phase, output)
    maximum1 = max(maximum1, output)

maximum2 = 0
for permutation in list(permutations([5, 6, 7, 8, 9])):
    amplifiers = [
        deepcopy(amplifier),
        deepcopy(amplifier),
        deepcopy(amplifier),
        deepcopy(amplifier),
        deepcopy(amplifier),
    ]
    output = 0
    for phase in range(5):
        amplifier_running, output = run_amplifier(
            amplifiers[phase], permutation[phase], output
        )

    while amplifier_running:
        for phase in range(5):
            amplifier_running, output = run_amplifier(amplifiers[phase], output)
        maximum2 = max(maximum2, output)


print("2019 Day 07")
print("\tPart 1:", maximum1)
print("\tPart 2:", maximum2)
