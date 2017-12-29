import itertools

def p53():
  tot = 0
  chaine = "o" * 23
  while len(chaine) <= 100:
    print len(chaine),"/100 => ", tot
    for x in xrange(1, len(chaine)):
      lenPermuts = 0 
      for y in itertools.combinations(chaine, x):
        lenPermuts += 1
        if lenPermuts > 1000000:
          break
      if lenPermuts > 1000000:
        tot += (len(chaine) + 1) - x*2
        break
    chaine += 'o'
  print tot
