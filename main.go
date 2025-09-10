package main

import (
	"fmt"
	projet "projet_red_NOM-DU-PROJET/char"
)

func main() {
	c1 := initCharacter("Aragorn", "Elf", "Chevalier", 10, 100, 40, []string{"Épée", "Potion", "Bouclier"})

	displayInfo(c1)
	accessInventory(c1)
	takePot(&c1)
	displayInfo(c1)
	accessInventory(c1)
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

func accessInventory(c projet.Character) {
	fmt.Println("===== Inventaire =====")
	if len(c.Inventory) == 0 {
		fmt.Println("L'inventaire est vide.")
	} else {
		for i, item := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
	fmt.Println("======================")
}

func takePot(c *projet.Character) {
	// Chercher une potion dans l'inventaire
	found := -1
	for i, item := range c.Inventory {
		if item == "Potion" {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Println("❌ Vous n'avez pas de potion dans l'inventaire.")
		return
	}

	// Retirer la potion de l'inventaire
	c.Inventory = append(c.Inventory[:found], c.Inventory[found+1:]...)

	// Soigner le personnage
	c.HP += 50
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}

	fmt.Printf("✅ Vous avez utilisé une potion ! PV : %d/%d\n", c.HP, c.MaxHP)
}

func Menu() {
	for {
		var choix int
		fmt.Println("===== Menu =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Quitter")
		fmt.Println("================")
		fmt.Print("Entrez votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println(" Informations du personnage :")
			fmt.Println("Nom: Héros, Niveau: 5, PV: 100/100")
		case 2:
			fmt.Println(" Inventaire :")
			fmt.Println("- Épée")
			fmt.Println("- Potion de soin")
		case 3:
			fmt.Println(" Au revoir !")
			return
		default:
			fmt.Println("❌ Choix invalide. Veuillez réessayer.")
		}
		fmt.Println()
	}
}
