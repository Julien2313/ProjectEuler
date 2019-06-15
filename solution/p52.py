import itertools

def p52():
  x = 1
  while True:
    x2 = str(2 * x)
    x3 = str(3 * x)
    x4 = str(4 * x)
    x5 = str(5 * x)
    x6 = str(6 * x)
    permuts = list(map("".join, itertools.permutations(str(x))))
    if x2 in permuts:
      if x3 in permuts:
        if x4 in permuts:
          if x5 in permuts:
            if x6 in permuts:
              print x
              break
    x += 1
      
