from math import hypot
from collections import defaultdict

def p39():
  maxiP = defaultdict(lambda: 0, {})
  
  for y in xrange(1, 1000):
      for x in xrange(y, 1000 - y):
        z=hypot(x,y)
        if y + x + z > 1000:
          break
        if int(z) != z:
          continue
        maxiP[int(y + x + z)] += 1
  print max(maxiP.iterkeys(), key=(lambda key: maxiP[key]))
