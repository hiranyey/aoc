def convert(c):
    if c == "(":
        return 1
    else:
        return -1


def Part1():
    with open("input.txt") as f:
        return sum([convert(d) for d in f.read()])


def Part2():
    with open("input.txt") as f:
        count = 0
        for i, d in enumerate(f.read()):
            count += convert(d)
            if count < 0:
                return i + 1


print(Part1())
print(Part2())
