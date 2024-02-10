def FinaLocation(finalData, time):
    return finalData[0] * finalData[1] * int(
        time / (finalData[2] + finalData[1])
    ) + finalData[0] * min(finalData[1], time % (finalData[2] + finalData[1]))


def Part1(listOfReindeers, maxTime):
    return max([FinaLocation(finalData, maxTime) for finalData in listOfReindeers])


def Part2(listOfReindeers, maxTime):
    time = 1
    points = [0]*len(listOfReindeers)
    while time <= maxTime:
        newList = [FinaLocation(finalData, time) for finalData in listOfReindeers]
        newMax = max(newList)
        for i,x in enumerate(newList):
            if(x==newMax):
                points[i]+=1
        time += 1
    return max(points)

with open("input.txt") as f:
    maxTime = 2503
    listOfReindeers = []
    for d in f.readlines():
        x = d.strip()
        x = x.replace("can fly", "")
        x = x.replace("for", "")
        x = x.replace("seconds, but then must rest", "")
        x = x.replace("seconds.", "")
        data = x.split(" ")
        listOfReindeers.append(
            (int(data[2].strip()), int(data[5].strip()), int(data[-2].strip()))
        )
    print(Part1(listOfReindeers, maxTime))
    print(Part2(listOfReindeers, maxTime))
