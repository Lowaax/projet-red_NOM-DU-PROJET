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
			fightGoblin(c)
			return
		case 6:
			trainingFight(c)
		case 7:
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
		for i, item := range c.Inventory {
			if item == "Chapeau de l'aventurier" {
				c.Equip.Head = item
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				fmt.Println("Chapeau équipé.")
				return
			}
		}
		fmt.Println("Aucun chapeau trouvé dans l'inventaire.")
	case 2:
		for i, item := range c.Inventory {
			if item == "Tunique de l'aventurier" {
				c.Equip.Chestplate = item
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				fmt.Println("Tunique équipée.")
				return
			}
		}
		fmt.Println("Aucune tunique trouvée dans l'inventaire.")
	case 3:
		for i, item := range c.Inventory {
			if item == "Bottes de l'aventurier" {
				c.Equip.Feet = item
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				fmt.Println("Bottes équipées.")
				return
			}
		}
		fmt.Println("Aucune botte trouvée dans l'inventaire.")
	case 0:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}

func goblinPattern(c *projet.Character) {
	goblin := projet.InitGoblin()
	tour := 1

	fmt.Println("Début du combat")

	for c.HP > 0 {
		var damage int

		if tour%3 == 0 {
			damage = goblin.Attack * 2
		} else {
			damage = goblin.Attack
		}

		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}

		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, c.Name, damage)
		fmt.Printf("PV : %d/%d\n", c.HP, c.MaxHP)

		if c.HP <= 0 {
			fmt.Println("Vous êtes mort!")
			break
		}

		tour++

		time.Sleep(1 * time.Second)
	}
}

func characterTurn(player *projet.Character, monster *projet.Monster) bool {
	for {
		var choix int
		fmt.Println("===== Tour du joueur =====")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			attaque := "Attaque basique"
			degats := 5
			monster.HP -= degats
			if monster.HP < 0 {
				monster.HP = 0
			}

			fmt.Printf("🗡️  %s utilise %s et inflige %d dégâts.\n", player.Name, attaque, degats)
			fmt.Printf("👹 %s PV : %d/%d\n", monster.Name, monster.HP, monster.MaxHP)

			if monster.HP <= 0 {
				fmt.Printf("✅ %s a été vaincu !\n", monster.Name)
				return true
			}

			time.Sleep(1 * time.Second)
			monsterTurn(player, monster)
			return false

		case 2:
			if len(player.Inventory) == 0 {
				fmt.Println("📦 Inventaire vide. Choisissez une autre action.")
				continue
			}

			fmt.Println("🎒 Inventaire :")
			keys := []string{}
			i := 1
			for item, qty := range player.Inventory {
				fmt.Printf("%d. %s (x%d)\n", i, item, qty)
				keys = append(keys, item)
				i++
			}
			fmt.Println("0. Annuler")
			fmt.Print("Choix : ")
			var invChoice int
			fmt.Scan(&invChoice)

			if invChoice == 0 {
				continue
			}

			if invChoice < 1 || invChoice > len(keys) {
				fmt.Println("❌ Choix invalide.")
				continue
			}

			item := keys[invChoice-1]
			switch item {
			case "Potion":
				player.HP += 50
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				player.Inventory[item]--
				if player.Inventory[item] <= 0 {
					delete(player.Inventory, item)
				}
				fmt.Printf("🧪 Vous utilisez %s. PV : %d/%d\n", item, player.HP, player.MaxHP)

			case "Potion de poison":
				player.Inventory[item]--
				if player.Inventory[item] <= 0 {
					delete(player.Inventory, item)
				}
				fmt.Println("☠️ Vous utilisez Potion de poison !")
				for i := 1; i <= 3; i++ {
					time.Sleep(1 * time.Second)
					player.HP -= 10
					if player.HP <= 0 {
						player.HP = 0
						fmt.Printf("☠️ Poison tick %d → PV : %d/%d\n", i, player.HP, player.MaxHP)
						fmt.Println("💀 Vous êtes mort.")
						IsDead(player)
						break
					}
					fmt.Printf("☠️ Poison tick %d → PV : %d/%d\n", i, player.HP, player.MaxHP)
				}

			case "Livre de Sort : Boule de Feu":
				alreadyKnown := false
				for _, skill := range player.Skills {
					if skill == "Boule de Feu" {
						alreadyKnown = true
						break
					}
				}
				if alreadyKnown {
					fmt.Println("✨ Vous connaissez déjà ce sort.")
				} else {
					player.Skills = append(player.Skills, "Boule de Feu")
					fmt.Println("🔥 Nouveau sort appris : Boule de Feu !")
				}
				player.Inventory[item]--
				if player.Inventory[item] <= 0 {
					delete(player.Inventory, item)
				}
			default:
				fmt.Println("et objet ne peut pas être utilisé en combat.")
			}
			time.Sleep(1 * time.Second)
			monsterTurn(player, monster)
			return false

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func monsterTurn(player *projet.Character, monster *projet.Monster) {
	degats := monster.Attack
	player.HP -= degats
	if player.HP < 0 {
		player.HP = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts !\n", monster.Name, player.Name, degats)
	fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)

	if player.HP <= 0 {
		fmt.Println("Vous êtes mort.")
		IsDead(player)
	}
}

func fightGoblin(player *projet.Character) {
	goblin := projet.InitGoblin()

	fmt.Println("Un Gobelin d'entraînement apparaît!")

	for goblin.HP > 0 && player.HP > 0 {
		victoire := characterTurn(player, &goblin)
		if victoire {
			break
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Fin du combat")
}

func trainingFight(player *projet.Character) {
	goblin := projet.InitGoblin()
	tour := 1

	fmt.Println("Début du combat")

	for goblin.HP > 0 && player.HP > 0 {
		fmt.Printf("\n===== Tour %d =====\n", tour)

		victoire := characterTurn(player, &goblin)
		if victoire || player.HP <= 0 {
			break
		}

		var damage int
		if tour%3 == 0 {
			damage = goblin.Attack * 2
			fmt.Println("Le gobelin lance une grosse attaque !")
		} else {
			damage = goblin.Attack
		}

		player.HP -= damage
		if player.HP < 0 {
			player.HP = 0
		}

		fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", goblin.Name, damage, player.Name)
		fmt.Printf("PV de %s : %d/%d\n", player.Name, player.HP, player.MaxHP)

		if player.HP <= 0 {
			fmt.Println("Vous avez été vaincu lors de l'entraînement.")
			IsDead(player)
			break
		}

		tour++
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Fin du combat")
}
