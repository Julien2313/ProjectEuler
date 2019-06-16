
def solve():
	n =[1, 1, 2]
	for i in xrange(3, 100+2):
		if i%3 == 0:
			d = n[i-1] + n[i-2]
		elif i%3 == 1:
			d = 2*(i-1)*n[i-1]/3 + n[i-2]
		else:
			d = n[i-1] + n[i-2]
		n.append(d)

	s = 0
	while n[-1]:
		n[-1], remainder = divmod(n[-1], 10)
		s += remainder
	print s
