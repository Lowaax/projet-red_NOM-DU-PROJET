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
		Inventory:    map[string]int{"Potion": 3},
		Skills:       []string{"Coup de poing"},
		Gold:         100,
		MaxInventory: 10,
		Equip:        projet.Equipment{},
	}
}

func characterCreation() *projet.Character {
	var name, race string
	fmt.Print("Choisissez un nom : ")
	fmt.Scan(&name)
	name = strings.Title(strings.ToLower(name))

	fmt.Println("Choisissez une race (Humain, Elfe, Nain) :")
	fmt.Scan(&race)

	var maxHP int
	switch strings.ToLower(race) {
	case "humain":
		maxHP = 100
		race = "Humain"
	case "elfe":
		maxHP = 80
		race = "Elfe"
	case "nain":
		maxHP = 120
		race = "Nain"
	default:
		maxHP = 100
		race = "Humain"
	}

	fmt.Println("Choisissez une classe : (Chevalier, Sorcier, Archer, Assassin, Prêtre, Necromancien, Berserker)")
	var class string
	fmt.Scan(&class)
	class = strings.Title(strings.ToLower(class))

	return initCharacter(name, race, class, maxHP)
}

func displayInfo(c *projet.Character) {
	fmt.Println("===== Informations du personnage =====")
	fmt.Println("Nom       :", c.Name)
	fmt.Println("Race      :", c.Race)
	fmt.Println("Classe    :", c.Class)
	fmt.Println("Niveau    :", c.Level)
	fmt.Println("PV        :", c.HP, "/", c.MaxHP)
	fmt.Printf("Équipé    : Tête[%s] Torse[%s] Pieds[%s]\n", c.Equip.Head, c.Equip.Chestplate, c.Equip.Feet)
	fmt.Println("Or        :", c.Gold)

	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire: vide")
	} else {
		fmt.Print("Inventaire: ")
		for item, qty := range c.Inventory {
			fmt.Printf("%s (x%d), ", item, qty)
		}
		fmt.Println()
	}

	if len(c.Skills) == 0 {
		fmt.Println("Skills    : Aucun")
	} else {
		fmt.Println("Skills    :", strings.Join(c.Skills, ", "))
	}
	fmt.Println("======================================")
}

func addItem(c *projet.Character, item string) bool {
	count := 0
	for _, qty := range c.Inventory {
		count += qty
	}
	if count >= c.MaxInventory {
		fmt.Println("❌ Inventaire plein !")
		return false
	}

	c.Inventory[item]++
	return true
}

func removeItem(c *projet.Character, item string) bool {
	qty, ok := c.Inventory[item]
	if !ok || qty <= 0 {
		return false
	}

	if qty == 1 {
		delete(c.Inventory, item)
	} else {
		c.Inventory[item]--
	}
	return true
}

// La fonction AccessInventory adaptée à map[string]int
func AccessInventory(c *projet.Character) {
	fmt.Println("===== Inventaire =====")
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	i := 1
	items := []string{}
	for item, qty := range c.Inventory {
		fmt.Printf("%d. %s (x%d)\n", i, item, qty)
		items = append(items, item)
		i++
	}
	fmt.Println("0. Retour")

	var choix int
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	if choix == 0 || choix > len(items) {
		return
	}

	item := items[choix-1]

	switch item {
	case "Potion":
		if removeItem(c, "Potion") {
			takePot(c)
		}
	case "Potion de poison":
		if removeItem(c, "Potion de poison") {
			poisonPot(c)
		}
	case "Livre de Sort : Boule de Feu":
		if removeItem(c, "Livre de Sort : Boule de Feu") {
			spellBook(c)
		}
	default:
		fmt.Println("Rien ne se passe…")
	}
}

func takePot(c *projet.Character) {
	c.HP += 50
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	fmt.Printf("✅ Potion utilisée ! PV : %d/%d\n", c.HP, c.MaxHP)
}

