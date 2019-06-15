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
  
def p35():
  cpt = 0
  for num in xrange(1, 1000000+1):
    num = str(num)
    for x in xrange(0, len(num)):
      if not isPrime(int(num[x:] + num[:x])):
        break
    else:
      cpt += 1
      
  print cpt
      
