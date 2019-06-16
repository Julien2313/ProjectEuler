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

def solve():
    nbrPrime = 0
    cpt = 3
    while True:
        num1 = cpt * cpt
        num2 = num1 - cpt + 1
        num3 = num2 - cpt + 1
        num4 = num3 - cpt + 1
        # print num1, num2,num3, num4
        if isPrime(num2):
            nbrPrime += 1
        if isPrime(num3):
            nbrPrime += 1
        if isPrime(num4):
            nbrPrime += 1
        if nbrPrime / (cpt * 2.0 - 1) < 0.1:
            print cpt
            break

        cpt += 2