func poisonPot(c *projet.Character) {
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("☠️ Poison tick %d → PV : %d/%d\n", i, c.HP, c.MaxHP)
		if c.HP == 0 {
			IsDead(c)
			break
		}
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
	fmt.Println("🔥 Nouveau sort appris : Boule de Feu !")
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
		fmt.Println("4. Forgeron")
		fmt.Println("5. Combattre un Gobelin")
		fmt.Println("6. Combat d'entraînement")
		fmt.Println("7. Équipement")
		fmt.Println("8. Quitter")
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
			Forgeron(c)
		case 5:
			Equipement(c)
		case 6:
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
	fmt.Println("1 : Fabriquer Chapeau de l'aventurier (1 Plume de Corbeau + 1 Cuir de Sanglier + 5 or)")
	fmt.Println("2 : Fabriquer Tunique de l'aventurier (2 Fourrure de Loup + 1 Peau de Troll + 5 or)")
	fmt.Println("3 : Fabriquer Bottes de l'aventurier (1 Cuir de Sanglier + 1 Fourrure de Loup + 5 or)")
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
				{"Plume de Corbeau", 1},
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

	CraftItem := func(it Item) {
		if c.Gold < it.Cost {
			fmt.Println("❌ Pas assez d’or.")
			return
		}
		if len(c.Inventory) >= c.MaxInventory {
			fmt.Println("❌ Inventaire plein, impossible d’ajouter l’équipement.")
			return
		}

		missing := false
		for _, mat := range it.Materials {
			count := c.Inventory[mat.Name]
			if count < mat.Quantity {
				fmt.Printf("❌ Il manque %d × %s\n", mat.Quantity-count, mat.Name)
				missing = true
			}
		}
		if missing {
			return
		}

		for _, mat := range it.Materials {
			c.Inventory[mat.Name] -= mat.Quantity
			if c.Inventory[mat.Name] <= 0 {
				delete(c.Inventory, mat.Name)
			}
		}

		c.Gold -= it.Cost

		c.Inventory[it.Name]++
		fmt.Printf("✅ %s forgé pour %d or (reste %d).\n", it.Name, it.Cost, c.Gold)
	}

	CraftItem(item)
}

func Equipement(c *projet.Character) {
	var choix int
	fmt.Println("===== Équipement =====")
	fmt.Println("1 : Chapeau de l'aventurier")
	fmt.Println("2 : Tunique de l'aventurier")
	fmt.Println("3 : Bottes de l'aventurier")
	fmt.Println("0 : Retour")
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		if qty, ok := c.Inventory["Chapeau de l'aventurier"]; ok && qty > 0 {
			c.Equip.Head = "Chapeau de l'aventurier"
			c.Inventory["Chapeau de l'aventurier"]--
			if c.Inventory["Chapeau de l'aventurier"] == 0 {
				delete(c.Inventory, "Chapeau de l'aventurier")
			}
			fmt.Println("Chapeau équipé.")
		} else {
			fmt.Println("Aucun chapeau trouvé dans l'inventaire.")
		}
	case 2:
		if qty, ok := c.Inventory["Tunique de l'aventurier"]; ok && qty > 0 {
			c.Equip.Chestplate = "Tunique de l'aventurier"
			c.Inventory["Tunique de l'aventurier"]--
			if c.Inventory["Tunique de l'aventurier"] == 0 {
				delete(c.Inventory, "Tunique de l'aventurier")
			}
			fmt.Println("Tunique équipée.")
		} else {
			fmt.Println("Aucune tunique trouvée dans l'inventaire.")
		}
	case 3:
		if qty, ok := c.Inventory["Bottes de l'aventurier"]; ok && qty > 0 {
			c.Equip.Feet = "Bottes de l'aventurier"
			c.Inventory["Bottes de l'aventurier"]--
			if c.Inventory["Bottes de l'aventurier"] == 0 {
				delete(c.Inventory, "Bottes de l'aventurier")
			}
			fmt.Println("Bottes équipées.")
		} else {
			fmt.Println("Aucune botte trouvée dans l'inventaire.")
		}
	case 0:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}
