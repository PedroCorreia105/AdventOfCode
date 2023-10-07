from collections import deque
from aocd import get_data

data = [
    list(map(int, player.splitlines()[1:]))
    for player in get_data(day=22, year=2020).split("\n\n")
]
player1, player2 = deque(data[0]), deque(data[1])


def play1(player1, player2):
    while len(player1) != 0 and len(player2) != 0:
        a, b = player1.popleft(), player2.popleft()
        winner = 1 if a > b else 2

        if winner == 1:
            player1.extend([a, b])
        else:
            player2.extend([b, a])

    return 1, player1 if len(player2) == 0 else 2, player2

def play2(player1, player2):
    states = []
    while len(player1) != 0 and len(player2) != 0:
        if repr(player1) + repr(player2) in states:
            return 1
        states.append(repr(player1) + repr(player2))

        a, b = player1.popleft(), player2.popleft()

        if len(player1) >= a and len(player2) >= b:
            copy1, copy2 = player1.copy(), player2.copy()
            while len(copy1) != a:
                copy1.pop()
            while len(copy2) != b:
                copy2.pop()
            winner = play2(copy1, copy2)
        else:
            winner = 1 if a > b else 2

        if winner == 1:
            player1.extend([a, b])
        else:
            player2.extend([b, a])

    return 1 if len(player2) == 0 else 2


print("2020 Day 22")
winner = play1(player1.copy(), player2.copy())[1]
print("\tPart 1:", sum(winner.pop() * (i + 1) for i in range(len(winner))))

winner = player1 if play2(player1, player2) == 1 else player2
print("\tPart 2:", sum(winner.pop() * (i + 1) for i in range(len(winner))))
