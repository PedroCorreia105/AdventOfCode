from re import search
from re import findall
from aocd import get_data

data = get_data(day=7, year=2020).splitlines()
# wavy olive bags contain 1 striped olive bag, 1 dull cyan bag.
# plaid cyan bags contain no other bags.

rules = []
# ['wavy beige', [('5', 'muted silver'), ('5', 'pale teal')]]
for line in data:
    container = search(r"(.*) bags contain", line).group(1)
    contents = findall(r"(\d+) (.*?) bag", line)
    rules.append([container, contents])

looking_for = ["shiny gold"]
original_length = 0

while original_length != len(looking_for):
    original_length = len(looking_for)
    for rule in rules:
        if rule[-1] != "*" and set(looking_for).intersection(
            [content[1] for content in rule[1]]
        ):
            looking_for.append(rule[0])
            rule.append("*")


def recursive_search(bag: str):
    for rul in rules:
        if rul[0] == bag:
            return 1 + sum(
                [
                    int(contents[0]) * recursive_search(contents[1])
                    for contents in rul[1]
                ]
            )
    return 0


print("2020 Day 07")
print("\tPart 1:", len([rule for rule in rules if rule[-1] == "*"]))
print("\tPart 2:", recursive_search("shiny gold") - 1)
