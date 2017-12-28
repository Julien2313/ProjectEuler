import math

def p34():
  sumTot = 0
  for num in xrange(3, 10000000):
    subSum = 0
    for digit in str(num):
      subSum += math.factorial(int(digit))
    if subSum == num:
      sumTot += subSum
      print "tot : ", sumTot, ", subTot : ", subSum
