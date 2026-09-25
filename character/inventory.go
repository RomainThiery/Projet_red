package character

import (
	"Projet-Red/equipment"
	"fmt"
)

func (p *Character) OpenEquipmentMenu() {
	for {
		fmt.Println("\n=== 🎒 INVENTAIRE ===")
		if len(p.Inventory) == 0 {
			fmt.Println("(Aucun objet dans l'inventaire)")
		} else {
			for i, item := range p.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		}
		fmt.Println("0. Quitter")
		fmt.Print("Quel objet veux-tu équiper ? : ")

		var choice int
		fmt.Scanln(&choice)
		if choice == 0 {
			fmt.Println("Fermeture de l'inventaire.")
			break
		} else if choice > 0 && choice <= len(p.Inventory) {
			selectedItem := p.Inventory[choice-1]
			equipment.EquipItem(&p.Equipment, &p.Inventory, &p.MaxHP, selectedItem)
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}
