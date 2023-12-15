package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	Warrior = iota + 1
	Rogue
	Mage
)

const (
	Fire = iota + 1
	Ice
	Lightning
)

type PlayerFactory struct{}

func NewPlayerFactory() *PlayerFactory {
	return &PlayerFactory{}
}

type PlayerPerkFactory struct{}

func NewPlayerClassFactory() *PlayerPerkFactory {
	return &PlayerPerkFactory{}
}

func (pf *PlayerFactory) CreatePlayer(name string, classType int, classPerk []string) *Player {
	var player Player
	playerBuilder(name, classType, classPerk, &player)
	return &player
}

func playerBuilder(name string, class int, classPerk []string, player *Player) *Player {
	player.Name = name
	player.Check.Stage = 1
	className, _ := player.displayClassToString(class, 0)

	switch class {
	case Warrior:
		player.Class.ClassPerk = classPerk
		player.Stats = Stats{STR: 5, INT: 3, DEX: 2}
		player.Hp.curr = 20
		player.Hp.max = 20
		player.Class.Name = className
		player.Level = 1
	case Rogue:
		player.Class.ClassPerk = classPerk
		player.Stats = Stats{STR: 2, INT: 3, DEX: 5}
		player.Hp.curr = 10
		player.Hp.max = 10
		player.Class.Name = className
		player.Level = 1
	case Mage:
		player.Class.ClassPerk = classPerk
		player.Stats = Stats{STR: 2, INT: 10, DEX: 3}
		player.Hp.curr = 5
		player.Hp.max = 5
		player.Class.Name = className
		player.Level = 1

	}

	return player
}

func (ppf *PlayerPerkFactory) perkBuilder(class int) []string {
	var classPerks []string

	switch class {
	case Warrior:
		warriorPerks := []string{"Unrelenting force", "Iron will", "Weapon specialist", "Battle expert", "Unyielding stance", "Heavy Strike", "Second wind"}
		classPerks = append(classPerks, warriorPerks...)
	case Rogue:
		roguePerks := []string{"Death blade", "Glass Cannon", "Quick attack"}
		classPerks = append(classPerks, roguePerks...)
	case Mage:
		magePerks := []string{"Highborn", "Glass Cannon", "Mana Shield"}
		classPerks = append(classPerks, magePerks...)
	default:

	}

	return classPerks
}

func (ppf *PlayerPerkFactory) getPerks(class, index int) []string {

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
	case Warrior:
		warriorPerks := []string{"Unrelenting force", "Iron will",
			"Weapon specialist", "Battle expert", "Unyielding stance",
			"Heavy Strike", "Second wind"}
		for i, v := range warriorPerks {
			fmt.Println(i+1, ": ", v)
		}
	case Rogue:
		roguePerks := []string{"Death blade", "Glass Cannon", "Quick attack"} // Example perks
		for i, v := range roguePerks {
			fmt.Println(i+1, ": ", v)
		}
	case Mage:
		magePerks := []string{"Highborn", "Glass Cannon", "Mana Shield"} // Example perks
		for i, v := range magePerks {
			fmt.Println(i+1, ": ", v)
		}
	default:

	}
}

