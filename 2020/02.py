from re import findall
from aocd import get_data

data = get_data(day=2, year=2020).splitlines()

total1 = 0
for line in data:
    mini, maxi, char, word = findall(r"(\d+)-(\d+) (.): (.*)", line)[0]
    mini, maxi = int(mini), int(maxi)

    if mini <= word.count(char) <= maxi:
        total1 += 1


total2 = 0
for line in data:
    i1, i2, char, word = findall(r"(\d+)-(\d+) (.): (.*)", line)[0]
    i1, i2 = int(i1) - 1, int(i2) - 1

    if (word[i1] == char) + (word[i2] == char) == 1:
        total2 += 1

print("2020 Day 02")
print("\tPart 1:", total1)
print("\tPart 2:", total2)
