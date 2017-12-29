def p46():
  sum = 0
  for x in xrange(1, 1000+1):
    sum = (sum + x**x)%10000000000
  print sum
