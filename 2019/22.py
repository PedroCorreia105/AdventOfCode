from aocd import get_data

data = get_data(day=22, year=2019).splitlines()


def shuffle(lenght: int):
    stack = list(range(lenght))

    for line in data:
        commands = line.split()
        if commands[0] == "cut":
            value = int(commands[1])
            stack = stack[value:] + stack[:value]

        elif commands[0] == "deal" and commands[1] == "with":
            value = int(commands[3])
            new_stack = [0] * lenght
            step = 0
            for i in range(lenght):
                new_stack[step] = stack[i]
                step = (step + value) % lenght
            stack = new_stack

        elif line == "deal into new stack":
            stack = stack[::-1]

    return stack.index(2019)

# convert rules to linear polynomial.
# (g∘f)(x) = g(f(x))
def parse(lenth):
    a, b = 1, 0
    for line in data[::-1]:
        commands = line.split()
        if commands[0] == "cut":
            value = int(commands[1])
            b = (b + value) % lenth
        elif commands[0] == "deal" and commands[1] == "with":
            value = int(commands[3])
            z = pow(value, lenth - 2, lenth)
            a = a * z % lenth
            b = b * z % lenth
        elif line == "deal into new stack":
            a *= -1
            b = lenth - b - 1
    return a, b


# modpow the polynomial: (ax+b)^m % n
# f(x) = ax+b
# g(x) = cx+d
# f^2(x) = a(ax+b)+b = aax + ab+b
# f(g(x)) = a(cx+d)+b = acx + ad+b
def polypow(a, b, m, n):
    if m == 0:
        return 1, 0
    if m % 2 == 0:
        return polypow(a * a % n, (a * b + b) % n, m // 2, n)
    else:
        c, d = polypow(a, b, m - 1, n)
        return a * c % n, (a * d + b) % n

# Based on https://www.reddit.com/r/adventofcode/comments/ee0rqi/comment/fbwauzi/?utm_source=share&utm_medium=web2x&context=3
def shuffle2(length, times, position):
    a, b = parse(length)
    a, b = polypow(a, b, times, length)
    return (position * a + b) % length

print("2019 Day 22")
print("\tPart 1:", shuffle(10007))
length = 119315717514047
times = 101741582076661
print("\tPart 2:", shuffle2(length, times, 2020))
