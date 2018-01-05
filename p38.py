from collections import Counter

def p38():
  maxi = 0
  #max 9999 because 1xxxx will always be smaller
  for num in xrange(0, 9999):
    if str(num).count('0') > 0:
      continue
    strNum = ''
    for n in xrange(1, 9):
      strNum += str(num * n)
      if len(strNum) < 9:
        continue
      if len(strNum) > 9:
        break
      if strNum.count('0') > 0:
        break
      
          
      occur     = Counter(strNum)
      maxOccur  = max(occur.iterkeys(), key=(lambda key: occur[key]))
      if occur[str(maxOccur)] > 1:
        break
      if n > 1:
        if maxi < strNum:
          maxi = strNum
        break
  print maxi

p38()
