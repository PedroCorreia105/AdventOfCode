from copy import deepcopy


def write(program, position, value):
    if position >= len(program):
        program += [0] * (position - len(program) + 1)
    program[position] = value


def read(program, position):
    if position >= len(program):
        program += [0] * (position - len(program) + 1)
    return program[position]


def get_parameter(program, pointer, parameter, relative_base):
    if read(program, pointer) // 10 ** (parameter + 1) % 10 == 1:
        return pointer + parameter
    if read(program, pointer) // 10 ** (parameter + 1) % 10 == 2:
        return read(program, pointer + parameter) + relative_base
    return read(program, pointer + parameter)


def run_program(program, process_input, process_output, condition=lambda: True, id=None):
    pointer = relative_base = 0
    program = deepcopy(program)

    while condition() and read(program, pointer) != 99:
        instruction = read(program, pointer) % 100
        param1 = get_parameter(program, pointer, 1, relative_base)
        param2 = get_parameter(program, pointer, 2, relative_base)
        param3 = get_parameter(program, pointer, 3, relative_base)

        if instruction == 1:
            write(program, param3, read(program, param1) + read(program, param2))
            pointer += 4
        elif instruction == 2:
            write(program, param3, read(program, param1) * read(program, param2))
            pointer += 4
        elif instruction == 3:
            write(program, param1, process_input(id))
            pointer += 2
        elif instruction == 4:
            process_output(read(program, param1), id)
            pointer += 2
        elif instruction == 5:
            if read(program, param1):
                pointer = read(program, param2)
            else:
                pointer += 3
        elif instruction == 6:
            if not read(program, param1):
                pointer = read(program, param2)
            else:
                pointer += 3
        elif instruction == 7:
            write(
                program,
                param3,
                1 if read(program, param1) < read(program, param2) else 0,
            )
            pointer += 4
        elif instruction == 8:
            write(
                program,
                param3,
                1 if read(program, param1) == read(program, param2) else 0,
            )
            pointer += 4
        elif instruction == 9:
            relative_base += read(program, param1)
            pointer += 2
        else:
            print("Error")
