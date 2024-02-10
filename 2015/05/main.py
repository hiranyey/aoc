def FirstNaughtyLine(line):
    numberOfVowels = 0
    foundDoubleCharacter = False
    foundMonotonousCharacter = False
    vowels = ['a', 'e', 'i', 'o', 'u']
    tempChar = '0'
    for c in line:
        if tempChar == c:
            foundDoubleCharacter = True
        if ord(c)-ord(tempChar)==1 and (tempChar == 'a' or tempChar =='c' or tempChar == 'p' or tempChar == 'x'):
            foundMonotonousCharacter = True
        if c in vowels:
            numberOfVowels+=1
        tempChar = c
    return numberOfVowels > 2 and foundDoubleCharacter and not foundMonotonousCharacter

def SecondNaughtyLine(line):
    valleyCheck = False
    for i in range(0,len(line)-2):
        if line[i]==line[i+2]:
            valleyCheck = True
    if not valleyCheck :
        return valleyCheck
    for i in range(0,len(line)-2):
        tempChar = line[i]
        tempChar2 = line[i+1]
        for j in range(i+2,len(line)-1):
            if tempChar == line[j] and tempChar2 == line[j+1]:
                return True
    return False

def Part(processor):
    with open("input.txt") as f:
        return len([d for d in f.readlines() if processor(d)])

print(Part(FirstNaughtyLine))
print(Part(SecondNaughtyLine))