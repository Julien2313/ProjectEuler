import fractions

def p33():
	numerator = 1
	denominator = 1
	for dDenom in xrange(10, 100):
		for dNomi in xrange(10, dDenom):
			nD0 = dNomi % 10
			nD1 = dNomi // 10
			dD0 = dDenom % 10
			dD1 = dDenom // 10
			
			if (nD1 == dD0 and nD0 * dDenom == dNomi * dD1) or (nD0 == dD1 and nD1 * dDenom == dNomi * dD0):
				numerator   *= dNomi
				denominator *= dDenom
	print str(denominator / fractions.gcd(numerator, denominator))

