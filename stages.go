package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func NewGame() {
	var player *Player
	player = CreateCharacter()
	player.ShowPlayer()
}

func CreateCharacter() *Player {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What is your name?")
	scanner.Scan() // Read input
	name := scanner.Text()

	fmt.Println("Before we continue, please choose a class...")
	fmt.Println("1: Warrior\n2: Rogue\n3: Mage")
	scanner.Scan() // Read input
	classInput := scanner.Text()

	// Convert string to int
	class, err := strconv.Atoi(classInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return nil
	}
	fmt.Println("Great choice...")
	fmt.Println("To help you on your journey please choose a perk to help guide you.")
	perks := NewPlayerClassFactory()
	perks.ShowPerks(class)
	scanner.Scan() // Read input
	perkInput := scanner.Text()
	perk, err := strconv.Atoi(perkInput)
	if err != nil {
		fmt.Println("Try again with a real perk this time...")
		return nil
	}
	perk -= 1

	ps := perks.ChoosePerk(perk, class)
	pf := NewPlayerFactory()
	p := pf.CreatePlayer(name, class, ps)
	// Here you can do something with the 'player' object
	return p
}

func LoadGame() {

}
