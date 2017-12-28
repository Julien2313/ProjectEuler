def isPalindrome(string):
  for x in xrange(0, len(string)):
    char = string[x]
    if char != string[-x-1]:
      return False
  return True

def p36():
  sum = 0
  for num in xrange(0, 1000000):
    numBin = bin(num)
    if isPalindrome(str(num)) and isPalindrome(str(numBin)[2:]):
      sum += num
  print sum
