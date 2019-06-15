def sumDigitsSquare(n):
    r = 0
    while n:
        m = n % 10
        n = n / 10
        r = r + m * m
    return r


def solve():
    tab = [0] * (567 + 1)
    cpt89 = 0

    for cpt in xrange(1, 10000000):
        tab[sumDigitsSquare(cpt)] += 1
    for x in xrange(1, 567 + 1):
        y = x
        while True:
            if y == 1:
                break
            if y == 89:
                cpt89 += tab[x]
                break
            y = sumDigitsSquare(y)
    print cpt89

