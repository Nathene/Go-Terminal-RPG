package main

import "fmt"

var enemies []Enemy

func CreateEnemies() {
	Ghoul := Enemy{
		hp:    10,
		level: 1,
		name:  "Ghoul",
		race:  "Necrophage",
		stats: Stats{
			STR: 2,
			INT: 1,
			DEX: 2,
		},
		spells: Spells{
			name:        "Heal",
			description: "Basic heal",
			potency:     5,
			uses:        1,
		},
	}

	enemies = append(enemies, Ghoul)

	fmt.Printf("%+v", enemies)
}
