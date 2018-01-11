import math


def solve():
    lines = [map(int, x.split(',')) for x in open("text.txt", 'r').read().strip().splitlines()]
    bigger = 0
    biggerLine = 0
    for line in lines:
        num = line[1] * math.log(line[0], 10)
        if num > bigger:
            bigger = num
            biggerLine = lines.index(line)

    print biggerLine + 1