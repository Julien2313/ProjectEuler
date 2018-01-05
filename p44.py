def pentagonal(n):
  return n * (3 * n - 1) / 2
    
def isPentagonal(N):
  # N = n * (3 * n - 1) / 2
  # 3n^2 - n - 2N = 0
  # roots :
  # n = (1+-sqrt(1+24N))/6
  # n must be positive, there is no solution for n = (1-sqrt(1+24N))/6
  # n is natural so :
  # (1+sqrt(1+24N)) % 6 = 0
  # (1   +  sqrt(1+24N))% 6   = 0
  # (       sqrt(1+24N))% 6   = 5
  return (pow(1+24*N, 1.0/2) % 6) == 5

def p44():
  
  for n1 in xrange(1, 10000):
    p1 = pentagonal(n1)
    for n2 in xrange(1, n1):
      p2 = pentagonal(n2)
      if isPentagonal(p1 - p2) and isPentagonal(p1 + p2):
        print abs(p1 - p2)
        return
  
