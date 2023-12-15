package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var state = 1

func NewGame() {
	var player *Player
	player = CreateCharacter()
	player.ShowPlayer()
	Stage1(player)

}

func LoadGame() (*GameState, error) {
	data, err := os.ReadFile(".gamestate.json")
	if err != nil {
		return nil, err
	}
	var state GameState
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func SaveGame(state GameState) error {

	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return os.WriteFile(".gamestate.json", data, 0644)
}

func Saving(p *Player, state int) {
	gs := GameState{Player: p, Level: state}
	err := SaveGame(gs)
	if err != nil {
		return
	}
}

func Stage1(p *Player) {
	Saving(p, state)

	fmt.Println(p.Name, p.Level, p.Class.Name)

	ShowText("\nThis is stage 1")

	state = 2
	Stage2(p)
}

func Stage2(p *Player) {
	Saving(p, state)
	stage2IntroBattles(p)
	Saving(p, state)
	stage2MidBattle(p)
	Saving(p, state)
	stage2LateBattle(p)
	state = 3
	Stage3(p)
}

func Stage3(p *Player) {

	ShowText("This is stage 3")
	fmt.Printf("Your name: %s, your class is %v, your level is %d\n", p.Name, p.Class.Name, p.Level)
	time.Sleep(10 * time.Second)

}

func LoadGameFile(gs *GameState) (*Player, int) {
	var p *Player
	p = gs.Player
	l := gs.Level

	switch l {
	case 1:
		Stage1(p)
	case 2:
		Stage2(p)
	case 3:
		Stage3(p)

	default:
		panic("error")
	}
	return p, 0
}

func ShowText(text string) {
	for _, runeValue := range text {
		fmt.Print(string(runeValue))
		time.Sleep(20 * time.Millisecond) // Adjust the duration for typing speed
	}
	fmt.Println()

}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin": // Linux and MacOS
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}
