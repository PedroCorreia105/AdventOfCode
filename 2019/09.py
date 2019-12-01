from aocd import get_data
from intcode import run_program

program = list(map(int, get_data(day=9, year=2019).split(",")))


def process_input(_):
    return 1


def process_input2(_):
    return 2


def process_output(output, _):
    global outputVal
    outputVal = output


print("2019 Day 09")
run_program(program, process_input, process_output)
print("\tPart 1:", outputVal)
run_program(program, process_input2, process_output)
print("\tPart 2:", outputVal)
