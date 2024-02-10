def processNumber(x):
    counter = []
    temp = x[0]
    count = 0
    for y in x:
        if y==temp:
            count+= 1
        else:
            counter.append((count,temp))
            count = 1
            temp = y
    counter.append((count,temp))
    result = ""
    for x in counter:
        result += str(x[0])+x[1]
    return result

def Part1(times):
    with open("input.txt") as f:
        lines = f.readlines()
        counter = 0
        target = lines[0].strip()
        while counter < times:
            target = processNumber(target)
            counter += 1
        return len(target)

print(Part1(40))
print(Part1(50))