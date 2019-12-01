from aocd import get_data

mini, maxi = get_data(day=4, year=2019).split("-")


def valid1(number):
    double = False
    decreases = False
    for i in range(1, len(number)):
        if number[i] == number[i - 1]:
            double = True
        if number[i - 1] > number[i]:
            decreases = True
    return double and not decreases


def valid2(number):
    double = False
    decreases = False
    for i in range(1, len(number)):
        if (
            number[i] == number[i - 1]
            and (i == 1 or number[i - 2] != number[i])
            and (i == 5 or number[i + 1] != number[i])
        ):
            double = True
        if number[i - 1] > number[i]:
            decreases = True
    return double and not decreases


total1 = total2 = 0
for i in range(int(mini), int(maxi)):
    if valid1(str(i)):
        total1 += 1
    if valid2(str(i)):
        total2 += 1

print("2019 Day 04")
print("\tPart 1:", total1)
print("\tPart 2:", total2)
