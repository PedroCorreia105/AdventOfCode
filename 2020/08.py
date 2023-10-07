from copy import deepcopy
from aocd import get_data

program = [
    [c[0], int(c[1])]
    for c in list(map(str.split, get_data(day=8, year=2020).splitlines()))
]


def run_program(program: list, mutation: int = -1):
    program = deepcopy(program)
    if mutation > -1:
        program[mutation][0] = "jmp" if program[mutation][0] == "nop" else "nop"
    acumulator = pointer = 0
    current_command = program[pointer]

    while len(current_command) == 2:
        if current_command[0] == "acc":
            acumulator += current_command[1]

        pointer += current_command[1] if current_command[0] == "jmp" else 1

        if pointer == len(program) - 1:
            return acumulator

        current_command.append("*")
        current_command = program[pointer]

    return acumulator if mutation == -1 else False


print("2020 Day 08")
print("\tPart 1:", run_program(program))

for mutation_line in range(len(program)):
    result = run_program(program, mutation_line)
    if result:
        print("\tPart 2:", result)
