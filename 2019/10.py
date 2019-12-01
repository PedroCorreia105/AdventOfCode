from aocd import get_data
import numpy as np

data = get_data(day=10, year=2019).splitlines()
width, height = len(data[0]), len(data)
asteroids = [(x, y) for y in range(height) for x in range(width) if data[y][x] == "#"]
asteroids_in_sight = [0] * len(asteroids)


def angle_between(a, b):
    angle = np.rad2deg(np.arctan2(np.cross(a, b), np.dot(a, b)))
    return angle + 360 if angle < 0 else angle


def can_see_each_other(asteroid1, asteroid2):
    vector = np.subtract(asteroid1, asteroid2)
    divisor = np.gcd(vector[0], vector[1])
    vector = vector // divisor
    position = np.sum([asteroid2, vector], axis=0)

    while not np.array_equal(asteroid1, position):
        if data[position[1]][position[0]] == "#":
            return False
        position = np.sum([position, vector], axis=0)
    return True


def asteroids_in_between(center, asteroid):
    vector = np.subtract(center, asteroid)
    divisor = np.gcd(vector[0], vector[1])
    vector = vector // divisor
    position = np.sum([asteroid, vector], axis=0)
    order = 0

    while not np.array_equal(center, position):
        if data[position[1]][position[0]] == "#":
            order += 1
        position = np.sum([position, vector], axis=0)
    return order


for asteroid1_index, asteroid1 in enumerate(asteroids):
    for asteroid2_index in range(asteroid1_index + 1, len(asteroids)):
        if can_see_each_other(asteroid1, asteroids[asteroid2_index]):
            asteroids_in_sight[asteroid1_index] += 1
            asteroids_in_sight[asteroid2_index] += 1

maximum = max(asteroids_in_sight)
center = asteroids[asteroids_in_sight.index(maximum)]
asteroids.remove(center)
ordered_asteroids = sorted(
    [
        (
            asteroids_in_between(center, asteroid),
            angle_between((0, -1), np.subtract(asteroid, center)),
            asteroid,
        )
        for asteroid in asteroids
    ]
)


print("2019 Day 10")
print("\tPart 1:", maximum)
print("\tPart 2:", ordered_asteroids[199][2][0] * 100 + ordered_asteroids[199][2][1])
