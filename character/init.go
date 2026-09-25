package character

import "fmt"

func Init() Character {
	var nom string

	fmt.Print("Entrez le nom de votre personnage")
	fmt.Scan(&nom)

	p := Character{
		Name:      "Harmony",
		Classe:    "Elfe",
		Niveau:    1,
		MaxHP:     100,
		CurrentHP: 40,
		Inventory: []string{
			"potion de vie",
			"potion de vie",
			"potion de vie",
		},

		Skill: []string{"Coup de poing"},
	}
	return p
}
