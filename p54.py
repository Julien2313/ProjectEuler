def p54():
  maxi = 0
  for a in xrange(0, 100):
    for b in xrange(0, 100):
      num = a**b
      somme = sum(map(int, str(num)))
      if somme > maxi:
        maxi = somme
  print maxi
