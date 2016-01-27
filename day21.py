import itertools
import math

WEAPONS = {
    "Dagger": (8, 4, 0),
    "Shortsword": (10, 5, 0),
    "Warhammer": (25, 6, 0),
    "Longsword": (40, 7, 0),
    "Greataxe": (74, 8, 0),
}

ARMOR = {
    "": (0, 0, 0),
    "Leather": (13, 0, 1),
    "Chainmail": (31, 0, 2),
    "Splintmail": (53, 0, 3),
    "Bandedmail": (75, 0, 4),
    "Platemail": (102, 0, 5),
}

RINGS = {
    "": (0, 0, 0),
    "": (0, 0, 0),
    "Damage +1": (25, 1, 0),
    "Damage +2": (50, 2, 0),
    "Damage +3": (100, 3, 0),
    "Defense +1": (20, 0, 1),
    "Defense +2": (40, 0, 2),
    "Defense +3": (80, 0, 3),
}

MONSTER = {
    "Hit Points": 104.0,
    "Damage": 8.0,
    "Armor": 1.0,
}

def defeat_monster(comb):
    w = comb[0]
    a = comb[1]
    rs = comb[2]

    damage = WEAPONS[w][1] + sum(RINGS[r][1] for r in rs)
    armor = ARMOR[a][2] + sum(RINGS[r][2] for r in rs)
    cost = WEAPONS[w][0] + ARMOR[a][0] + sum(RINGS[r][0] for r in rs)

    my_dps = max(damage - MONSTER["Armor"], 1)
    mo_dps = max(MONSTER["Damage"] - armor, 1)

    if mo_dps > my_dps:
        return False, cost

    my_rounds = math.ceil(MONSTER["Hit Points"] / my_dps)
    mo_rounds = math.ceil(100 / mo_dps)

    if my_rounds <= mo_rounds:
        return True, cost

    return False, cost

wa_combos = [[w, a] for w in WEAPONS for a in ARMOR]
combos = []

for r in itertools.combinations(RINGS, 2):
    for c in wa_combos:
        combos.append([c[0], c[1], r])

highest_cost = 0
loadout = None

for c in combos:
    success, cost = defeat_monster(c)
    if not success and cost > highest_cost:
        highest_cost = cost
        loadout = c

print loadout, highest_cost
