from aocd import get_data

data = get_data(day=18, year=2020).splitlines()


def evaluate1(line):
    while isinstance(line, str) and "(" in line:
        last_index = line.find(")")
        first_index = line[:last_index].rfind("(")
        line = (
            line[:first_index]
            + str(evaluate1(line[first_index + 1 : last_index]))
            + line[last_index + 1 :]
        )
    while isinstance(line, str) and ("*" in line or "+" in line):
        index = max(line.rfind("*"), line.rfind("+"))
        if index == line.rfind("*"):
            line = evaluate1(line[:index]) * evaluate1(line[index + 1 :])
        else:
            line = evaluate1(line[:index]) + evaluate1(line[index + 1 :])
    return int(line)


def evaluate2(line):
    while isinstance(line, str) and "(" in line:
        last_index = line.find(")")
        first_index = line[:last_index].rfind("(")
        line = (
            line[:first_index]
            + str(evaluate2(line[first_index + 1 : last_index]))
            + line[last_index + 1 :]
        )
    while isinstance(line, str) and "*" in line:
        index = line.find("*")
        line = evaluate2(line[:index]) * evaluate2(line[index + 1 :])
    while isinstance(line, str) and "+" in line:
        index = line.find("+")
        line = evaluate2(line[:index]) + evaluate2(line[index + 1 :])
    return int(line)


print("2020 Day 18")
print("\tPart 1:", sum([evaluate1(line.replace(" ", "")) for line in data]))
print("\tPart 2:", sum([evaluate2(line.replace(" ", "")) for line in data]))
