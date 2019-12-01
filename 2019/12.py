from math import gcd
from re import match
from aocd import get_data

data = get_data(day=12, year=2019).splitlines()
planet_position = [
    [int(i) for i in match("<x=(.*), y=(.*), z=(.*)>", line).groups()] for line in data
]
planet_velocity = [[0] * 3, [0] * 3, [0] * 3, [0] * 3]
step = total_energy = 0

# Records the location and speed of the planets on the first step for each axis
memory = [[], [], []]

# Records how long it takes for each axis to return to the same exact position and speed
steps = [0, 0, 0]

while 0 in steps:
    planet_gravity = [[0] * 3, [0] * 3, [0] * 3, [0] * 3]
    for planet1 in range(4):
        for planet2 in range(planet1 + 1, 4):
            for axis in range(3):
                if planet_position[planet1][axis] < planet_position[planet2][axis]:
                    planet_gravity[planet1][axis] += 1
                    planet_gravity[planet2][axis] -= 1
                elif planet_position[planet1][axis] > planet_position[planet2][axis]:
                    planet_gravity[planet1][axis] -= 1
                    planet_gravity[planet2][axis] += 1

    for planet in range(4):
        for axis in range(3):
            planet_velocity[planet][axis] += planet_gravity[planet][axis]
            planet_position[planet][axis] += planet_velocity[planet][axis]

    for axis in range(3):
        if step == 0:
            memory[axis] = [planet[axis] for planet in planet_position] + [
                planet[axis] for planet in planet_velocity
            ]
        elif not steps[axis]:
            position_and_velocity = [planet[axis] for planet in planet_position] + [
                planet[axis] for planet in planet_velocity
            ]
            if position_and_velocity == memory[axis]:
                steps[axis] = step
    step += 1

    if step == 1000:
        for planet in range(4):
            potential = sum([abs(planet_position[planet][axis]) for axis in range(3)])
            kinetic = sum([abs(planet_velocity[planet][axis]) for axis in range(3)])
            total_energy += potential * kinetic
        print("2019 Day 12")
        print("\tPart 1:", total_energy)


# Calculate Least Common Multiple of the 3 axis
lcm = steps[0]
for step in steps[1:]:
    lcm = lcm * step // gcd(lcm, step)

print("\tPart 2:", lcm)
