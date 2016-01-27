package main

import (
	"fmt"
)

type Game struct {
	MyHitPoints      int
	MonsterHitPoints int
	MonsterDamage    int
	Mana             int
	ManaSpent        int
	Shield           int
	Poison           int
	Recharge         int
	Turn             int
}

type Spell struct {
	Cost   int
	Damage int
	Health int
	Mana   int
	Armor  int
	Effect int
}

var MONSTER_DAMAGE = 9
var MIN_SPEND = 100000

// Cost, Damage to boss, health restored, mana restored, armor added,
// effect length
var SPELLS = map[string]Spell{
	"Magic Missile": Spell{53, 4, 0, 0, 0, 0},
	"Drain":         Spell{73, 2, 2, 0, 0, 0},
	"Shield":        Spell{113, 0, 0, 0, 7, 6},
	"Poison":        Spell{173, 3, 0, 0, 0, 6},
	"Recharge":      Spell{229, 4, 0, 101, 0, 5},
}

func Turn(game Game, spells []string) {
	if game.Turn == 0 {
		game.MyHitPoints -= 1
		if game.MyHitPoints <= 0 {
			return
		}
	}

	// Apply Effects
	if game.Shield > 0 {
		game.MonsterDamage = MONSTER_DAMAGE - SPELLS["Shield"].Armor
		game.Shield--
	} else {
		game.MonsterDamage = MONSTER_DAMAGE
	}

	if game.Poison > 0 {
		game.MonsterHitPoints -= SPELLS["Poison"].Damage
		game.Poison--
	}

	if game.Recharge > 0 {
		game.Mana += SPELLS["Recharge"].Mana
		game.Recharge--
	}

	// Possible for the boss to be dead at this point
	if game.MonsterHitPoints <= 0 {
		if game.ManaSpent < MIN_SPEND {
			MIN_SPEND = game.ManaSpent
			fmt.Println("EFFECTS", game.ManaSpent, spells)
		}
		return
	}

	if game.Turn == 1 {
		// Boss Turn
		game.MyHitPoints -= game.MonsterDamage
		if game.MyHitPoints <= 0 {
			return
		}
		game.Turn = 0
		Turn(game, spells)
		return
	}

	game.Turn = 1

	for name, spell := range SPELLS {
		if game.Mana < spell.Cost {
			continue
		}

		switch name {
		case "Magic Missile":
			game.MonsterHitPoints -= spell.Damage
			game.Mana -= spell.Cost
			game.ManaSpent += spell.Cost

			if game.MonsterHitPoints <= 0 {
				if game.ManaSpent < MIN_SPEND {
					fmt.Println("SPELL", game.ManaSpent, append(spells, name))
					MIN_SPEND = game.ManaSpent
				}
				game.MonsterHitPoints += spell.Damage
				game.Mana += spell.Cost
				game.ManaSpent -= spell.Cost
				continue
			}

			Turn(game, append(spells, name))
			game.MonsterHitPoints += spell.Damage
			game.Mana += spell.Cost
			game.ManaSpent -= spell.Cost
		case "Drain":
			game.MonsterHitPoints -= spell.Damage
			game.MyHitPoints += spell.Health
			game.Mana -= spell.Cost
			game.ManaSpent += spell.Cost

			if game.MonsterHitPoints <= 0 {
				if game.ManaSpent < MIN_SPEND {
					fmt.Println("SPELL", game.ManaSpent, append(spells, name))
					MIN_SPEND = game.ManaSpent
				}
				game.MonsterHitPoints += spell.Damage
				game.MyHitPoints -= spell.Health
				game.Mana += spell.Cost
				game.ManaSpent -= spell.Cost
				continue
			}

			Turn(game, append(spells, name))
			game.MonsterHitPoints += spell.Damage
			game.MyHitPoints -= spell.Health
			game.Mana += spell.Cost
			game.ManaSpent -= spell.Cost
		case "Shield":
			if game.Shield > 0 {
				continue
			}
			game.Mana -= spell.Cost
			game.ManaSpent += spell.Cost
			game.Shield = spell.Effect
			Turn(game, append(spells, name))
			game.Mana += spell.Cost
			game.ManaSpent -= spell.Cost
			game.Shield = 0
		case "Poison":
			if game.Poison > 0 {
				continue
			}
			game.Mana -= spell.Cost
			game.ManaSpent += spell.Cost
			game.Poison = spell.Effect
			Turn(game, append(spells, name))
			game.Mana += spell.Cost
			game.ManaSpent -= spell.Cost
			game.Poison = 0
		case "Recharge":
			if game.Recharge > 0 {
				continue
			}
			game.Mana -= spell.Cost
			game.ManaSpent += spell.Cost
			game.Recharge = spell.Effect
			Turn(game, append(spells, name))
			game.Mana += spell.Cost
			game.ManaSpent -= spell.Cost
			game.Recharge = 0
		}
	}
}

func main() {
	game := Game{
		MyHitPoints:      50,
		MonsterHitPoints: 58,
		MonsterDamage:    9,
		Mana:             500,
		ManaSpent:        0,
		Shield:           0,
		Poison:           0,
		Recharge:         0,
		Turn:             0,
	}
	Turn(game, []string{})
	fmt.Println(MIN_SPEND)
}
