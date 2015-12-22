import sys
from itertools import permutations

GUESTS = set(['Evan'])
RULES = {}

rules = open('day13.input', 'r').readlines()

for r in rules:
    parts = r.split()
    modifier = -1 if parts[2] == "lose" else 1
    happiness = int(parts[3])
    pair = (parts[0], parts[-1][:-1]) # Strip period
    GUESTS.add(pair[0])
    GUESTS.add(pair[1])
    RULES[pair] = happiness * modifier

def happiness(seating):
    total = 0
    l = len(seating)
    for i in xrange(len(seating)):
        left = seating[(i-1)%l]
        right = seating[(i+1)%l]
        middle = seating[i]
        total += RULES.get((middle, left), 0)
        total += RULES.get((middle, right), 0)
    return total

highest = -sys.maxint
best = None
for seating in permutations(list(GUESTS)):
    happy = happiness(seating)
    if happy > highest:
        highest = happy
        best = seating

print best, highest
