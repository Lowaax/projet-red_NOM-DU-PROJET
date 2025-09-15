package main

import (
	"fmt"
	"math/rand"
	projet "projet_red_NOM-DU-PROJET/char"
	"strings"
	"time"
)

func main() {
	c1 := characterCreation()
	rand.Seed(time.Now().UnixNano())
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
		MaxMana:           10,
		Mana:              10,
		Initiative:        5,
		Exp:               0,
		ExpMax:            10,
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
	fmt.Println("11) Vampire")
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
	case 11: // Vampire
		race, maxHP = "Vampire", 100
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

	switch class {
	case "Chevalier":
		c.Inventory = append(c.Inventory, "Épée courte", "Armure rembourrée")
		c.Equip.Arme = "Épée courte"
		c.Equip.Torse = "Armure rembourrée"

	case "Sorcier":
		c.Inventory = append(c.Inventory, "Bâton usé", "Robe simple", "Grimoire débutant")
		c.Equip.Arme = "Bâton usé"
		c.Equip.Torse = "Robe simple"

	case "Archer":
		c.Inventory = append(c.Inventory, "Arc court", "Carquois (x20)", "Tunique légère")
		c.Equip.Arme = "Arc court"
		c.Equip.Torse = "Tunique légère"

	case "Assassin":
		c.Inventory = append(c.Inventory, "Dague", "Cape sombre", "Bottes souples")
		c.Equip.Arme = "Dague"
		c.Equip.Torse = "Cape sombre"
		c.Equip.Pieds = "Bottes souples"

	case "Prêtre":
		c.Inventory = append(c.Inventory, "Masse légère", "Robe bénie", "Amulette de protection")
		c.Equip.Arme = "Masse légère"
		c.Equip.Torse = "Robe bénie"
		c.Equip.Talisman = "Amulette de protection"

	case "Necromancien":
		c.Inventory = append(c.Inventory, "Bâton d’os", "Robe noire", "Talisman occulte")
		c.Equip.Arme = "Bâton d’os"
		c.Equip.Torse = "Robe noire"
		c.Equip.Talisman = "Talisman occulte"

	case "Berserker":
		c.Inventory = append(c.Inventory, "Hache rouillée", "Bandeau", "Ceinture de cuir")
		c.Equip.Arme = "Hache rouillée"
		c.Equip.Tête = "Bandeau"
		c.Equip.Ceinture = "Ceinture de cuir"
	}

	recomputeMaxHP(c)
	return c
}

