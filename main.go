package main

import (
	"fmt"
	projet "projet_red_NOM-DU-PROJET/char"
)

func main() {
	c1 := initCharacter("Aragorn", "Elf", "Chevalier", 10, 100, 40, []string{"Épée", "Potion", "Bouclier"})

	//displayInfo(c1)
	//accessInventory(c1)
	//takePot(c1)
	//displayInfo(c1)
	//accessInventory(c1)
	Menu(c1)
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

func AccessInventory(c projet.Character) {
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

func takePot(c projet.Character) {
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
	c.Inventory = append(c.Inventory[:found], c.Inventory[found+1:]...)
	c.HP += 50

	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	fmt.Printf("✅ Vous avez utilisé une potion ! PV : %d/%d\n", c.HP, c.MaxHP)
}

func Menu(c projet.Character) {
	for {
		var choix int
		fmt.Println("===== Menu =====")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Accéder au marchand")
		fmt.Println("4. Quitter")
		fmt.Println("================")
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			displayInfo(c)
		case 2:
			AccessInventory(c)
		case 3:
			Shop(c)
		case 4:
			fmt.Println(" Au revoir !")
			return
		default:
			fmt.Println("❌ Choix invalide. Veuillez réessayer.")
		}
		fmt.Println()
	}
}

func Shop(c projet.Character) {
	var choice_shop int
	fmt.Println("===== Marchand =====")
	fmt.Println("Marchand: 'Salut, gros fils de pute, \nqu'est ce qui t'intéresserait dans mon magasin de merde ?'")
	fmt.Println("===== Stock =====")
	fmt.Println("1: Potion de vie de fdp (Gratuit)")
	fmt.Println("0: Annuler")
	fmt.Println("======================")
	fmt.Println("Marchand: C'est tout ce que j'ai, sale merde de pigeon")
	fmt.Println("Vous avez choisi: ")
	fmt.Scan(&choice_shop)
	if choice_shop == 1 {
		fmt.Println("Marchand: Une potion de vie ? voila pour toi")
		c.Inventory = append(c.Inventory, "Potion")
	} else if choice_shop == 0 {
		fmt.Println("Marchand: Tu change d'avis ? bah dégage alors")
	} else {
		fmt.Println("T'es vraiment con toi, ", choice_shop, "n'est pas une option")
	}
}
