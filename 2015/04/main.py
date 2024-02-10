import hashlib
import itertools

def mine(startingString):
    with open("input.txt") as f:
        data = f.readline()
        for i in itertools.count():
            if hashlib.md5(f'{data}{i}'.encode('utf-8')).hexdigest().startswith(startingString):
                return i

print(mine("00000"))
print(mine("000000"))