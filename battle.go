package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Battle(p *Player, e *Enemy) {
	ShowText("A wild " + e.Name + " appears!\n")
	time.Sleep(1 * time.Second)
	ShowText("Get ready for battle, " + p.Name + "!")

	for {
		displayBattle(p, e)
		ShowText("Choose your move \n1 - Attack \n2 - Defend \n3 - Run \n")
		var action string
		_, err := fmt.Scanln(&action)
		if err != nil {
			return
		}
		pDamage := max(p.Stats.STR, p.Stats.INT, p.Stats.DEX)
		eDamage := e.Level * 2
		switch action {
		case "1":
			clearScreen()
			// chance for crits or misses
			src := rand.NewSource(time.Now().UnixNano())
			r := rand.New(src)
			randNum := r.Intn(10-1+1) + 1

			if randNum <= 1 {
				ShowText("You missed!")
				goto missed
			}
			if randNum >= 8 {
				ShowText("Critical hit!!")
				pDamage *= 2
			}
			displayBattle(p, e)
			e.Hp.curr -= pDamage
			ShowText("You attack " + e.Name + " for " + strconv.Itoa(pDamage) + "!\n")
			if e.checkIfAlive() {
				goto finished
			}
		missed:
			time.Sleep(time.Second)
			clearScreen()
			displayBattle(p, e)
			ShowText(e.Name + " attacks you for " + strconv.Itoa(eDamage) + "!\n")
			p.Hp.curr -= eDamage
			if p.checkIfAlive() {
				goto died
			}
		case "2":
			ShowText("You defend against " + e.Name + "'s attack!")
			ShowText(e.Name + " attacks you for " + strconv.Itoa(eDamage/2) + "!\n")
			p.Hp.curr -= eDamage / 2
			if p.checkIfAlive() {
				goto died
			}
		case "3":
			ShowText("You try to run from " + e.Name)
			ShowText("tsk tsk tsk.. cant do that!")
		default:
			fmt.Println("Invalid choice! Choose 1 to Attack, 2 to Defend, or 3 to Run.")
		}
		clearScreen()

	}
died:
	ShowText("Sorry... You died..")
	IntroText()
finished:
	displayBattle(p, e)
	time.Sleep(time.Second * 2)
	p.Hp.curr = p.Hp.max
	p.exp += e.Level*20 + 20
	p.checkXp()
	clearScreen()
}

func displayBattle(p *Player, e *Enemy) {
	playerInfo := fmt.Sprintf("%s (lvl %d)", p.Name, p.Level)
	playerStats := fmt.Sprintf("HP: %d/%d   EXP: %d", p.Hp.curr, p.Hp.max, p.exp)
	enemyInfo := e.Name
	enemyStats := fmt.Sprintf("HP: %d/%d", e.Hp.curr, e.Hp.max)

	// Calculate the border width based on the length of the longest line
	maxLineLength := len(playerInfo) * 2
	lines := []string{playerInfo, playerStats, enemyInfo, enemyStats}
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	border := "+" + strings.Repeat("-", maxLineLength) + "+"

	fmt.Println(border)
	fmt.Printf("| %-25s%-*s |\n", playerInfo, maxLineLength-len(enemyInfo)-4, enemyInfo)
	fmt.Printf("| %-25s%-*s |\n", playerStats, maxLineLength-len(playerStats), enemyStats)
	fmt.Println(border + "\n")
}
