def solve():
  tot = 1
  for x in xrange(0, 7830457):
    tot *= 2
    tot %= 10000000000
  print (tot * 28433 + 1) % 10000000000