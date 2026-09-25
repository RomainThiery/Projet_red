package character

import "fmt"

func DisplayInfo(p *Character) {
	fmt.Println("=== PERSONNAGE ===")

	fmt.Println("Nom :", p.Name)
	fmt.Println("Classe :", p.Classe)
	fmt.Println("Niveau :", p.Niveau)
	fmt.Println("PV :", p.CurrentHP, "/", p.MaxHP)

	fmt.Println("=== COMPÉTENCES ===")

	for _, skill := range p.Skill {
		fmt.Println("-", skill)
	}
}
