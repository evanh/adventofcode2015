start = "3113322113"

def translate(num):
    output = ""
    curdig = ""
    for c in num:
        if curdig and c not in curdig:
            output += "%d%s" % (len(curdig), curdig[0])
            curdig = ""
        curdig += c

    output += "%d%s" % (len(curdig), curdig[0])
    curdig = ""

    return output

for i in xrange(50):
    start = translate(start)

print len(start)