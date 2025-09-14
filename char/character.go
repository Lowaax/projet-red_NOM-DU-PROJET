package projet

type Equipment struct {
	Head       string
	Chestplate string
	Feet       string
}

type Character struct {
	Name         string
	Race         string
	Class        string
	Level        int
	MaxHP        int
	HP           int
	Inventory    map[string]int
	Skills       []string
	Gold         int
	MaxInventory int
	Equip        Equipment
}

type Monster struct {
	Name   string
	MaxHP  int
	HP     int
	Attack int
}

func InitGoblin() Monster {
	return Monster{
		Name:   "Gobelin d'entraînement",
		MaxHP:  40,
		HP:     40,
		Attack: 5,
	}
}
