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
	Name              string
	Race              string
	Class             string
	Level             int
	BaseMaxHP         int
	MaxHP             int
	HP                int
	Inventory         []string
	Skills            []string
	Gold              int
	MaxInventory      int
	InventoryUpgrades int
	Equip             Equipment
}

type Equipment struct {
	Head       string
	Chestplate string
	Feet       string
	Weapon     string
	Shield     string
	Accessory  string
}

type WeaponStats struct {
	Min int
	Max int
}

type ArmorStats struct {
	Slot    string // "Head", "Chestplate", "Feet"
	HPBonus int
}

var WeaponsDB = map[string]WeaponStats{
	"Coup de poing":  {Min: 3, Max: 5},
	"Épée courte":    {Min: 5, Max: 8},
	"Bâton usé":      {Min: 4, Max: 7},
	"Arc court":      {Min: 5, Max: 8},
	"Dague":          {Min: 4, Max: 9},
	"Masse légère":   {Min: 5, Max: 7},
	"Bâton d’os":     {Min: 4, Max: 8},
	"Hache rouillée": {Min: 6, Max: 9},
}

var ArmorsDB = map[string]ArmorStats{
	"Chapeau de l'aventurier": {Slot: "Head", HPBonus: 10},
	"Tunique de l'aventurier": {Slot: "Chestplate", HPBonus: 25},
	"Bottes de l'aventurier":  {Slot: "Feet", HPBonus: 15},

	// Stuff de départ (petits bonus)
	"Armure rembourrée": {Slot: "Chestplate", HPBonus: 15},
	"Robe simple":       {Slot: "Chestplate", HPBonus: 10},
	"Tunique légère":    {Slot: "Chestplate", HPBonus: 8},
	"Cape sombre":       {Slot: "Chestplate", HPBonus: 5},
	"Robe bénie":        {Slot: "Chestplate", HPBonus: 12},
	"Robe noire":        {Slot: "Chestplate", HPBonus: 10},
	"Bottes souples":    {Slot: "Feet", HPBonus: 5},
}
