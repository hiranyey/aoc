def Part1():
    with open("input.txt") as f:
        matrix = [[0] * 1000 for i in range(1000)]
        for line in f.readlines():
            numbers = line.split(" ")
            firstCoordinate = numbers[-3]
            x1, y1 = firstCoordinate.split(",")
            firstX, firstY = int(x1), int(y1)
            secondCoordinate = numbers[-1]
            x2, y2 = secondCoordinate.split(",")
            secondX, secondY = int(x2) + 1, int(y2) + 1
            if line.startswith("toggle"):
                for x in range(firstX, secondX):
                    for y in range(firstY, secondY):
                        matrix[x][y] = 1 - matrix[x][y]
            elif line.startswith("turn off"):
                for x in range(firstX, secondX):
                    for y in range(firstY, secondY):
                        matrix[x][y] = 0
            elif line.startswith("turn on"):
                for x in range(firstX, secondX):
                    for y in range(firstY, secondY):
                        matrix[x][y] = 1
        return sum([sum(x) for x in matrix])
    
def Part2():
    with open("input.txt") as f:
        matrix = [[0] * 1000 for i in range(1000)]
        for line in f.readlines():
            numbers = line.split(" ")
            firstCoordinate = numbers[-3]
            x1, y1 = firstCoordinate.split(",")
            firstX, firstY = int(x1), int(y1)
            secondCoordinate = numbers[-1]
            x2, y2 = secondCoordinate.split(",")
            secondX, secondY = int(x2) + 1, int(y2) + 1
            if line.startswith("toggle"):
                for x in range(firstX, secondX):
                    for y in range(firstY, secondY):
                        matrix[x][y]= matrix[x][y]+2
            elif line.startswith("turn off"):
                for x in range(firstX, secondX):
                    for y in range(firstY, secondY):
                        matrix[x][y] = max(matrix[x][y]-1,0)
            elif line.startswith("turn on"):
                for x in range(firstX, secondX):
                    for y in range(firstY, secondY):
                        matrix[x][y] = matrix[x][y]+1
        return sum([sum(x) for x in matrix])

print(Part1())
print(Part2())