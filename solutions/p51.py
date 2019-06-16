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
    cpt = 100001
    while True:
        if not isPrime(cpt):
            cpt += 2
            continue

        strNum = str(cpt)

        for maskBin in xrange(6, pow(2, len(strNum)) - 1):
            cptPrime = 10
            for digit in xrange(0, 10):
                mask = str(int(str(bin(maskBin))[2:]) * digit).zfill(len(strNum))
                newNum = ""
                for x in xrange(0, len(strNum)):
                    if mask[x] == '0':
                        newNum += strNum[x]
                    else:
                        newNum += mask[x]
                if len(str(int(newNum))) != len(strNum):
                    cptPrime -= 1
                    if cptPrime < 8:
                        break
                elif not isPrime(int(newNum)):
                    cptPrime -= 1
                    if cptPrime < 8:
                        break
                else:
                    print newNum,
            else:
                print
                print cpt, cptPrime, newNum, str(int(str(bin(maskBin))[2:])).zfill(len(strNum))
                return
            print

        print cpt
        cpt += 2
