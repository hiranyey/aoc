def Part1():
    with open("input.txt") as f:
        return sum([len(d.strip()) - len(eval(d.strip())) for d in f.readlines()])


def Part2():
    with open("input.txt") as f:
        return sum(
            [2 + d.strip().count("\\") + d.strip().count('"') for d in f.readlines()]
        )


print(Part1())
print(Part2())
