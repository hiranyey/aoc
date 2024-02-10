from enum import Enum


class OPERATION(Enum):
    AND = 1
    OR = 2
    LSHIFT = 3
    RSHIFT = 4
    NOT = 5
    ASSIGNMENT = 6


def processGates(line):
    lhs, rhs = line.strip().split("->")
    rhs = rhs.strip()
    if "AND" in lhs:
        ops = lhs.split(" ")
        return (rhs, OPERATION.AND, ops[0], ops[2])
    elif "OR" in lhs:
        ops = lhs.split(" ")
        return (rhs, OPERATION.OR, ops[0], ops[2])
    elif "LSHIFT" in lhs:
        ops = lhs.split(" ")
        return (rhs, OPERATION.LSHIFT, ops[0], ops[2])
    elif "RSHIFT" in lhs:
        ops = lhs.split(" ")
        return (rhs, OPERATION.RSHIFT, ops[0], ops[2])
    elif "NOT" in lhs:
        ops = lhs.split(" ")
        return (rhs, OPERATION.NOT, ops[1])
    else:
        ops = lhs.strip()
        return (rhs, OPERATION.ASSIGNMENT, ops)


def processAssignment(currentGate, value):
    try:
        return int(currentGate[2])
    except:
        if currentGate[2] in value:
            return value[currentGate[2]]
        else:
            return None


def processAND(currentGate, value):
    try:
        a = int(currentGate[2])
    except:
        a = value.get(currentGate[2])

    try:
        b = int(currentGate[3])
    except:
        b = value.get(currentGate[3])
    if a != None and b != None:
        return a & b
    else:
        return None


def processOR(currentGate, value):
    try:
        a = int(currentGate[2])
    except:
        a = value.get(currentGate[2])

    try:
        b = int(currentGate[3])
    except:
        b = value.get(currentGate[3])
    if a != None and b != None:
        return a | b
    else:
        return None


def processLSHIFT(currentGate, value):
    a = value.get(currentGate[2])
    b = int(currentGate[3])
    if a != None and b != None:
        return a << b
    else:
        return None


def processRSHIFT(currentGate, value):
    a = value.get(currentGate[2])
    b = int(currentGate[3])
    if a != None and b != None:
        return a >> b
    else:
        return None


def processNOT(currentGate, value):
    a = value.get(currentGate[2])
    if a != None:
        return ~a & 0xFFFF
    else:
        return None


def Part(values):
    with open("input.txt") as f:
        gates = []
        for line in f.readlines():
            gates.append(processGates(line))
        counter = 0
        while len(gates) > 0:
            currentGate = gates[counter]
            newValue = None
            if currentGate[1] == OPERATION.ASSIGNMENT:
                newValue = processAssignment(currentGate,values)
            elif currentGate[1] == OPERATION.AND:
                newValue = processAND(currentGate, values)
            elif currentGate[1] == OPERATION.OR:
                newValue = processOR(currentGate, values)
            elif currentGate[1] == OPERATION.LSHIFT:
                newValue = processLSHIFT(currentGate, values)
            elif currentGate[1] == OPERATION.RSHIFT:
                newValue = processRSHIFT(currentGate, values)
            elif currentGate[1] == OPERATION.NOT:
                newValue = processNOT(currentGate, values)

            if newValue is not None:
                values[currentGate[0]] = newValue
                del gates[counter]
            if len(gates) > 0:
                counter = (counter + 1) % len(gates)
            else:
                break
        return values['a']

print(Part({}))
print(Part({'b':Part({})}))

