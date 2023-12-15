package main

type Stats struct {
	STR int
	INT int
	DEX int
}

type Spells struct {
	Element     string
	Name        string
	Description string
	Potency     int
	Uses        int
}

type Class struct {
	Name      string
	ClassPerk []string
}

type Enemy struct {
	Hp     Hp
	Level  int
	Name   string
	Race   string
	Stats  Stats
	Spells Spells
}

type Hp struct {
	curr int
	max  int
}

type Player struct {
	Hp     Hp
	Level  int
	exp    int
	Name   string
	Class  Class
	Stats  Stats
	Spells []Spells
	Check  Checkpoint
}

type Checkpoint struct {
	Stage int
}

type GameState struct {
	Player *Player
	Level  int
}
