def geneTriangle():
    n = 1
    while True:
        yield n * (n + 1) / 2
        n += 1

def geneSquare():
    n = 1
    while True:
        yield n*n
        n += 1

def genePentagonal():
    n = 1
    while True:
        yield n * (3 * n - 1) / 2
        n += 1

def geneHexagonal():
    n = 1
    while True:
        yield n * (2 * n - 1)
        n += 1

def geneHeptagonal():
    n = 1
    while True:
        yield n * (5 * n - 3) / 2
        n += 1

def geneOctogonal():
    n = 1
    while True:
        yield n * (3 * n - 2)
        n += 1

def solve():
    trig = geneTriangle()
    square = geneSquare()
    penta = genePentagonal()
    hexa = geneHexagonal()
    hepta = geneHeptagonal()
    octo = geneOctogonal()
    for x in xrange(0, 10):
        print next(octo)
    print "not yet, soon"