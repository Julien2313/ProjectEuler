def solve():
    for x in xrange(1000000, 0, -1):
        if (x - 5) % 7 == 0:
            print ((x - 5) / 7 - 1) * 3
            break
