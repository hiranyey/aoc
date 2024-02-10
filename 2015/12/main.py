import json

def findIterator(x):
    if type(x) == dict:
        return x.values()
    else:
        return x

def Part1(x):
    total = 0
    for value in findIterator(x):
        if type(value) == dict:
            total += Part1(value)
        elif type(value) == list:
            total += Part1(value)
        elif type(value) == int:
            total += value
    return total


def Part2(x):
    total = 0
    for value in findIterator(x):
        if type(x) == dict and type(value) == str and value == "red":
            return 0
        if type(value) == dict:
            total += Part2(value)
        elif type(value) == list:
            total += Part2(value)
        elif type(value) == int:
            total += value
    return total


with open("input.txt") as f:
    x = json.loads(f.readline().strip())
    print(Part1(x))
    print(Part2(x))
