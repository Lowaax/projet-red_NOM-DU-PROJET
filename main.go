package main

import (
	"fmt"
)

func main() {
	c1 := initCharacter("Aragorn", "Elf", "Chevalier", 10, 100, 100, []string{"Épée", "Bouclier"})
	fmt.Println("Nom:", c1.Name)
	fmt.Println("Race:", c1.Race)
	fmt.Println("Classe:", c1.Class)
	fmt.Println("Niveau:", c1.Level)
	fmt.Println("PV Max:", c1.MaxHP)
	fmt.Println("PV Actuel:", c1.HP)
	fmt.Println("Inventaire:", c1.Inventory)
}

func initCharacter(name string, race string, class string, level int, maxHP int, hp int, inventory []string) Character {
	return Character{
		Name:      name,
		Race:      race,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		HP:        hp,
		Inventory: inventory,
	}
}
