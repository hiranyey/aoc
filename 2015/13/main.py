peopleCount = 0
maximumHapiness = 0

def CalculateHapiness(graph,current):
    count = 0 
    for i,x in enumerate(current):
        count += graph[(x,current[(i+1)%len(current)])]
        count += graph[(x,current[(i-1)%len(current)])]
    global maximumHapiness
    maximumHapiness = max(maximumHapiness,count)


def TravellingSalesman(graph,current,remaining):
    global peopleCount
    if len(current) == peopleCount:
        CalculateHapiness(graph,current)
    for x in remaining:
        peopleCopy = remaining.copy()
        peopleCopy.remove(x)
        currentCopy = current.copy()
        currentCopy.append(x)
        TravellingSalesman(graph,currentCopy,peopleCopy)

with open("input.txt") as f:
    graph = {}
    people = set()
    for line in f.readlines():
        line = line.replace("would","")
        line = line.replace("happiness units by sitting next to","")
        parts = line.strip().split(" ")
        people.add(parts[0])
        people.add(parts[-1][:-1])
        if "gain" in parts[2]:
            graph[(parts[0],parts[-1][:-1])] = int(parts[3])
        else:
            graph[(parts[0],parts[-1][:-1])] = -1*int(parts[3])
    peopleCount = len(people)
    maximumHapiness = 0
    for x in people:
        peopleCopy = people.copy()
        peopleCopy.remove(x)
        TravellingSalesman(graph,[x],peopleCopy)
    print(maximumHapiness) #Part 1
    for x in people:
        graph[("Me",x)]=0
        graph[(x,"Me")]=0
    people.add("Me")
    peopleCount = len(people)
    maximumHapiness = 0
    for x in people:
        peopleCopy = people.copy()
        peopleCopy.remove(x)
        TravellingSalesman(graph,[x],peopleCopy)
    print(maximumHapiness) #Part 2

    