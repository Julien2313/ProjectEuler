import math

def solve():
    cpt = 0
    num = 1
    while True:
        x = 1
        while True:
            if int(math.log10(pow(num, x))) + 1 == x:
                cpt += 1
            else:
                break
            x += 1
        if x == 1:
            break
        num += 1
    print cpt

