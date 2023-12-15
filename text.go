package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Your struct definitions...

func IntroText() {
	s1 := stage1Text{}
	ShowText("Hello and welcome!\n")
	ShowText("Would you like to: \n1. New Game\n2. Load Game\n3. Quit\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "1":
			clearScreen()
			ShowText(s1.NewGameIntroText())
			NewGame()
			return
		case "2":
			clearScreen()
			loadedState, err := LoadGame()
			if err != nil {
				ShowText("Error loading game..")
			}
			fmt.Printf("Loaded game state: %+v\n", loadedState)
			LoadGameFile(loadedState)
		case "3":
			os.Exit(0)
		default:
			ShowText("Please try again.. it's either 1, 2, or 3...")
		}
	}
}

type stage1Text struct{}

func (s1 *stage1Text) NewGameIntroText() string {
	return "Welcome, brave soul, to the land of Eldoria, a realm of wonder, whimsy, and untold\n" +
		"adventures! Your story begins in the quaint village of Sunnydale, known for its overly friendly\n" +
		"inhabitants and an uncanny number of sunny days per year. You are in your humble abode, a small but\n" +
		"cozy room filled with the bare necessities and a few personal trinkets. A poster on the wall boldly declares,\n" +
		"'Adventure Awaits!' and indeed, it does.\n"
}

func stage2IntroBattles(p *Player) {
	e := GetMeAnEnemy(p.Check.Stage)
	ShowText("\nAs you step out of your house, a rustling in the nearby bushes catches your attention. \n" +
		"You investigate, and suddenly, a mischievous mouse jumps out!\n\nMouse Battle:\nThe mouse eyes you \n" +
		"with a determined look, its tiny sword gleaming in the sunlight. It charges at you with surprising speed. \n" +
		"Time to show off your skills!")
	time.Sleep(time.Second * 2)
	clearScreen()
	Battle(p, &e)
	clearScreen()
	p.Check.Stage++
	ShowText("You have defeated the ferocious Mouse!")

	time.Sleep(1 * time.Second)
	ShowText("You stand victorious, catching your breath after the intense battle with the \n" +
		"surprisingly formidable mouse. As you gaze at the tiny sword it wielded, you can't help but\n" +
		"wonder if there's more to this world than meets the eye.")

}

func stage2MidBattle(p *Player) {
	e := GetMeAnEnemy(p.Check.Stage)
	ShowText("You decide to explore further and venture into the nearby woods. The trees, ancient and towering, \n" +
		"cast a comforting shade. Birds chirp melodiously, creating a symphony that accompanies your journey deeper \n" +
		"into the forest.")
	ShowText("Suddenly, the ground trembles beneath your feet. You brace yourself as a shadow looms over you. \n" +
		"It's a Giant Boar, its tusks sharp and menacing. It grunts aggressively, declaring its territory.")
	time.Sleep(time.Second * 2)
	clearScreen()
	Battle(p, &e)
	clearScreen()
	ShowText("With a final valiant effort, you defeat the Giant Boar! Your heart races with adrenaline as you \n" +
		"realize you're becoming quite the " + p.Class.Name + ".")
	ShowText("As you continue your journey, the forest starts to clear, revealing a quaint village. The villagers, \n" +
		"upon seeing you approach, begin to gather around. They've heard of your deeds, defeating the mischievous mouse \n" +
		"and the Giant Boar. They look at you with a mix of awe and hope.")
	p.Check.Stage++
}

func stage2LateBattle(p *Player) {
	e := GetMeAnEnemy(p.Check.Stage)
	ShowText("An elder steps forward, \"Brave adventurer, our village is in dire need of someone \n" +
		"with your skills. A dragon has been terrorizing us, and we are too weak to fight it. Will you help us?")
pathetic:
	ShowText("\n\nDo you: \n1. nod your head, feeling a sense of duty?\n or \n2. run away while you can...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()
	if answer != "1" {
		ShowText("Have you no shame...")
		goto pathetic
	}
	ShowText("\nDragon Quest:\nYou agree to help the villagers and set off towards the dragon's lair, guided \n" +
		"by their directions. The path is perilous, but your courage never falters.")
	ShowText("As you approach the dragon's lair, a sense of foreboding fills the air. The entrance is marked by \n" +
		"scorched earth and bones of fallen heroes. Taking a deep breath, you step inside, \n" +
		"ready to face your greatest challenge yet... or is it?")
	ShowText("\nStanding before you is a baby dragon, you steel yourself for the battle ahead. Despite its size, \n" +
		"you know the risk it poses to the village. 'This is a necessary challenge,' you remind yourself, \n" +
		"bracing for what's to come.")
	time.Sleep(time.Second * 2)
	clearScreen()
	Battle(p, &e)
	clearScreen()
	ShowText("\nWith the dragon defeated, you feel a mixture of relief and strength. As the sun sets, casting long \n" +
		"shadows over the village, you reflect on the day's events. 'What a day it has been,' you think, feeling the \n" +
		"weight of your actions but also recognizing the growth in your own strength and resolve.")
	p.Check.Stage++
}
