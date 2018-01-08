import math

def solve():
    lastNum = [1, 3]
    lastDenum = [1, 2]
    num = 3
    denum = 2
    cpt = 0
    for x in xrange(0, 1000):
        num = num * 2 + lastNum[-2]
        denum = denum * 2 + lastDenum[-2]

        lastNum.append(num)
        lastDenum.append(denum)

        lastNum = lastNum[1:]
        lastDenum = lastDenum[1:]

        if int(math.log(num, 10)) > int(math.log(denum, 10)):
            cpt += 1
    print cpt