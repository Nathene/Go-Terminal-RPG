package main

import "fmt"

var enemies []Enemy

func GetMeAnEnemy(level int) Enemy {
	CreateEnemies(level)
	return enemies[level-1]
}
func CreateEnemies(level int) []Enemy {
	switch level {
	case 1:

		Mouse := Enemy{
			Hp: Hp{
				curr: 18,
				max:  18,
			},
			Level: 1,
			Name:  "Mouse",
			Race:  "Rodent",
			Stats: Stats{
				STR: 1,
				INT: 1,
				DEX: 1,
			},
			Spells: Spells{
				Name:        "Heal",
				Description: "Basic heal",
				Potency:     5,
				Uses:        1,
			},
		}
		enemies = append(enemies, Mouse)
	case 2:

		Boar := Enemy{
			Hp: Hp{
				curr: 42,
				max:  42,
			},
			Level: 2,
			Name:  "Boar",
			Race:  "Pig",
			Stats: Stats{
				STR: 2,
				INT: 1,
				DEX: 2,
			},
			Spells: Spells{
				Name:        "Heal",
				Description: "Basic heal",
				Potency:     5,
				Uses:        1,
			},
		}
		enemies = append(enemies, Boar)
	case 3:
		BabyDragon := Enemy{
			Hp: Hp{
				curr: 74,
				max:  74,
			},
			Level: 4,
			Name:  "Baby Dragon",
			Race:  "Draconoid",
			Stats: Stats{
				STR: 5,
				INT: 2,
				DEX: 3,
			},
			Spells: Spells{
				Name:        "Heal",
				Description: "Basic heal",
				Potency:     5,
				Uses:        1,
			},
		}
		enemies = append(enemies, BabyDragon)
	default:
		Ghoul := Enemy{
			Hp: Hp{
				curr: 10,
				max:  10,
			},
			Level: 1,
			Name:  "Ghoul",
			Race:  "Necrophage",
			Stats: Stats{
				STR: 2,
				INT: 1,
				DEX: 2,
			},
			Spells: Spells{
				Name:        "Heal",
				Description: "Basic heal",
				Potency:     5,
				Uses:        1,
			},
		}

		enemies = append(enemies, Ghoul)

		fmt.Printf("%+v", enemies)

	}
	return nil
}

func (e *Enemy) checkIfAlive() bool {
	if e.Hp.curr <= 0 {
		return true
	}
	return false
}
