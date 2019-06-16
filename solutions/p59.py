def solve():
    chars = map(int, file("cypher.txt", 'r').read().strip().split(','))
    maxE = 0
    maxT = 0
    maxThe = 0
    maxc1 = 0
    maxc2 = 0
    maxc3 = 0
    for c1 in xrange(ord('a'), ord('z') + 1):
        for c2 in xrange(ord('a'), ord('z') + 1):
            for c3 in xrange(ord('a'), ord('z') + 1):
                sentence = ""
                for x in xrange(0, len(chars)):
                    if x % 3 == 0:
                        char = chr(chars[x] ^ c1)
                    elif x % 3 == 1:
                        char = chr(chars[x] ^ c2)
                    else:
                        char = chr(chars[x] ^ c3)
                    sentence += char
                sentence = sentence.lower()
                nbrE = sentence.count('e')
                nbrT = sentence.count('t')
                nbrThe = sentence.count('the')
                if nbrE > maxE and nbrT > maxT and nbrThe > maxThe:
                    print sentence[:20], c1, c2, c3
                    maxE = nbrE
                    maxT = nbrT
                    maxThe = nbrThe
                    maxc1 = c1
                    maxc2 = c2
                    maxc3 = c3

    sum = 0
    sentence = ""
    for x in xrange(0, len(chars)):
        if x % 3 == 0:
            char = chars[x] ^ maxc1
        elif x % 3 == 1:
            char = chars[x] ^ maxc2
        else:
            char = chars[x] ^ maxc3
        sum += char
        sentence += chr(char)
    print "Clear :"
    print sentence
    print "Sum :", sum
    print "key :", chr(maxc1) + chr(maxc2) + chr(maxc3)


