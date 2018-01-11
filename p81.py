size = 80 - 1


def deplacement(data, x=0, y=0, tot=0, cpt=0, mini=9999999999999999):
    if data[x][y] < 0:
        score = -data[x][y] + cpt
        x = y = size
    else:
        score = data[x][y] + cpt

    if score > mini:
        return mini

    if x < size:
        mini = deplacement(data, x + 1, y, tot, score, mini)
    if y < size:
        mini = deplacement(data, x, y + 1, tot, score, mini)
    if x == size and y == size:
        if score < mini:
            mini = score

    return mini


def solve():
    file = open("text.txt", 'r').read().splitlines()
    data = []
    for x in xrange(0, len(file)):
        data.append(map(int, file[x].split(',')))

    for x in xrange(0, size + 1):
        for y in xrange(0, size + 1):
            data[size - x][size - y] = -deplacement(data, size - x, size - y)
    print -data[0][0]

