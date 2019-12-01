from aocd import get_data
from intcode import run_program
import threading
import concurrent.futures
from time import sleep

program = list(map(int, get_data(day=23, year=2019).split(",")))
lock = threading.Lock()
number_of_computers = 50

messages = {i: [i] for i in range(number_of_computers)}
messages_to_process = {i: [] for i in range(number_of_computers)}
messages_to_process[255] = []
nat_x, nat_y, prev_nat_y = None, None, None
should_thread_die = False
idle_count = 0

print("2019 Day 23")


def launch_computer(i):
    run_program(program, process_input, process_output, should_threads_continue, id=i)


def should_threads_continue():
    return not should_thread_die


def is_idle():
    return idle_count > 200 and sum(len(messages[i]) for i in messages) == 0


def process_input(id):
    global prev_nat_y, should_thread_die, messages, idle_count
    with lock:
        if is_idle() and nat_x != None:
            idle_count = 0
            messages[0] += [nat_x, nat_y]
            if prev_nat_y == nat_y:
                print("\tPart 2:", nat_y)
                should_thread_die = True
            prev_nat_y = nat_y
        if len(messages[id]) > 0:
            idle_count = 0
            return messages[id].pop(0)
        else:
            idle_count += 1
            return -1


def process_output(output, id):
    global nat_x, nat_y, messages_to_process, messages
    with lock:
        messages_to_process[id].append(output)
        if len(messages_to_process[id]) == 3:
            destination, x, y = messages_to_process[id]
            messages_to_process[id] = []
            if destination == 255:
                if nat_x == None:
                    print("\tPart 1:", y)
                nat_x, nat_y = x, y
            else:
                messages[destination] += [x, y]


with concurrent.futures.ThreadPoolExecutor(max_workers=number_of_computers) as executor:
    futures = [executor.submit(launch_computer, i) for i in range(number_of_computers)]

    for future in concurrent.futures.as_completed(futures):
        result = future.result()
        # Wait for threads

# 19724
# 15252
