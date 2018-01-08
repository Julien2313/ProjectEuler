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

def countDigits(num):
    digits = [0]*10
    for digit in num:
        digits[int(digit)] += 1
    return digits


def solve():
    for num in xrange(1000, 9999):
        if not isPrime(num):
            continue
        digits = countDigits(str(num))
        for add in xrange(1, 9999):
            num2 = num + add
            num3 = num2 + add
            if num2 > 9999 or num3> 9999:
                break
            if not isPrime(num2):
                continue
            if not isPrime(num3):
                continue
            if digits == countDigits(str(num2)):
                if digits == countDigits(str(num3)):
                    print "  ", num, num2, num3, str(num) + str(num2) + str(num3), add


