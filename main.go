package main

import (
	"fmt"
	projet "projet_red_NOM-DU-PROJET/char"
	"strings"
	"time"
)

func main() {
	c1 := characterCreation()
	Menu(c1)
	//displayInfo(c1)
	//accessInventory(c1)
	//takePot(c1)
	//displayInfo(c1)
	//accessInventory(c1)
}

func initCharacter(name, race, class string, maxHP int) *projet.Character {
	return &projet.Character{
		Name:         name,
		Race:         race,
		Class:        class,
		Level:        1,
		MaxHP:        maxHP,
		HP:           maxHP / 2,
		Inventory:    []string{"Potion", "Potion", "Potion"},
		Skills:       []string{"Coup de poing"},
		Gold:         100,
		MaxInventory: 10,
	}
}

func characterCreation() *projet.Character {
	var name, class string
	fmt.Print("Choisissez un nom : ")
	fmt.Scan(&name)
	name = strings.Title(strings.ToLower(name))

	fmt.Println("Choisissez une classe : (Humain, Elfe, Nain)")
	fmt.Scan(&class)

	var maxHP int
	switch strings.ToLower(class) {
	case "humain":
		maxHP = 100
	case "elfe":
		maxHP = 80
	case "nain":
		maxHP = 120
	default:
		maxHP = 100
		class = "Humain"
	}

	return initCharacter(name, "A définir", class, maxHP)
}

func displayInfo(c *projet.Character) {
	fmt.Println("===== Informations du personnage =====")
	fmt.Println("Nom       :", c.Name)
	fmt.Println("Race      :", c.Race)
	fmt.Println("Classe    :", c.Class)
	fmt.Println("Niveau    :", c.Level)
	fmt.Println("PV        :", c.HP, "/", c.MaxHP)
	fmt.Println("Or        :", c.Gold)
	fmt.Println("Inventaire:", c.Inventory)
	fmt.Println("Skills    :", c.Skills)
	fmt.Println("======================================")
}

func AccessInventory(c *projet.Character) {
	fmt.Println("===== Inventaire =====")
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")

	var choix int
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	if choix <= 0 || choix > len(c.Inventory) {
		return
	}

	item := c.Inventory[choix-1]
	switch item {
	case "Potion":
		takePot(c)
	case "Potion de poison":
		poisonPot(c)
	case "Livre de Sort : Boule de Feu":
		spellBook(c)
	default:
		fmt.Println("Rien ne se passe…")
	}
}

func takePot(c *projet.Character) {
	removeItem(c, "Potion")
	c.HP += 50
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	fmt.Printf("✅ Potion utilisée ! PV : %d/%d\n", c.HP, c.MaxHP)
}

func poisonPot(c *projet.Character) {
	removeItem(c, "Potion de poison")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.HP -= 10
		fmt.Printf("☠️ Poison tick %d → PV : %d/%d\n", i, c.HP, c.MaxHP)
	}
	if c.HP <= 0 {
		IsDead(c)
	}
}

func spellBook(c *projet.Character) {
	for _, s := range c.Skills {
		if s == "Boule de Feu" {
			fmt.Println("✨ Sort déjà appris.")
			return
		}
	}
	c.Skills = append(c.Skills, "Boule de Feu")
	removeItem(c, "Livre de Sort : Boule de Feu")
	fmt.Println("🔥 Nouveau sort appris : Boule de Feu !")
}

func addItem(c *projet.Character, item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println("❌ Inventaire plein !")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func removeItem(c *projet.Character, item string) {
	for i, it := range c.Inventory {
		if it == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
}

func Shop(c *projet.Character) {
	var choix int
	fmt.Println("===== Marchand =====")
	fmt.Println("1: Potion de vie (3 or)")
	fmt.Println("2: Potion de poison (6 or)")
	fmt.Println("3: Livre de Sort : Boule de Feu (25 or)")
	fmt.Println("4: Fourrure de Loup (4 or)")
	fmt.Println("5: Peau de Troll (7 or)")
	fmt.Println("6: Cuir de Sanglier (3 or)")
	fmt.Println("7: Plume de Corbeau (1 or)")
	fmt.Println("0: Retour")
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	type Offer struct {
		Name string
		Cost int
	}

	offers := map[int]Offer{
		1: {"Potion", 3},
		2: {"Potion de poison", 6},
		3: {"Livre de Sort : Boule de Feu", 25},
		4: {"Fourrure de Loup", 4},
		5: {"Peau de Troll", 7},
		6: {"Cuir de Sanglier", 3},
		7: {"Plume de Corbeau", 1},
	}

	if offer, ok := offers[choix]; ok {
		if c.Gold < offer.Cost {
			fmt.Println("❌ Pas assez d’or.")
			return
		}
		if addItem(c, offer.Name) {
			c.Gold -= offer.Cost
			fmt.Printf("✅ %s acheté pour %d or (reste %d).\n", offer.Name, offer.Cost, c.Gold)
		}
	}
}

func IsDead(c *projet.Character) bool {
	if c.HP <= 0 {
		fmt.Println("💀 Vous êtes mort.")
		c.HP = c.MaxHP / 2
		fmt.Printf("✨ Résurrection avec %d/%d PV.\n", c.HP, c.MaxHP)
		return true
	}
	return false
}

func Menu(c *projet.Character) {
	for {
		var choix int
		fmt.Println("===== Menu =====")
		fmt.Println("1. Infos personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			displayInfo(c)
		case 2:
			AccessInventory(c)
		case 3:
			Shop(c)
		case 4:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func Forgeron(c *projet.Character) {
	var choix int
	fmt.Println("===== Forgeron =====")
	fmt.Println("1 : Chapeau de l'aventurier")
	fmt.Println("2 : Tunique de l'aventurier")
	fmt.Println("3 : Bottes de l'aventurier")
	fmt.Println("0 : Retour")
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	if choix == 0 {
		return
	}

	type Material struct {
		Name     string
		Quantity int
	}

	type Item struct {
		Name      string
		Cost      int
		Materials []Material
	}

	items := map[int]Item{
		1: {
			Name: "Chapeau de l'aventurier",
			Cost: 5,
			Materials: []Material{
				{"Plume de Corbeau", 2},
				{"Cuir de Sanglier", 1},
			},
		},
		2: {
			Name: "Tunique de l'aventurier",
			Cost: 5,
			Materials: []Material{
				{"Fourrure de Loup", 2},
				{"Peau de Troll", 1},
			},
		},
		3: {
			Name: "Bottes de l'aventurier",
			Cost: 5,
			Materials: []Material{
				{"Cuir de Sanglier", 1},
				{"Fourrure de Loup", 1},
			},
		},
	}

	item, ok := items[choix]
	if !ok {
		fmt.Println("Choix invalide.")
		return
	}

	canCraft := true
	for _, mat := range item.Materials {
		count := 0
		for _, invItem := range c.Inventory {
			if invItem == mat.Name {
				count++
			}
		}
		if count < mat.Quantity {
			canCraft = false
			break
		}
	}

	if !canCraft {
		fmt.Println("Matériaux insuffisants.")
		return
	}

	if c.Gold < item.Cost {
		fmt.Println("Pas assez d’or.")
		return
	}

	for _, mat := range item.Materials {
		for i := 0; i < mat.Quantity; i++ {
			removeItem(c, mat.Name)
		}
	}

	c.Gold -= item.Cost
	addItem(c, item.Name)

	fmt.Printf("%s forgé pour %d or (reste %d).\n", item.Name, item.Cost, c.Gold)
}
