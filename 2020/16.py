from re import match
from aocd import get_data

data = get_data(day=16, year=2020).split("\n\n")

# The 2 ranges are [0] to [1] and [2] to [3], field name is in [5]
# The set in [4] are the indexes of the ticket that cannot be that field
ticket_fields = []
for line in data[0].splitlines():
    matches = match("(.*): (.*)-(.*) or (.*)-(.*)", line).groups()
    ticket_fields.append(list(map(int, matches[1:])) + [set(), matches[0]])

my_ticket = list(map(int, data[1].splitlines()[1].split(",")))
nearby_tickets = [list(map(int, line.split(","))) for line in data[2].splitlines()[1:]]


def belongs_to_field(field, number):
    return field[0] <= number <= field[1] or field[2] <= number <= field[3]


error_rate = 0
for ticket in nearby_tickets:
    for field_index, field in enumerate(ticket):
        # If ticket_index does not belong to any field then the ticket in invalid
        if not any(
            [belongs_to_field(ticket_field, field) for ticket_field in ticket_fields]
        ):
            error_rate += field
        else:
            for ticket_field in ticket_fields:
                if not belongs_to_field(ticket_field, field):
                    ticket_field[4].add(field_index)

# Sort by the length of the list of impossible indexes in ticket
ticket_fields = sorted(
    ticket_fields, key=lambda ticket_field: len(ticket_field[4]), reverse=True
)
dic = {
    ticket_field[5]: [
        possible
        for possible in range(len(my_ticket))
        if possible not in ticket_field[4]
    ]
    for ticket_field in ticket_fields
}
# {Field Name: Possible indexes in ticket}

total = 1
for element in dic:
    # Since the list of possible indexes in ticket is always equal to the previous + the correct element
    dic[element] = list(filter(lambda x: x not in dic.values(), dic[element]))[0]
    if element.startswith("departure"):
        total *= my_ticket[dic[element]]


print("2020 Day 16")
print("\tPart 1:", error_rate)
print("\tPart 2:", total)
