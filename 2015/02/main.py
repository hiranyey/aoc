def wrappingPaper(line):
    l,w,h = line.strip().split("x")
    length,width,height = int(l),int(w),int(h)
    return 2*length*width + 2*width*height + 2*height*length + min(length*width,width*height,height*length)

def ribbonPaper(line):
    l,w,h = line.strip().split("x")
    length,width,height = int(l),int(w),int(h)
    return length*width*height + 2*min(length+width,width+height,height+length)

def Part1():
    with open("input.txt") as f:
        return sum([wrappingPaper(x) for x in f.readlines()])
    
def Part2():
    with open("input.txt") as f:
        return sum([ribbonPaper(x) for x in f.readlines()])
    
print(Part1())
print(Part2())