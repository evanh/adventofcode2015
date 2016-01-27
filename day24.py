from itertools import combinations, chain, permutations

def split_list(data, n):
    for splits in combinations(range(1, len(data)), n-1):
        result = []
        prev = None
        for split in chain(splits, [None]):
            result.append(data[prev:split])
            prev = split
        yield result

presents = [1, 3, 5, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113]
# presents = [1, 2, 3, 4, 5, 7, 8, 9, 10, 11]

possibles = None
for perm in permutations(presents):
    for p in split_list(perm, 3):
        if sum(p[0]) == sum(p[1]) and sum(p[1]) == sum(p[2]):
            score = min(len(a) for a in p)
            if not possibles:
                possibles = [(p, score)]
                continue

            if score < possibles[0][1]:
                possibles = [(p, score)]
            elif score == possibles[0][1]:
                possibles.append((p, score))

possibles = map(lambda x: sorted(x, key=lambda y: len(y)), possibles)

smallest_qq = 100000000
grouping = None
for p in possibles:
    qq = reduce(lambda x, y: x * y, p[0])
    if qq < smallest_qq:
        smallest_qq = qq
        grouping = p

print grouping, qq
