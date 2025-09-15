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
		Name:              name,
		Race:              race,
		Class:             class,
		Level:             1,
		BaseMaxHP:         maxHP,
		MaxHP:             maxHP,
		HP:                maxHP / 2,
		Inventory:         []string{"Potion", "Potion", "Potion"},
		Skills:            []string{"Coup de poing"},
		Gold:              100,
		MaxInventory:      10,
		InventoryUpgrades: 0,
	}
}

func characterCreation() *projet.Character {
	var name string
	fmt.Print("Choisissez un nom : ")
	fmt.Scan(&name)
	name = strings.Title(strings.ToLower(name))

	fmt.Println("Choisissez une RACE :")
	fmt.Println(" 1) Human")
	fmt.Println(" 2) Elf")
	fmt.Println(" 3) Nain")
	fmt.Println(" 4) Orc")
	fmt.Println(" 5) Dragon")
	fmt.Println(" 6) Mort-vivant")
	fmt.Println(" 7) Ange")
	fmt.Println(" 8) Orque")
	fmt.Println(" 9) Centaure")
	fmt.Println("10) Fée")
	fmt.Println("11) Lycanthrope")
	fmt.Println("12) Farfadet")
	fmt.Println("13) Antromorphe")
	fmt.Print("Votre choix (numéro) : ")

	var raceChoice int
	fmt.Scan(&raceChoice)

	race := "Human"
	maxHP := 100
	switch raceChoice {
	case 1: // Human
		race, maxHP = "Human", 100
	case 2: // Elf
		race, maxHP = "Elf", 80
	case 3: // Nain
		race, maxHP = "Nain", 120
	case 4: // Orc
		race, maxHP = "Orc", 110
	case 5: // Dragon
		race, maxHP = "Dragon", 140
	case 6: // Mort-vivant
		race, maxHP = "Mort-vivant", 90
	case 7: // Ange
		race, maxHP = "Ange", 100
	case 8: // Orque (syno d’orc séparé dans ta liste)
		race, maxHP = "Orque", 110
	case 9: // Centaure
		race, maxHP = "Centaure", 110
	case 10: // Fée
		race, maxHP = "Fée", 70
	case 11: // Lycanthrope
		race, maxHP = "Lycanthrope", 120
	case 12: // Farfadet
		race, maxHP = "Farfadet", 75
	case 13: // Antromorphe
		race, maxHP = "Antromorphe", 100
	default:
		fmt.Println("Entrée invalide → race par défaut : Human (100 PV).")
	}

	fmt.Println("Choisissez une CLASSE :")
	fmt.Println(" 1) Chevalier")
	fmt.Println(" 2) Sorcier")
	fmt.Println(" 3) Archer")
	fmt.Println(" 4) Assassin")
	fmt.Println(" 5) Prêtre")
	fmt.Println(" 6) Necromancien")
	fmt.Println(" 7) Berserker")
	fmt.Print("Votre choix (numéro) : ")

	var classChoice int
	fmt.Scan(&classChoice)

	class := "Chevalier"
	switch classChoice {
	case 1:
		class = "Chevalier"
	case 2:
		class = "Sorcier"
	case 3:
		class = "Archer"
	case 4:
		class = "Assassin"
	case 5:
		class = "Prêtre"
	case 6:
		class = "Necromancien"
	case 7:
		class = "Berserker"
	default:
		fmt.Println("Entrée invalide → classe par défaut : Chevalier.")
	}

	// Création du perso avec PV selon la race (HP init = 50% dans initCharacter)
	c := initCharacter(name, race, class, maxHP)

	// --- STUFF DE BASE PAR CLASSE (simple et direct) ---
	switch class {
	case "Chevalier":
		c.Inventory = append(c.Inventory, "Épée courte", "Bouclier en bois", "Armure rembourrée")
	case "Sorcier":
		c.Inventory = append(c.Inventory, "Bâton usé", "Robe simple", "Grimoire débutant")
	case "Archer":
		c.Inventory = append(c.Inventory, "Arc court", "Carquois (x20)", "Tunique légère")
	case "Assassin":
		c.Inventory = append(c.Inventory, "Dague", "Cape sombre", "Bottes souples")
	case "Prêtre":
		c.Inventory = append(c.Inventory, "Masse légère", "Robe bénie", "Amulette")
	case "Necromancien":
		c.Inventory = append(c.Inventory, "Bâton d’os", "Robe noire", "Talisman occulte")
	case "Berserker":
		c.Inventory = append(c.Inventory, "Hache rouillée", "Bandeau", "Ceinture de cuir")
	}

	return c
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
		fmt.Println("Inventaire:", strings.Join(c.Inventory, ", "))
	}

	fmt.Println("Skills    :", strings.Join(c.Skills, ", "))
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
	fmt.Println("8: Augmentation d'inventaire (+10 slots, 30 or)  — max 3")
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
		return
	}

	if choix == 8 {
		if c.InventoryUpgrades >= 3 {
			fmt.Println("❌ Limite atteinte : vous avez déjà utilisé 3 augmentations.")
			return
		}
		if c.Gold < 30 {
			fmt.Println("❌ Pas assez d’or.")
			return
		}
		c.Gold -= 30
		c.MaxInventory += 10
		c.InventoryUpgrades++
		fmt.Printf("✅ Capacité d’inventaire augmentée à %d (utilisations : %d/3). Or restant : %d.\n",
			c.MaxInventory, c.InventoryUpgrades, c.Gold)
		return
	}

	if choix != 0 {
		fmt.Println("Choix invalide.")
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
		fmt.Println("7. Quitter")
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
			fmt.Println("Au revoir !")
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

	var CraftItem = func(it Item) {
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
			count := 0
			for _, invItem := range c.Inventory {
				if invItem == mat.Name {
					count++
				}
			}
			if count < mat.Quantity {
				fmt.Printf("❌ Il manque %d × %s\n", mat.Quantity-count, mat.Name)
				missing = true
			}
		}
		if missing {
			return
		}

		for _, mat := range it.Materials {
			for i := 0; i < mat.Quantity; i++ {
				removeItem(c, mat.Name)
			}
		}

		c.Gold -= it.Cost
		added := addItem(c, it.Name)
		if !added {
			fmt.Println("❌ Impossible d’ajouter l’objet (inventaire plein).")
			return
		}

		fmt.Printf("✅ %s forgé pour %d or (reste %d).\n", it.Name, it.Cost, c.Gold)
	}

	CraftItem(item)
}

