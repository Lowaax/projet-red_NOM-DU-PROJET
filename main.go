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
		Skills:    []string{"Coup de poing"},
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
	fmt.Println("Skills    :", c.Skills)
	fmt.Println("======================================")
}

func AccessInventory(c projet.Character) {
	fmt.Println("===== Inventaire =====")
	if len(c.Inventory) == 0 {
		fmt.Println("L'inventaire est vide.")
		fmt.Println("======================")
		return
	}
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")
	fmt.Println("======================")
	fmt.Print("Choisissez un item à utiliser : ")

	var choix int
	fmt.Scan(&choix)

	if choix == 0 || choix < 0 || choix > len(c.Inventory) {
		return
	}

	item := c.Inventory[choix-1]

	switch item {
	case "Livre de Sort : Boule de Feu":
		fmt.Println("📖 Vous lisez le grimoire…")
		spellBook(c)
	case "Potion":
		takePot(c)
	default:
		fmt.Println("Rien ne se passe…")
	}
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

func Shop(c projet.Character) projet.Character {
	var choice_shop int
	fmt.Println("===== Marchand =====")
	fmt.Println("Marchand: Bienvenue jeune aventurier, \nqu'est ce qui t'intéresserait dans ma boutique ?")
	fmt.Println("===== Stock =====")
	fmt.Println("1: Potion de vie (Gratuit)")
	fmt.Println("2: Livre de Sort : Boule de Feu (Gratuit)") // <-- ajout
	fmt.Println("0: Annuler")
	fmt.Println("======================")
	fmt.Println("Marchand: C'est tout ce que j'ai, jeune aventurier")
	fmt.Print("Vous avez choisi: ")
	fmt.Scan(&choice_shop)

	if choice_shop == 1 {
		fmt.Println("Marchand: Une potion de vie ? voila pour toi")
		c.Inventory = append(c.Inventory, "Potion")
	} else if choice_shop == 2 {
		fmt.Println("Marchand: Un grimoire enflammé, prends soin de tes moustaches~")
		c.Inventory = append(c.Inventory, "Livre de Sort : Boule de Feu")
	} else if choice_shop == 0 {
		fmt.Println("Marchand: Tu change d'avis ? A bientôt alors !")
	} else {
		fmt.Println("Tu ne peux pas,", choice_shop, "n'est pas disponible")
	}
	return c
}

func IsDead(c projet.Character) bool {
	if c.HP <= 0 {
		fmt.Println("Le personnage est mort.")
		c.HP = c.MaxHP / 2
		fmt.Println("Le personnage a été ressuscité avec la moitié de ses PV max :", c.HP, "/", c.MaxHP)
		return true
	}
	return false
}

func spellBook(c projet.Character) {
	for _, s := range c.Skills {
		if s == "Boule de Feu" {
			fmt.Println("✨ Vous connaissez déjà ce sort.")
			return
		}
	}
	c.Skills = append(c.Skills, "Boule de Feu")
	fmt.Println("🔥 Nouveau sort appris : Boule de Feu !")
}