func displayInfo(c *projet.Character) {
	fmt.Println("===== Informations du personnage =====")
	fmt.Println("Nom       :", c.Name)
	fmt.Println("Race      :", c.Race)
	fmt.Println("Classe    :", c.Class)
	fmt.Println("Niveau    :", c.Level)
	fmt.Println("PV        :", c.HP, "/", c.MaxHP)
	fmt.Printf("Mana      : %d / %d\n", c.Mana, c.MaxMana)
	fmt.Printf("Initiative: %d\n", c.Initiative)
	fmt.Printf("XP        : %d / %d\n", c.Exp, c.ExpMax)
	fmt.Printf("Équipé    : Arme[%s] Tête[%s] Torse[%s] Gant[%s] Centure[%s] Pieds[%s] Talisman[%s]\n",
		func() string {
			if c.Equip.Arme == "" {
				return "Coup de poing"
			}
			return c.Equip.Arme
		}(),
		c.Equip.Tête, c.Equip.Torse, c.Equip.Pieds, c.Equip.Gants, c.Equip.Ceinture, c.Equip.Talisman)
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
	case "Potion de mana":
		manaPot(c)

	default:
		if _, ok := projet.ArmeDB[item]; ok && item != "Coup de poing" {
			equipWeapon(c, item)
			return
		}
		if _, ok := projet.ArmureDB[item]; ok {
			equipArmor(c, item)
			return
		}
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

func manaPot(c *projet.Character) {
	removeItem(c, "Potion de mana")
	restore := 15
	c.Mana += restore
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}
	fmt.Printf("✨ Potion de mana utilisée ! Mana : %d/%d\n", c.Mana, c.MaxMana)
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
	fmt.Println("9: Potion de mana (5 or)")
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
		9: {"Potion de mana", 5},
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
			goblinFight(c)
		case 6:
			trainingFight(c)
		case 7:
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

func recomputeMaxHP(c *projet.Character) {
	base := c.BaseMaxHP
	if a, ok := projet.ArmureDB[c.Equip.Tête]; ok {
		base += a.HPBonus
	}
	if a, ok := projet.ArmureDB[c.Equip.Torse]; ok {
		base += a.HPBonus
	}
	if a, ok := projet.ArmureDB[c.Equip.Pieds]; ok {
		base += a.HPBonus
	}
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

		if c.Equip.Tête != "" {
			c.Inventory = append(c.Inventory, c.Equip.Tête)
		}

		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
		c.Equip.Tête = "Chapeau de l'aventurier"
		recomputeMaxHP(c)
		fmt.Println("Chapeau équipé.")
	case 2:
		idx, ok := pickFromInv("Tunique de l'aventurier")
		if !ok {
			fmt.Println("Aucune tunique trouvée dans l'inventaire.")
			return
		}
		if c.Equip.Torse != "" {
			c.Inventory = append(c.Inventory, c.Equip.Torse)
		}
		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
		c.Equip.Torse = "Tunique de l'aventurier"
		recomputeMaxHP(c)
		fmt.Println("Tunique équipée.")
	case 3:
		idx, ok := pickFromInv("Bottes de l'aventurier")
		if !ok {
			fmt.Println("Aucune botte trouvée dans l'inventaire.")
			return
		}
		if c.Equip.Pieds != "" {
			c.Inventory = append(c.Inventory, c.Equip.Pieds)
		}
		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
		c.Equip.Pieds = "Bottes de l'aventurier"
		recomputeMaxHP(c)
		fmt.Println("Bottes équipées.")
	case 0:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}

type Monster struct {
	Name       string
	MaxHP      int
	HP         int
	ATK        int
	Initiative int // ← utilisé pour comparer
	XPReward   int
}

func goblinPattern(g *Monster, c *projet.Character, turn int) {
	if g == nil || c == nil || g.HP <= 0 {
		return
	}
	dmg := g.ATK
	if turn%3 == 0 {
		dmg = g.ATK * 2
		fmt.Printf("%s déclenche une attaque renforcée ! %s subit %d dégâts.\n", g.Name, c.Name, dmg)
	} else {
		fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", g.Name, dmg, c.Name)
	}
	c.HP -= dmg
	if c.HP < 0 {
		c.HP = 0
	}
	fmt.Printf("PV de %s : %d / %d\n", c.Name, c.HP, c.MaxHP)
	if c.HP <= 0 {
		_ = IsDead(c)
	}
}

func characterTurn(c *projet.Character, g *Monster) (ended bool) {
	for {
		var choix int
		fmt.Println("===== Votre tour =====")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Fuir")
		fmt.Println("4. Sorts")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			weapons := []string{"Coup de poing"}
			if c.Equip.Arme != "" {
				weapons = append(weapons, c.Equip.Arme)
			}

			fmt.Println("— Choisir l’arme —")
			for i, w := range weapons {
				min, max := weaponDamageRange(w)
				fmt.Printf("%d) %s (%d–%d)\n", i+1, w, min, max)
			}
			fmt.Print("Arme : ")
			var wch int
			fmt.Scan(&wch)
			if wch < 1 || wch > len(weapons) {
				fmt.Println("Choix invalide.")
				continue
			}
			wname := weapons[wch-1]
			min, max := weaponDamageRange(wname)
			dmg := rollDamage(min, max)

			g.HP -= dmg
			if g.HP < 0 {
				g.HP = 0
			}
			fmt.Printf("%s utilise %s et inflige %d dégâts à %s.\n", c.Name, wname, dmg, g.Name)
			fmt.Printf("PV de %s : %d / %d\n", g.Name, g.HP, g.MaxHP)
			return false

		case 2:
			AccessInventory(c)
			return FalseIfBothAlive(c, g)
		case 3:
			fmt.Println("Vous fuyez le combat.")
			return true
		case 4:
			castSpell(c, g)
			return FalseIfBothAlive(c, g)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func FalseIfBothAlive(c *projet.Character, g *Monster) bool {
	if c.HP <= 0 || g.HP <= 0 {
		return TrueIfAnyDown(c, g)
	}
	return false
}

func TrueIfAnyDown(c *projet.Character, g *Monster) bool {
	return c.HP <= 0 || g.HP <= 0
}

func trainingFight(c *projet.Character) {
	g := initTrainingGoblin()
	turn := 1
	fmt.Printf("Un %s apparaît ! (ATK %d, Init %d)\n", g.Name, g.ATK, g.Initiative)

	firstPlayer := playerStarts(c, g)
	if firstPlayer {
		fmt.Printf("🟩 %s commence (Initiative %d) !\n", c.Name, c.Initiative)
	} else {
		fmt.Printf("🟥 %s commence (Initiative %d) !\n", g.Name, g.Initiative)
	}

	for {
		fmt.Printf("\n===== Tour %d =====\n", turn)
		if characterTurn(c, g) {
			fmt.Println("Fin du combat.")
			return
		}
		if g.HP <= 0 {
			fmt.Printf("%s est vaincu !\n", g.Name)
			rewardVictory(c, g)
			fmt.Println("Fin du combat.")
			return
		}
		fmt.Println("Le gobelin d'entraînement ne riposte pas.")
		turn++
	}
}

func goblinFight(c *projet.Character) {
	g := initGoblin()
	turn := 1
	fmt.Printf("Un %s apparaît ! (ATK %d, Init %d)\n", g.Name, g.ATK, g.Initiative)

	firstPlayer := playerStarts(c, g)
	if firstPlayer {
		fmt.Printf("🟩 %s commence (Initiative %d) !\n", c.Name, c.Initiative)
	} else {
		fmt.Printf("🟥 %s commence (Initiative %d) !\n", g.Name, g.Initiative)
	}

	for {
		fmt.Printf("\n===== Tour %d =====\n", turn)

		if firstPlayer {
			if characterTurn(c, g) {
				fmt.Println("Fin du combat.")
				return
			}
			if g.HP <= 0 {
				fmt.Printf("%s est vaincu !\n", g.Name)
				rewardVictory(c, g)
				fmt.Println("Fin du combat.")
				return
			}
			goblinPattern(g, c, turn)
			if c.HP <= 0 {
				fmt.Println("Fin du combat.")
				return
			}
		} else {
			goblinPattern(g, c, turn)
			if c.HP <= 0 {
				fmt.Println("Fin du combat.")
				return
			}
			if characterTurn(c, g) {
				fmt.Println("Fin du combat.")
				return
			}
			if g.HP <= 0 {
				fmt.Printf("%s est vaincu !\n", g.Name)
				rewardVictory(c, g) // ← ICI aussi
				fmt.Println("Fin du combat.")
				return
			}
		}
		turn++
	}
}

func initGoblin() *Monster {
	return &Monster{
		Name:       "Gobelin",
		MaxHP:      40,
		HP:         40,
		ATK:        5,
		Initiative: 6,
		XPReward:   10,
	}
}

func initTrainingGoblin() *Monster {
	return &Monster{
		Name:       "Gobelin d'entraînement",
		MaxHP:      40,
		HP:         40,
		ATK:        0,
		Initiative: 4,
		XPReward:   2,
	}
}

func weaponDamageRange(weaponName string) (min, max int) {
	if ws, ok := projet.ArmeDB[weaponName]; ok {
		return ws.Min, ws.Max
	}
	// fallback
	return projet.ArmeDB["Coup de poing"].Min, projet.ArmeDB["Coup de poing"].Max
}

func rollDamage(min, max int) int {
	if max < min {
		max = min
	}
	return rand.Intn(max-min+1) + min
}

func equipWeapon(c *projet.Character, name string) {
	if _, ok := projet.ArmeDB[name]; !ok {
		fmt.Println("Ce n’est pas une arme équipable.")
		return
	}
	if c.Equip.Arme != "" {
		c.Inventory = append(c.Inventory, c.Equip.Arme)
	}
	removeItem(c, name)
	c.Equip.Arme = name
	fmt.Printf("🔪 Arme équipée : %s\n", name)
}

func equipArmor(c *projet.Character, name string) {
	a, ok := projet.ArmureDB[name]
	if !ok {
		fmt.Println("Ce n’est pas une armure connue.")
		return
	}
	removeItem(c, name)
	switch a.Slot {
	case "Head":
		if c.Equip.Tête != "" {
			c.Inventory = append(c.Inventory, c.Equip.Tête)
		}
		c.Equip.Tête = name
	case "Chestplate":
		if c.Equip.Torse != "" {
			c.Inventory = append(c.Inventory, c.Equip.Torse)
		}
		c.Equip.Torse = name
	case "Feet":
		if c.Equip.Pieds != "" {
			c.Inventory = append(c.Inventory, c.Equip.Pieds)
		}
		c.Equip.Pieds = name
	}

	recomputeMaxHP(c)
	fmt.Printf("🛡️ Armure équipée : %s (+%d PV max)\n", name, a.HPBonus)
}

func playerStarts(c *projet.Character, g *Monster) bool {
	if c.Initiative > g.Initiative {
		return true
	}
	if c.Initiative < g.Initiative {
		return false
	}
	return rand.Intn(2) == 0
}

func rewardVictory(c *projet.Character, g *Monster) {
	if g.XPReward <= 0 {
		return
	}
	fmt.Printf("🏅 %s gagne %d XP.\n", c.Name, g.XPReward)
	c.Exp += g.XPReward
	for c.Exp >= c.ExpMax {
		c.Exp -= c.ExpMax
		levelUp(c)
	}
	displayInfo(c)
}

func levelUp(c *projet.Character) {
	c.Level++
	c.BaseMaxHP += 5
	recomputeMaxHP(c)
	c.MaxMana += 2
	c.Initiative += 1
	c.ExpMax += 10
	c.HP = c.MaxHP
	c.Mana = c.MaxMana
	fmt.Printf("⬆️  Niveau %d ! (+5 PV max, +2 Mana max, +1 Initiative)\n", c.Level)
}

func castSpell(c *projet.Character, g *Monster) (ended bool) {
	if len(c.Skills) == 0 {
		fmt.Println("Vous ne connaissez aucun sort.")
		return false
	}

	spellCost := map[string]int{
		"Coup de poing": 3,
		"Boule de Feu":  7,
	}
	spellDamage := map[string]int{
		"Coup de poing": 8,
		"Boule de Feu":  18,
	}

	known := []string{}
	for _, s := range c.Skills {
		if _, ok := spellDamage[s]; ok {
			known = append(known, s)
		}
	}
	if len(known) == 0 {
		fmt.Println("Aucun sort offensif disponible.")
		return false
	}

	fmt.Println("— Choisir un sort —")
	for i, s := range known {
		fmt.Printf("%d) %s (coût %d mana, %d dégâts)\n", i+1, s, spellCost[s], spellDamage[s])
	}
	fmt.Print("Sort : ")
	var sch int
	fmt.Scan(&sch)
	if sch < 1 || sch > len(known) {
		fmt.Println("Choix invalide.")
		return false
	}
	name := known[sch-1]
	cost := spellCost[name]
	dmg := spellDamage[name]

	if c.Mana < cost {
		fmt.Printf("❌ Mana insuffisant (%d/%d). Ce sort nécessite %d mana.\n", c.Mana, c.MaxMana, cost)
		return false
	}
	c.Mana -= cost
	g.HP -= dmg
	if g.HP < 0 {
		g.HP = 0
	}
	fmt.Printf("%s lance %s et inflige %d dégâts à %s. (Mana %d/%d)\n", c.Name, name, dmg, g.Name, c.Mana, c.MaxMana)
	fmt.Printf("PV de %s : %d / %d\n", g.Name, g.HP, g.MaxHP)
	return false
}
