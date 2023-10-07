from aocd import get_data

data = list(map(int, get_data(day=9, year=2020).splitlines()))
pointer = 25
vulnerability = False

while not vulnerability:
    vulnerability = data[pointer]
    for lower in range(pointer - 25, pointer - 1):
        for upper in range(lower, pointer):
            if data[lower] + data[upper] == data[pointer]:
                vulnerability = False
    pointer += 1

print("2020 Day 09")
print("\tPart 1:", vulnerability)

for lower, total in enumerate(data):
    for upper in range(lower + 1, len(data)):
        total += data[upper]
        if total == vulnerability:
            print(
                "\tPart 2:", min(data[lower : upper + 1]) + max(data[lower : upper + 1])
            )
        elif total > vulnerability:
            break
