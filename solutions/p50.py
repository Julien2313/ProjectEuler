def isPrime(n):
    if n == 2 or n == 3: return True
    if n < 2 or n % 2 == 0: return False
    if n < 9: return True
    if n % 3 == 0: return False
    r = int(n ** 0.5)
    f = 5
    while f <= r:
        if n % f == 0: return False
        if n % (f + 2) == 0: return False
        f += 6
    return True


def genPrimes():
    primes = [2, 3]
    yield 2
    yield 3
    n = 5
    while True:
        for prime in primes:
            if n % prime == 0:
                break
        else:
            primes.append(n)
            yield n
        n += 2


def p50():
    MAXI = 1000000
    GEN = genPrimes()
    primes = []
    cpt = 0
    maxiLen = 0
    while True:
        prime = next(GEN)
        if sum(primes[-maxiLen:]) > MAXI:
            break
        primes.append(prime)
        for offSet in xrange(0, len(primes)):
            somme = 0
            for x in xrange(offSet, len(primes)):
                somme += primes[x]
                if somme > MAXI:
                    break
                if isPrime(somme):
                    if maxiLen < x - offSet + 1:
                        maxiLen = x - offSet + 1
                        print somme, maxiLen

        cpt += 1