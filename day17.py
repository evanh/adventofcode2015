from itertools import permutations

containers = open('day17.input', 'r').readlines()

containers = [int(c) for c in containers]

combos = 0

for i in xrange(len(containers)+1):
    for p in permutations(containers, i):
        if sum(p) == 150:
            combos += 1

print combos