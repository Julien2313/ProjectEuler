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

def p46():
  x = 3
  while True:
    y = 0
    while True:
      if isPrime(x - 2 * pow(y, 2)):
        break
      if (x - 2 * pow(y, 2)) < y:
        print x
        return
      y += 1
    x += 2
    
