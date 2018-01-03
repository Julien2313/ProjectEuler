import itertools

def p32():
  products = []
  sumProducts = 0
  for y in itertools.permutations('123456789'):
    chaine = ''.join(y)
    for indiceMultiplicand in xrange(0, len(chaine) - 2):
      for indiceMultiplier in xrange(indiceMultiplicand + 1, len(chaine) - 1):
        multiplicand  = int(chaine[:indiceMultiplicand + 1])
        multiplier    = int(chaine[indiceMultiplicand + 1:indiceMultiplier + 1])
        product       = int(chaine[-(len(chaine) - indiceMultiplier - 1):])
        if multiplicand * multiplier == product:
          if not (product in products):
            sumProducts += product
            products.append(product)
    if y[0] == '5':
      break
    
  print sumProducts
