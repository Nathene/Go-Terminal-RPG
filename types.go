package main

//type Battle interface {
//	fight()
//}

type Stats struct {
	STR int
	INT int
	DEX int
}

type Spells struct {
	name        string
	description string
	potency     int
	uses        int
}

type Class struct {
	classPerk []string
}

type Enemy struct {
	hp     int
	level  int
	name   string
	race   string
	stats  Stats
	spells Spells
}

type Player struct {
	hp     int
	level  int
	name   string
	class  Class
	stats  Stats
	spells Spells
}
