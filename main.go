package main

func initCharacter() {

func main() {

}

func initCharacter(name string, race Race, class Classes, niveau int, pv_max int, pv_actuel int, inventaire []string) Character {
	return Character{
		Name:       name,
		Race:       race,
		Class:      class,
		Niveau:     niveau,
		Pv_max:     pv_max,
		Pv_actuel:  pv_actuel,
		Inventaire: inventaire,
	}
}
