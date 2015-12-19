import itertools

edges = open('day9.input', 'r').readlines()

EDGES = {}
CITIES = set()

for e in edges:
    city1, extra = e.split(" to ")
    city2, dist = extra.split(" = ")
    EDGES[(city1, city2)] = int(dist)
    EDGES[(city2, city1)] = int(dist)
    CITIES.add(city1)
    CITIES.add(city2)

longest = -1
path = []
for perm in itertools.permutations(list(CITIES)):
    dist = 0
    valid_path = True
    for i in xrange(len(perm)-1):
        edge = (perm[i], perm[i+1])
        if edge in EDGES:
            dist += EDGES[edge]
        else:
            valid_path = False
            break

    if valid_path and dist > longest:
        longest = dist
        path = perm

print path, longest

#('Tristram', 'Tambi', 'Snowdin', 'AlphaCentauri', 'Faerun', 'Arbre', 'Straylight', 'Norrath') 207