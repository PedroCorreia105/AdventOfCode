from re import search
from aocd import get_data

data = [line.replace("\n", " ") for line in get_data(day=4, year=2020).split("\n\n")]
valid = total_valid = 0

for passport in data:
    if "pid:" in passport and "ecl:" in passport and "hcl:" in passport:
        if (
            "hgt:" in passport
            and "eyr:" in passport
            and "iyr:" in passport
            and "byr:" in passport
        ):
            valid += 1
print("2020 Day 04")
print("\tPart 1:", valid)

for passport in data:
    byr = search(r"byr:(\d{4})\b", passport)
    iyr = search(r"iyr:(\d{4})\b", passport)
    eyr = search(r"eyr:(\d{4})\b", passport)
    hgtcm = search(r"hgt:(\d+)cm\b", passport)
    hgtin = search(r"hgt:(\d+)in\b", passport)
    hcl = search(r"hcl:#[0-9a-f]{6}\b", passport)
    ecl = search(r"ecl:(...)\b", passport)
    pid = search(r"pid:(\d{9})\b", passport)

    valid = byr and iyr and eyr and hcl and ecl and pid
    valid = valid and 1920 <= int(byr.group(1)) <= 2002
    valid = valid and 2010 <= int(iyr.group(1)) <= 2020
    valid = valid and 2020 <= int(eyr.group(1)) <= 2030
    valid = valid and (
        hgtcm
        and 150 <= int(hgtcm.group(1)) <= 193
        or hgtin
        and 59 <= int(hgtin.group(1)) <= 76
    )
    valid = valid and ecl.group(1) in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]

    if valid:
        total_valid += 1

print("\tPart 2:", total_valid)
