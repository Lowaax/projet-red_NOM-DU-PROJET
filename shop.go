package main

import (
	"fmt"
	projet "projet_red_NOM-DU-PROJET/char"
)

func Shop(c *projet.Character) {
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
		c.Inventory()["Potion"] += 1
	} else if choice_shop == 0 {
		fmt.Println("Marchand: Tu change d'avis ? bah dégage alors")
	} else {
		fmt.Println("T'es vraiment con toi, ", choice_shop, "n'est pas une option")
	}
}
