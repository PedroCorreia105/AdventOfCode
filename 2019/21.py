from aocd import get_data
from intcode import run_program

program = list(map(int, get_data(day=21, year=2019).split(",")))

def process_input(_):
    global index
    index += 1
    return ord(input_text[index])


def process_output(output, part):
    if output > 1114111:
        if part == 1:
            print("2019 Day 21")
            print("\tPart 1:", output)
        else:
            print("\tPart 2:", output)

input_text = "NOT A T\nNOT B J\nOR T J\nNOT C T\nOR T J\nAND D J\nWALK\n"
index = -1
run_program(program, process_input, process_output, id=1)

input_text = "NOT A T\nNOT B J\nOR T J\nNOT C T\nOR T J\nAND D J\nNOT E T\nNOT T T\nOR H T\nAND T J\nRUN\n"
index = -1
run_program(program, process_input, process_output, id=2)
