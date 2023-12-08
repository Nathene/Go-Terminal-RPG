package main

import (
	"bufio"
	"fmt"
	"os"
)

// Constants for the classes
const (
	Warrior = iota + 1
	Rogue
	Mage
)

// Your struct definitions...

func IntroText() {
	fmt.Println("Hello and welcome!")
	fmt.Println("Would you like to: \n1. New Game\n2. Load Game\n3. Quit")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "1":
			NewGame()
			return
		case "2":
			LoadGame()
			return
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Please try again.. it's either 1, 2, or 3...")
		}
	}
}
