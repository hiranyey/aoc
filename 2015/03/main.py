def move(currentPosition, x):
    if x == "^":
        currentPosition = (currentPosition[0] + 1, currentPosition[1])
    elif x == "v":
        currentPosition = (currentPosition[0] - 1, currentPosition[1])
    elif x == ">":
        currentPosition = (currentPosition[0], currentPosition[1] + 1)
    elif x == "<":
        currentPosition = (currentPosition[0], currentPosition[1] - 1)
    return currentPosition


def Part1():
    with open("input.txt") as f:
        steps = {(0, 0): 1}
        currentPosition = (0, 0)
        for x in f.read():
            currentPosition = move(currentPosition, x)
            steps[currentPosition] = steps.get(currentPosition, 0) + 1
        return len(steps)


def Part2():
    with open("input.txt") as f:
        steps = {(0, 0): 1}
        santaPosition = (0, 0)
        robotPosition = (0, 0)
        for i, x in enumerate(f.read()):
            if i % 2 == 0:
                santaPosition = move(santaPosition, x)
                steps[santaPosition] = steps.get(santaPosition, 0) + 1
            else:
                robotPosition = move(robotPosition, x)
                steps[robotPosition] = steps.get(robotPosition, 0) + 1
        return len(steps)


print(Part1())
print(Part2())
