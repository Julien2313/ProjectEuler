from collections import defaultdict
import math

def primeFactors(n):
  factors = defaultdict(lambda: 0, {})
  
  while n % 2 == 0:
    factors[2] += 1
    n = n / 2
    
  for divisor in xrange(3, int(math.sqrt(n)) + 1, 2):
    while n % divisor == 0:
      factors[divisor] += 1
      n = n / divisor
      
  if n > 2:
      factors[n] += 1
  return factors
  

def p47():
  x = 0
  while True:
    x+=1
    lenX      = len(primeFactors(x))
    if lenX != 4:
      continue
    lenXPlus1 = len(primeFactors(x+1))
    if lenXPlus1 != 4:
      continue
    lenXPlus2 = len(primeFactors(x+2))
    if lenXPlus2 != 4:
      continue
    lenXPlus3 = len(primeFactors(x+3))
    if lenXPlus3 != 4:
      continue
    print x
    break
