globalMax = 0
def createDistribution(list, remaining, components,checkIf500):
    if len(list) == len(components):
        if sum(list) == 100:
            if checkIf500:
                cookieCalorie = 0
                sumOfComponent = 0
                for i,x in enumerate(list):
                    sumOfComponent += x * components[i][len(components[0])-1]
                sumOfComponent = max(0, sumOfComponent)
                if sumOfComponent !=500:
                    return
            product = 1
            for j in range(0, len(components[0])-1):
                sumOfComponent = 0
                for i,x in enumerate(list):
                    sumOfComponent += x * components[i][j]
                sumOfComponent = max(0, sumOfComponent)
                product *= sumOfComponent
            global globalMax
            globalMax = max(product, globalMax)
        return
    for i in range(0, remaining):
        listCopy = list.copy()
        listCopy.append(i)
        createDistribution(listCopy, remaining - i, components,checkIf500)


with open("input.txt") as f:
    components = []
    for d in f.readlines():
        x = d.strip().split(":")[1].split(",")
        ingredients = []
        for i, properties in enumerate(x):
            ingredients.append(int(properties.split(" ")[-1]))
        components.append(ingredients)
    for i in range(0, 101):
        createDistribution([i], 101 - i, components,False)
    print(globalMax)
    globalMax = 0 
    for i in range(0, 101):
        createDistribution([i], 101 - i, components,True)
    print(globalMax)
