def isPrime(n):
  if n == 2 or n == 3: return True
  if n < 2 or n%2 == 0: return False
  if n < 9: return True
  if n%3 == 0: return False
  r = int(n**0.5)
  f = 5
  while f <= r:
    if n%f == 0: return False
    if n%(f+2) == 0: return False
    f +=6
  return True  
  
def p37():
  cptPrime = 0
  cpt = 10
  sum = 0
  while cptPrime < 11:
    if isPrime(cpt):
      for x in xrange(1, len(str(cpt))):
        stringCpt = str(cpt)
        if not isPrime(int(stringCpt[x:])):
          break
        if not isPrime(int(stringCpt[:-x])):
          break
      else:
        sum += cpt
        print cptPrime, sum, cpt
        cptPrime += 1
    cpt += 1
  print sum
    
