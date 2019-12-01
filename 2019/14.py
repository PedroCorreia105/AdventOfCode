from re import match
from re import findall
from math import ceil
from random import randint
from aocd import get_data

data = get_data(day=14, year=2019).splitlines()


class Substance:
    def __init__(self, name, quantity, requirements):
        self.name = name
        self.quantity = quantity
        self.requirements = requirements

        self.node_height = 0
        self.needed = 0

    def calculate_height(self, substances):
        if len(self.requirements) == 0:
            self.node_height = 0
        else:
            self.node_height = 1 + max(
                [
                    substances[requirement[1]].calculate_height(substances)
                    for requirement in self.requirements
                ]
            )
        return self.node_height


def calculate(value):
    substances = {"ORE": Substance("ORE", 0, [])}
    ore = 0

    for line in data:
        number, name = match(r".+=> (.+) (.+)", line).groups()
        requirements = findall(r"(\d+) (\w+)", line)[:-1]
        substances[name] = Substance(
            name,
            int(number),
            [[int(requirement[0]), requirement[1]] for requirement in requirements],
        )

    substances["FUEL"].calculate_height(substances)
    substances["FUEL"].needed = value

    while ore == 0:
        substance = max_height = 0
        for s in substances.values():
            if s.needed > 0 and s.node_height > max_height:
                max_height = s.node_height
                substance = s
        if substance == 0:
            ore = substances["ORE"].needed
        else:
            multiplier = ceil(substance.needed / substance.quantity)
            substance.needed = 0
            for requirement in substance.requirements:
                substances[requirement[1]].needed += requirement[0] * multiplier
    return ore


# Smart brute force
limits = [1, 1000000000000]
while limits[0] + 1 != limits[1]:
    i = randint(limits[0], limits[1])
    if calculate(i) > 1000000000000:
        limits[1] = i
    else:
        limits[0] = i


print("2019 Day 14")
print("\tPart 1:", calculate(1))
print("\tPart 2:", limits[0])
