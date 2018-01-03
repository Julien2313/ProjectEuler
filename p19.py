import datetime

def p19():
	nbr = sum(1
		for year in xrange(1901, 2001)
		for month in xrange(1, 13)
		if datetime.date(year, month, 1).weekday() == 6) # 6 == sunday
	print str(nbr)
