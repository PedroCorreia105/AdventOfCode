from aocd import get_data

data = get_data(day=19, year=2020).split("\n\n")


def calculate_rules(part: int):
    rules = {}
    if part == 2:
        lines = (
            data[0]
            .replace("8: 42", "8: 42 | 42 8")
            .replace("11: 42 31", "11: 42 31 | 42 11 31")
            .splitlines()
        )
    else:
        lines = data[0].splitlines()

    for line in lines:
        rule = line.split(": ")
        if '"' in rule[1]:
            rules[rule[0]] = rule[1][1:-1]
        elif "|" in rule[1]:
            rules[rule[0]] = [
                possibility.split() for possibility in rule[1].split(" | ")
            ]
        else:
            rules[rule[0]] = [rule[1].split()]
    return rules


def match(rules, string, rule, pending):
    if isinstance(rules[rule], str):
        if pending == [] and len(string) == 1 and string[0] == rules[rule]:
            return True
        elif pending != [] and len(string) > 0 and string[0] == rules[rule]:
            return match(rules, string[1:], pending[0], pending[1:])
        else:
            return False
    else:
        return any(
            match(rules, string, possibility[0], possibility[1:] + pending)
            for possibility in rules[rule]
        )


print("2020 Day 19")
print(
    "\tPart 1:",
    sum(match(calculate_rules(1), m, "0", []) for m in data[1].splitlines()),
)
print(
    "\tPart 2:",
    sum(match(calculate_rules(2), m, "0", []) for m in data[1].splitlines()),
)
