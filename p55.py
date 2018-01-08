def solve():
    cpt = 0
    for x in xrange(1, 10000):
      for step in xrange(0, 50):
        x += int(str(x)[::-1])
        if str(x) == str(x)[::-1]:
          break
      else:
        cpt += 1
    print cpt