def geneTriangle():
  n = 1
  while True:
    yield n*(n+1)/2
    n +=1
def genePentagonal():
  n = 1
  while True:
    yield n*(3*n-1)/2
    n +=1
  
def geneHexagonal():
  n = 1
  while True:
    yield n*(2*n-1)
    n +=1
    
def p45():
  GT = geneTriangle()
  GP = genePentagonal()
  GH = geneHexagonal()

  found = 0
  valueT = next(GT)
  valueP = next(GP)
  valueH = next(GH)

  while found != 3:
    if valueT == valueP and valueT == valueH:
      found += 1
      print valueT, valueP, valueH
      valueT = next(GT)
      valueP = next(GP)
      valueH = next(GH)

    if valueT <= valueP and valueT <= valueH:
      valueT = next(GT)
    elif valueP <= valueT and valueP <= valueH:
      valueP = next(GP)
    elif valueH <= valueP and valueH <= valueT:
      valueH = next(GH)
      
  print valueT, valueP, valueH






  
