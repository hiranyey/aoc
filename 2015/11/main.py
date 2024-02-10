def hasIncreasingStraight(word):
    for i in range(len(word) - 2):
        if ord(word[i]) + 1 == ord(word[i + 1]) and ord(word[i]) + 2 == ord(word[i + 2]):
            return True
    return False

def hasTwoPairs(word):
    pairs = 0
    i = 0
    while i < len(word) - 1:
        if word[i] == word[i + 1]:
            pairs += 1
            i += 2
        else:
            i += 1
    return pairs >= 2

def validPassword(word):
    if "i" in word or "o" in word or "l" in word:
        return False
    if not hasIncreasingStraight(word):
        return False
    if not hasTwoPairs(word):
        return False
    return True

def increment(word):
    word = list(word)
    i = len(word) - 1
    while i >= 0:
        if word[i] == "z":
            word[i] = "a"
            i -= 1
        else:
            word[i] = chr(ord(word[i]) + 1)
            break
    return "".join(word)

def Part(currentPassword):
    while not validPassword(currentPassword):
        currentPassword = increment(currentPassword)
    return currentPassword

with open("input.txt") as f:
    lines = f.readlines()
    currentPassword = lines[0].strip()
    print(Part(currentPassword))
    print(Part(increment(Part(currentPassword))))