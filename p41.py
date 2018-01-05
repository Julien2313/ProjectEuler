import itertools

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

def p41():
  chaine = '987654321'
  allNum = []
  for x in xrange(1, len('987654321')):
    for y in itertools.permutations(chaine, len(chaine)):
      if isPrime(int(''.join(y))):
        allNum.append(int(''.join(y)))
    if len(allNum) > 0:
      print max(allNum)
      break
    chaine = chaine[1:]
  
