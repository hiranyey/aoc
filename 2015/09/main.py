globalMax = 0
globalMin = 2**14
globalCities = 2**14

def travellingSalesman(map, start, count):
    global globalCities
    if len(start) == globalCities:
        global globalMin
        globalMin = min(globalMin,count)
        global globalMax
        globalMax = max(globalMax,count)
    for k, v in map.items():
        newMap = map.copy()
        del newMap[k]
        newStart = start.copy()
        if k[0] not in start[0:-1] and k[1] not in start[0:-1] and k[0] == start[-1]:
            newStart.append(k[1])
            travellingSalesman(newMap, newStart, v + count)
        if k[0] not in start[0:-1] and k[1] not in start[:-1] and k[1] == start[-1]:
            newStart.append(k[0])
            travellingSalesman(newMap, newStart, v + count)


with open("input.txt") as f:
    map = {}
    cities = set()
    for line in f.readlines():
        data = line.strip().split(" ")
        cities.add(data[0])
        cities.add(data[2])
        map[(data[0], data[2])] = int(data[-1])
    globalCities = len(cities)
    for k, v in map.items():
        newMap = map.copy()
        del newMap[k]
        travellingSalesman(newMap, [k[0],k[1]], v)
        travellingSalesman(newMap, [k[1],k[0]], v)
    print(globalMax)
    print(globalMin)
