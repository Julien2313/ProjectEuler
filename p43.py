import itertools
def p43():
  sum = 0
  for num in itertools.permutations('9876543210'):
    num = ''.join(num)
    if int(num[1] + num[2] + num[3]) % 2 == 0:
      if int(num[2] + num[3] + num[4]) % 3 == 0:
        if int(num[3] + num[4] + num[5]) % 5 == 0:
          if int(num[4] + num[5] + num[6]) % 7 == 0:
            if int(num[5] + num[6] + num[7]) % 11 == 0:
              if int(num[6] + num[7] + num[8]) % 13 == 0:
                if int(num[7] + num[8] + num[9]) % 17 == 0:
                  sum += int(num)
                  print sum, num

  print sum
