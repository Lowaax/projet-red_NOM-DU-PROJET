package projet

var Races = []string{
	"Elf",
	"Nain",
	"Orc",
	"Human",
	"Dragon",
	"Mort-vivant",
	"Ange",
	"Orque",
	"Centaure",
	"Fée",
	"Lycanthrope",
	"Farfadet",
	"Antromorphe",
}

var Classes = []string{
	"Chevalier",
	"Sorcier",
	"Archer",
	"Assassin",
	"Prêtre",
	"Necromancien",
	"Berserker",
}

type Character struct {
	Name      string
	Race      string
	Class     string
	Level     int
	MaxHP     int
	HP        int
	Inventory []string
}
