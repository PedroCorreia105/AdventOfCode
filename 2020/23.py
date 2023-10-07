from aocd import get_data

data1 = list(map(int, get_data(day=23, year=2020)))
data2 = data1 + list(range(10, 1_000_001))

class Number:
    def __init__(self, number: int, next_number: int):
        self.number = number
        self.next = next_number

# Create a Linked List
def create_numbers_list(data):
    numbers = {data[len(data) - 1]: Number(data[len(data) - 1], 0)}
    for i in range(len(data) - 2, -1, -1):
        numbers[data[i]] = Number(data[i], numbers[data[i + 1]])
    numbers[data[len(data) - 1]].next = numbers[data[0]]
    return numbers

def process_numbers(data, numbers, loops):
    current_element = numbers[data[0]]
    for _ in range(loops):
        first = current_element.next
        second = first.next
        third = second.next

        searching_for = (
            current_element.number - 1 if current_element.number - 1 > 0 else len(data)
        )

        while searching_for in (first.number, second.number, third.number):
            searching_for = searching_for - 1 if searching_for - 1 > 0 else len(data)

        current_element.next = third.next
        third.next = numbers[searching_for].next
        numbers[searching_for].next = first

        current_element = current_element.next
    return numbers

numbers1 = process_numbers(data1, create_numbers_list(data1), 100)
numbers2 = process_numbers(data2, create_numbers_list(data2), 10_000_000)

current = numbers1[1]
answer = ""
while current.next.number != 1:
	answer += str(current.next.number)
	current = current.next
        
print("2020 Day 23")
print("\tPart 1:", answer)
print("\tPart 2:", numbers2[1].next.number * numbers2[1].next.next.number)
