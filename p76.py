def solve():
  tabWays = [0]*101
  tabWays[0] = 1
  for x in xrange(1, 100):
    for y in xrange(x, 101):
      tabWays[y] += tabWays[y-x]
  print tabWays[100]