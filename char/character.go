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
	"Vampire",
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
	Name              string
	Race              string
	Class             string
	Level             int
	BaseMaxHP         int
	MaxHP             int
	HP                int
	MaxMana           int
	Mana              int
	Initiative        int
	Exp               int
	ExpMax            int
	Inventory         []string
	MaxInventory      int
	InventoryUpgrades int
	Gold              int
	Skills            []string

	Equip struct {
		Tête       string
		Torse string
		Pieds       string
		Arme     string
		Talisman  string
		Gants      string
		Ceinture	string 
	}
}
type Equipment struct {
	Tête       string
	Torse string
	Gants      string
	Ceinture	string 
	Pieds       string
	Arme     string
	Talisman  string
}

type ArmeStats struct {
	Min int
	Max int
}

type ArmureStats struct {
	Slot    string // "Tête", "Torse", "Pieds"
	HPBonus int
}

var ArmeDB = map[string]struct{ Min, Max int }{
	"Coup de poing":  {Min: 3, Max: 5},
	"Épée courte":    {Min: 5, Max: 8},
	"Bâton usé":      {Min: 4, Max: 7},
	"Arc court":      {Min: 5, Max: 8},
	"Dague":          {Min: 4, Max: 9},
	"Masse légère":   {Min: 5, Max: 7},
	"Bâton d’os":     {Min: 4, Max: 8},
	"Hache rouillée": {Min: 6, Max: 9},
}

var ArmureDB = map[string]struct {
	Slot    string // "Tête" | "Torse" | "Pieds"
	HPBonus int
}{
	"Chapeau de l'aventurier": {Slot: "Tête", HPBonus: 10},
	"Tunique de l'aventurier": {Slot: "Torse", HPBonus: 25},
	"Bottes de l'aventurier":  {Slot: "Pieds", HPBonus: 15},
	"Armure rembourrée": {Slot: "Torse", HPBonus: 15},
	"Robe simple":       {Slot: "Torse", HPBonus: 10},
	"Tunique légère":    {Slot: "Torse", HPBonus: 8},
	"Cape sombre":       {Slot: "Torse", HPBonus: 5},
	"Robe bénie":        {Slot: "Torse", HPBonus: 12},
	"Robe noire":        {Slot: "Torse", HPBonus: 10},
	"Bottes souples":    {Slot: "Pieds", HPBonus: 5},
}
