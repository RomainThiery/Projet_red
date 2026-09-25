package character

import "fmt"

func Init() Character {
	var nom string

	fmt.Print("Entrez votre pseudo : ")
	fmt.Scanln(&nom)

	p := Character{
		Name:        nom,
		Classe:      "Elfe",
		Niveau:      1,
		MaxHP:       100,
		CurrentHP:   100,
		CurrentMana: 100,
		MaxMana:     100,
		Inventory: []string{
			"potion de vie",
			"potion de vie",
			"potion de vie",
		},

		Skill: []string{"Coup de poing"},
	}
	return p
}
