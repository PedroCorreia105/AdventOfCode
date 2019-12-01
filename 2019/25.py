from itertools import chain, combinations
from aocd import get_data
from intcode import run_program

program = list(map(int, get_data(day=25, year=2019).split(",")))
previous_location = previous_direction = direction_to_security = None
combinations_list, inventory, input_list = [], [], []
possible_directions = ["north", "east", "south", "west"]
input_index = -1
locations = {}
text = ""


class Location:
    def __init__(self, name: str, directions: list):
        self.name = name
        self.neighbours = {direction: None for direction in directions}
        self.visited = False


def parse_text(text: str):
    doors, items = [], []
    listing_doors = listing_items = False
    location = ""

    for line in text.splitlines():
        if "==" in line:
            location = line[3:-3]
        elif line == "Doors here lead:":
            listing_doors = True
        elif listing_doors and "- " in line:
            doors.append(line[2:])
        elif line == "Items here:":
            listing_items = True
        elif listing_items and "- " in line:
            items.append(line[2:])
        else:
            listing_doors = listing_items = False

    return location, doors, items


def go_to(current_location: Location, destination: Location):
    # reset locations visited which is used to prevent cycles
    for location in locations:
        locations[location].visited = False

    return go_to2(current_location, destination)


def go_to2(current_location: Location, destination: Location):
    current_location.visited = True
    # if the destination is a neighbour
    for direction in current_location.neighbours:
        if current_location.neighbours[direction] == destination:
            return direction

    # search for the destination in the neighbour's neighbours
    for direction in current_location.neighbours:
        neighbour = current_location.neighbours[direction]
        if not neighbour.visited and go_to2(neighbour, destination):
            return direction


def choose_next_move(location_name: str, directions: list, items: list):
    global previous_location, previous_direction, combinations_list, direction_to_security, inventory

    # step 1: record all locations
    if location_name not in locations:
        locations[location_name] = Location(location_name, directions)

    # link the previous with the current location
    if previous_location and previous_location.neighbours[previous_direction] == None:
        previous_location.neighbours[previous_direction] = locations[location_name]
        locations[location_name].neighbours[
            possible_directions[(possible_directions.index(previous_direction) + 2) % 4]
        ] = previous_location

    # while visiting locations, grab every acceptable item
    item_action = ""
    dont_take_items = [
        "escape pod",
        "infinite loop",
        "photons",
        "giant electromagnet",
        "molten lava",
    ]
    if combinations_list == [] and items != []:
        for item in items:
            if item not in dont_take_items:
                inventory.append(item)
                item_action += f"take {item}\n"

    # search for a path not taken before (represented be None)
    if direction := go_to(locations[location_name], None):
        previous_location = locations[location_name]
        previous_direction = direction
        return item_action + direction

    # step 2: go to Security Checkpoint
    if location_name != "Security Checkpoint":
        return (
            item_action
            + "inv\n"
            + go_to(locations[location_name], locations["Security Checkpoint"])
        )

    # calculate once, the direction to the Security Checkpoint
    if direction_to_security == None:
        direction_to_security = go_to(
            locations["Security Checkpoint"], locations["Security Checkpoint"]
        )

    # step 3: try all possible combinations of items
    if combinations_list == []:
        combinations_list = chain(
            *map(
                lambda x: combinations(inventory, x),
                range(len(inventory)),
            )
        )

    combination = next(combinations_list, "-1")

    for item in inventory:
        if item not in combination:
            item_action += f"drop {item}\n"
            inventory.remove(item)
            item_action += f"drop {item}\n"
    for item in combination:
        if item not in inventory:
            inventory.append(item)
            item_action += f"take {item}\n"
    return item_action + direction_to_security


def process_input(_):
    global text, input_list, input_index, inventory

    if input_index == len(input_list) - 1:
        input_list = list(map(ord, choose_next_move(*parse_text(text)) + "\n"))
        input_index = 0
        text = ""
    else:
        input_index += 1
    return input_list[input_index]


def process_output(output: int, _):
    global text
    text += chr(output)


run_program(program, process_input, process_output)

print("2019 Day 25")
print("\tPart 1:", text.split("\n")[-2])
