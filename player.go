package main

import "fmt"

// Player and other struct definitions from your previous code

// PlayerFactory is responsible for creating new Player instances
type PlayerFactory struct{}

// NewPlayerFactory creates a new instance of PlayerFactory
func NewPlayerFactory() *PlayerFactory {
	return &PlayerFactory{}
}

type PlayerPerkFactory struct{}

func NewPlayerClassFactory() *PlayerPerkFactory {
	return &PlayerPerkFactory{}
}

// CreatePlayer creates a new player with the given name and class type
func (pf *PlayerFactory) CreatePlayer(name string, classType int, classPerk []string) *Player {
	var player Player
	playerBuilder(name, classType, classPerk, &player)
	return &player
}

// class specific spells
func playerBuilder(name string, class int, classPerk []string, player *Player) *Player {
	player.name = name

	switch class {
	case Warrior:
		classType := Class{classPerk: classPerk}
		player.stats = Stats{STR: 10, INT: 3, DEX: 5}
		player.hp = 20
		player.class = classType
		player.level = 1
	case Rogue:

	}

	return player
}

func (ppf *PlayerPerkFactory) perkBuilder(class int) []string {
	var classPerks []string

	switch class {
	case 1: // Assuming 1 is Warrior
		warriorPerks := []string{"Unrelenting force", "Iron will", "Weapon specialist", "Battle expert", "Unyielding stance", "Heavy Strike", "Second wind"}
		classPerks = append(classPerks, warriorPerks...)
	case 2: // Assuming 2 is Rogue
		roguePerks := []string{"Death blade", "Glass Cannon", "Quick attack"} // Example perks
		classPerks = append(classPerks, roguePerks...)
	case 3: // Assuming 3 is Mage
		magePerks := []string{"Highborn", "Glass Cannon", "Mana Shield"} // Example perks
		classPerks = append(classPerks, magePerks...)
	default:
		// Handle unknown class
	}

	return classPerks
}

func (ppf *PlayerPerkFactory) getPerks(index int, class int) []string {

	switch class {
	case Warrior:
		warriorPerks := []string{"Unrelenting force", "Iron will",
			"Weapon specialist", "Battle expert", "Unyielding stance",
			"Heavy Strike", "Second wind"}

		return []string{warriorPerks[index]}
	case Rogue:
		roguePerks := []string{"Death blade", "Glass Cannon", "Quick attack"}
		return []string{roguePerks[index]}
	case Mage:
		magePerks := []string{"Highborn", "Glass Cannon", "Mana Shield"}
		return []string{magePerks[index]}
	}
	return nil
}

func (ppf *PlayerPerkFactory) ShowPerks(class int) {
	switch class {
	case 1: // Assuming 1 is Warrior
		warriorPerks := []string{"Unrelenting force", "Iron will",
			"Weapon specialist", "Battle expert", "Unyielding stance",
			"Heavy Strike", "Second wind"}
		for i, v := range warriorPerks {
			fmt.Println(i+1, ": ", v)
		}
	case 2: // Assuming 2 is Rogue
		roguePerks := []string{"Death blade", "Glass Cannon", "Quick attack"} // Example perks
		for i, v := range roguePerks {
			fmt.Println(i+1, ": ", v)
		}
	case 3: // Assuming 3 is Mage
		magePerks := []string{"Highborn", "Glass Cannon", "Mana Shield"} // Example perks
		for i, v := range magePerks {
			fmt.Println(i+1, ": ", v)
		}
	default:
		// Handle unknown class
	}
}

func (ppf *PlayerPerkFactory) ChoosePerk(perk int, class int) []string {
	var startingSetOfPerks []string
	perks := ppf.getPerks(perk, class) // Correct order of arguments
	if perks == nil {
		fmt.Println("Something went very wrong...")
		return nil
	}
	startingSetOfPerks = append(startingSetOfPerks, perks...)
	return startingSetOfPerks
}

func (p *Player) ShowPlayer() {
	fmt.Printf("%+v", p)
}
