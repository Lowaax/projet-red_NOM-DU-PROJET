package main

import (
	"fmt"
	projet "projet_red_NOM-DU-PROJET/char"
)

func main() {
	c1 := initCharacter("Aragorn", "Elf", "Chevalier", 10, 100, 80, []string{"Épée", "Bouclier"})

	displayInfo(c1)
}

func initCharacter(name string, race string, class string, level int, maxHP int, hp int, inventory []string) projet.Character {
	return projet.Character{
		Name:      name,
		Race:      race,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		HP:        hp,
		Inventory: inventory,
	}
}

func displayInfo(c projet.Character) {
	fmt.Println("===== Informations du personnage =====")
	fmt.Println("Nom       :", c.Name)
	fmt.Println("Race      :", c.Race)
	fmt.Println("Classe    :", c.Class)
	fmt.Println("Niveau    :", c.Level)
	fmt.Println("PV        :", c.HP, "/", c.MaxHP)
	fmt.Println("Inventaire:", c.Inventory)
	fmt.Println("======================================")
}
