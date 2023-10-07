from aocd import get_data

data = list(map(int, get_data(day=15, year=2020).split(",")))
dic = {element: data.index(element) for element in data}
length = len(data)
current_element = 0

print("2020 Day 15")

while length < 30000000 - 1:
    if current_element in dic:
        index = dic[current_element]
        dic[current_element] = length
        current_element = length - index
    else:
        dic[current_element] = length
        current_element = 0
    length += 1

    if length == 2019:
        print("\tPart 1:", current_element)

print("\tPart 2:", current_element)
