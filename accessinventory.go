package main

import "fmt"

func accessInventory(character class) {
	inventaire := character.Inventaire()
	fmt.Println("Inventaire du personnage :")
	for item, quantite := range inventaire {
		fmt.Printf("- %s : %d\n", item, quantite)
	}
}