func equipBonus(item string) int {
	switch item {
	case "Chapeau de l'aventurier":
		return 10
	case "Tunique de l'aventurier":
		return 25
	case "Bottes de l'aventurier":
		return 15
	default:
		return 0
	}
}

func recomputeMaxHP(c *projet.Character) {
	base := c.BaseMaxHP
	base += equipBonus(c.Equip.Head)
	base += equipBonus(c.Equip.Chestplate)
	base += equipBonus(c.Equip.Feet)
	c.MaxHP = base
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
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

	pickFromInv := func(needle string) (int, bool) {
		for i, item := range c.Inventory {
			if item == needle {
				return i, true
			}
		}
		return -1, false
	}

	switch choix {
	case 1:
		idx, ok := pickFromInv("Chapeau de l'aventurier")
		if !ok {
			fmt.Println("Aucun chapeau trouvé dans l'inventaire.")
			return
		}

		if c.Equip.Head != "" {
			c.Inventory = append(c.Inventory, c.Equip.Head)
		}

		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
		c.Equip.Head = "Chapeau de l'aventurier"
		recomputeMaxHP(c)
		fmt.Println("Chapeau équipé.")
	case 2:
		idx, ok := pickFromInv("Tunique de l'aventurier")
		if !ok {
			fmt.Println("Aucune tunique trouvée dans l'inventaire.")
			return
		}
		if c.Equip.Chestplate != "" {
			c.Inventory = append(c.Inventory, c.Equip.Chestplate)
		}
		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
		c.Equip.Chestplate = "Tunique de l'aventurier"
		recomputeMaxHP(c)
		fmt.Println("Tunique équipée.")
	case 3:
		idx, ok := pickFromInv("Bottes de l'aventurier")
		if !ok {
			fmt.Println("Aucune botte trouvée dans l'inventaire.")
			return
		}
		if c.Equip.Feet != "" {
			c.Inventory = append(c.Inventory, c.Equip.Feet)
		}
		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
		c.Equip.Feet = "Bottes de l'aventurier"
		recomputeMaxHP(c)
		fmt.Println("Bottes équipées.")
	case 0:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}
