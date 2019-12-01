from aocd import get_data

data = get_data(day=18, year=2019).splitlines()
vectors = [(0, 1), (0, -1), (1, 0), (-1, 0)]


class Node:
    def __init__(self, coordinates):
        self.coordinates = coordinates
        self.neighbours = []
        self.was_visited = False
        self.distance = self.doors_required = self.door = self.key = 0


class Path:
    def __init__(self, locations, keys, distance, max_keys):
        self.locations = locations
        self.max_keys = max_keys
        self.keys = keys
        self.distance = distance

    def __repr__(self):
        return repr(self.locations) + repr(self.keys)

    def is_completed(self):
        return self.keys == self.max_keys


def key_is_not_in_possession(keys, key):
    return int(str(keys), 2) & int(str(key), 2) == 0


def keys_unlock_doors(keys, doors):
    return int(str(doors), 2) & int(str(keys), 2) == int(str(doors), 2)


# For part2
def change_data(data):
    for y in range(1, len(data) - 1):
        for x in range(1, len(data[0]) - 1):
            if data[y][x] == "@":
                data[y - 1] = data[y - 1][: x - 1] + "@#@" + data[y - 1][x + 2 :]
                data[y] = data[y][: x - 1] + "###" + data[y][x + 2 :]
                data[y + 1] = data[y + 1][: x - 1] + "@#@" + data[y + 1][x + 2 :]
                return data


# Create a graph as an adjacency list where 1 coordinate is 1 node
def create_graph(data):
    graph = {}
    starting_points = []
    key_list = []
    for y in range(1, len(data) - 1):
        for x in range(1, len(data[0]) - 1):
            if data[y][x] != "#":
                node = Node((x, y))
                graph[repr((x, y))] = node

                if data[y][x] == "@":
                    starting_points.append(node)
                elif data[y][x].islower():
                    node.key = data[y][x]
                    key_list.append(node)
                elif data[y][x] != ".":
                    node.door = data[y][x]

                for v in vectors:
                    x1, y1 = x + v[0], y + v[1]
                    if data[y1][x1] != "#" and repr((x1, y1)) in graph:
                        node.neighbours.append(graph[repr((x1, y1))])
                        graph[repr((x1, y1))].neighbours.append(node)
    return graph, starting_points, key_list


# Recursively search for the shortest path as a DFS taking into account the doors
def visit(adjacency_matrix, point, keys, cache):
    if repr(point) + repr(keys) in cache:
        return cache[repr(point) + repr(keys)]

    minimum = 0
    for destiny in adjacency_matrix[point]:
        if (
            destiny != point
            and key_is_not_in_possession(keys, destiny.key)
            and keys_unlock_doors(keys, adjacency_matrix[point][destiny][0])
        ):
            result = (
                visit(adjacency_matrix, destiny, keys + destiny.key, cache)
                + adjacency_matrix[point][destiny][1]
            )
            minimum = result if minimum == 0 else min(minimum, result)
    cache[repr(point) + repr(keys)] = minimum
    return minimum


# Create a weighted graph as an adjacency_matrix where 1 node is 1 key (plus the starting positions)
# Weights are the distance + the doors between any two keys
def create_adjacency_matrix(graph, key_list, starting_points):
    adjacency_matrix = {}

    # Convert door and key char (a, A) to 10 ** index of the key in key_list, to allow "binary" operations
    for node in graph.values():
        if node.door:
            node.door = 10 ** next(
                i for i, v in enumerate(key_list) if node.door.lower() == v.key
            )

    for i, node in enumerate(key_list):
        node.key = 10**i

    for key1 in key_list + starting_points:
        adjacency_matrix[key1] = {}

        for key2 in key_list:
            adjacency_matrix[key1][key2] = [-1, float("inf")]

        to_visit = [key1]

        while to_visit != []:
            node = to_visit.pop(0)
            node.was_visited = True

            if (
                node.door
                and int(str(node.doors_required), 2) & int(str(node.door), 2) == 0
            ):
                node.doors_required += node.door

            for neighbour in node.neighbours:
                if not neighbour.was_visited:
                    to_visit.append(neighbour)
                    neighbour.distance = node.distance + 1
                    neighbour.doors_required = node.doors_required

            if node.key:
                adjacency_matrix[key1][node][0] = node.doors_required
                adjacency_matrix[key1][node][1] = node.distance

        for node in graph.values():
            node.distance = node.doors_required = 0
            node.was_visited = False
    return adjacency_matrix


graph1, starting_points1, key_list1 = create_graph(data)
adjacency_matrix1 = create_adjacency_matrix(graph1, key_list1, starting_points1)

print("2019 Day 18")
print("\tPart 1:", visit(adjacency_matrix1, starting_points1[0], 0, {}))

# Part 2
# Basically the same as the above without recursion
graph2, starting_points2, key_list2 = create_graph(change_data(data))
adjacency_matrix2 = create_adjacency_matrix(graph2, key_list2, starting_points2)
max_keys = sum(10**i for i in range(len(key_list2)))
paths = [Path(starting_points2, 0, 0, max_keys)]
cache = {}
minimum_distance = float("inf")

while paths != [] and minimum_distance > 2020:
    path = paths.pop(-1)

    if path.is_completed():
        minimum_distance = min(path.distance, minimum_distance)
    elif path.distance >= minimum_distance:
        continue
    elif repr(path) not in cache or cache[repr(path)] > path.distance:
        cache[repr(path)] = path.distance

        for i in range(4):
            for destiny in adjacency_matrix2[path.locations[i]]:
                if (
                    destiny != path.locations[i]
                    and key_is_not_in_possession(path.keys, destiny.key)
                    and keys_unlock_doors(
                        path.keys, adjacency_matrix2[path.locations[i]][destiny][0]
                    )
                ):
                    paths.append(
                        Path(
                            path.locations[:i] + [destiny] + path.locations[i + 1 :],
                            path.keys + destiny.key,
                            path.distance
                            + adjacency_matrix2[path.locations[i]][destiny][1],
                            max_keys,
                        )
                    )

print("\tPart 2:", minimum_distance)