func (ppf *PlayerPerkFactory) ChoosePerk(class int, perk int) []string {
	var startingSetOfPerks []string
	perks := ppf.getPerks(class, perk)
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

func CreateCharacter() *Player {
	scanner := bufio.NewScanner(os.Stdin)

	ShowText("What is your name?")
	scanner.Scan() // Read input
	name := scanner.Text()
	clearScreen()
	ShowText("Lets look around the room... i see \n" +
		"an old sword, a dusty spell book, and a lock picking kit with " +
		"a butter knife")
	ShowText("Which class do you want to be?\n1: Warrior\n2: Rogue\n3: Mage")
	scanner.Scan()
	classInput := scanner.Text()
	class, err := strconv.Atoi(classInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return nil
	}
	Cname := IntToStringClass(class)
	clearScreen()
	ShowText("Great choice... i would have chosen " + Cname + " aswell..")
	ShowText("To help you on your journey please choose a perk to help guide you.")
	perks := NewPlayerClassFactory()
	perks.ShowPerks(class)
	scanner.Scan()
	perkInput := scanner.Text()
	perk, err := strconv.Atoi(perkInput)
	if err != nil {
		fmt.Println("Try again with a real perk this time...")
		return nil
	}
	perk -= 1
	clearScreen()
	ps := perks.ChoosePerk(class, perk)
	pf := NewPlayerFactory()
	p := pf.CreatePlayer(name, class, ps)
	if class == 3 {
		p.mageSpellBuilderIntro()
	}
	strClass, strPerk := p.displayClassToString(class, perk)
	fmt.Printf("You have chosen the class %v and the perk %v\nGoodluck...\n", strClass, strPerk[0])

	return p
}

func (p *Player) displayClassToString(class, perk int) (string, []string) {
	ppf := PlayerPerkFactory{}

	switch class {
	case Warrior:
		return "Warrior", ppf.getPerks(class, perk)
	case Rogue:
		return "Rogue", ppf.getPerks(class, perk)
	case Mage:
		return "Mage", ppf.getPerks(class, perk)
	default:
		return "something went wrong...", nil
	}
}

func (p *Player) displayClassToInt(class string) int {
	switch class {
	case "Warrior":
		return Warrior
	case "Rogue":
		return Rogue
	case "Mage":
		return Mage
	default:
		fmt.Println("Not a class... Yet")
		return -1
	}
}

func (p *Player) mageSpellBuilderIntro() {
	ShowText("A magic user hey..\nWhat type of magic user are you?\n1: Fire\n2: Ice\n3: Lightning")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scan := scanner.Text()
	magicType, err := strconv.Atoi(scan)
	if err != nil {
		fmt.Println("Something went wrong..")
		p.mageSpellBuilderIntro()
	}
	clearScreen()
	tmp := Spells{}

	switch magicType {
	case Fire:
		tmp = Spells{
			Element:     "Fire",
			Name:        "Ember",
			Description: "A beginner's fire spell that's more likely to toast marshmallows than dragons.",
			Potency:     5,
			Uses:        30,
		}
		ShowText("Ah, the Fire Mage! Master of the barbecue, the bonfire, and the occasional accidental \n" +
			"eyebrow singe-off. You've always been hot-headed, both literally and metaphorically. The villagers \n" +
			"say you could start a fire in a rainstorm (and that one time, you actually did). Armed with your fiery \n" +
			"spells and an unquenchable thirst for adventure, you're ready to set Eldoria ablaze (hopefully, not \n" +
			"literally this time).")
	case Ice:
		tmp = Spells{
			Element:     "Ice",
			Name:        "Frosty Fingers",
			Description: "A beginner's ice magic that can barely freeze a cup of water, but hey, at least your drinks will never be warm again!",
			Potency:     5,
			Uses:        30,
		}
		ShowText("Choosing the Ice Mage, huh? Cool as a cucumber, cold as a freezer left open in winter. \n" +
			"You're the person villagers call when their drinks need chilling – stat. Your demeanor is so icy, \n" +
			"you've been accused of causing the last cold snap. But hey, with your frosty spells and knack for \n" +
			"'breaking the ice,' you're all set to give your enemies the cold shoulder as you skate gracefully \n" +
			"through Eldoria's challenges.")
	case Lightning:
		tmp = Spells{
			Element:     "Lightning",
			Name:        "Zap",
			Description: "A Beginner's attempt at lightning that's more likely to give a static shock than a thunderbolt.",
			Potency:     5,
			Uses:        30,
		}
		ShowText("Lightning Mage, the embodiment of shock and awe – and the occasional static shock. \n" +
			"You're so electrifying, villagers wonder if you're the reason the lights flicker. Fast-talking, \n" +
			"quick-thinking, and with reflexes so speedy you once dodged a falling apple, you're ready to charge \n" +
			"headfirst into adventure. Just remember to ground yourself now and then, or Eldoria might not have enough \n" +
			"hair products for all the static frizz.")
	}
	p.Spells = append(p.Spells, tmp)
	fmt.Println("You chose ", scanner.Text())

}

func (p *Player) checkXp() {
	if p.exp >= p.Level*20 {
		p.Level++
		LevelUp(p)
	} else {
		ShowText("You need " + strconv.Itoa(p.Level*20-p.exp) + " more exp to level up.")
	}
}

func LevelUp(p *Player) {
	ShowText("\n\n\nYou have leveled up!")
	time.Sleep(time.Second * 1)
	ShowText("\nYou are now level: " + strconv.Itoa(p.Level) + " Great work!")
	time.Sleep(time.Second * 1)
	ShowText("Please tell me which stat you would like to increase...\n" +
		"1. STR\n" +
		"2. INT\n" +
		"3. DEX")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	switch text {
	case "1":
		p.Stats.STR += 5
	case "2":
		p.Stats.INT += 5
	case "3":
		p.Stats.DEX += 5
	}
	p.Hp.max += 10
	p.Hp.curr = p.Hp.max
	time.Sleep(time.Second * 2)
	clearScreen()
}

func (p *Player) checkIfAlive() bool {
	if p.Hp.curr <= 0 {
		return true
	}
	return false
}

func IntToStringClass(class int) string {
	switch class {
	case Warrior:
		return "Warrior"
	case Rogue:
		return "Rogue"
	case Mage:
		return "Mage"
	default:
		return "something went wrong..."
	}
}
