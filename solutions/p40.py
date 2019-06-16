def p40():
  num = ""
  x = 0
  while len(num) <= 1000000:
    num += str(x)
    x += 1
  print int(num[1]) * int(num[10]) * int(num[100]) * int(num[1000]) * int(num[10000]) * int(num[100000]) * int(num[1000000])
