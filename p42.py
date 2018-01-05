def genTriangles():
  n = 1
  while True:
    yield n * (n + 1) / 2
    n += 1

def p42():
  
  file = open("word.txt", 'r').read()
  wordsSum = []
  cpt = 0
  GEN = genTriangles()
  for word in file.strip().split(','):
    wordsSum.append(sum([(ord(c) - ord('A') + 1) for c in word[1:-1]]))
    
  valueTrig = next(GEN)
  for wordSum in sorted(wordsSum):
    while valueTrig < wordSum:
      valueTrig = next(GEN)
    if valueTrig == wordSum:
      cpt += 1
  print cpt
    
