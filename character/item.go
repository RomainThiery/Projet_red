package character

import "fmt"

func (p *Character) SpellBook() bool {
	for _, skill := range p.Skill {
		if skill == "Boule de feu" {
			fmt.Println("Tu connais déjà ce sort !")
			return false
		}
	}
	p.Skill = append(p.Skill, "Boule de feu")
	fmt.Println("Tu as appris le sort : Boule de feu !")
	return true
}

func (p *Character) UseItem(itemName string) {
	switch itemName {
	case "Potion de vie", "potion de vie":
		if p.CurrentHP >= p.MaxHP {
			fmt.Println("Tes PV sont déjà au maximum !")
			return
		}
		p.CurrentHP += 30
		if p.CurrentHP > p.MaxHP {
			p.CurrentHP = p.MaxHP
		}
		fmt.Printf("Tu bois une Potion de vie (+30 PV). PV : %d/%d\n", p.CurrentHP, p.MaxHP)
		p.RemoveFromInventory(itemName)

	case "Potion de mana", "potion de mana":
		if p.CurrentMana >= p.MaxMana {
			fmt.Println("Ton Mana est déjà au maximum !")
			return
		}
		p.CurrentMana += 25
		if p.CurrentMana > p.MaxMana {
			p.CurrentMana = p.MaxMana
		}
		fmt.Printf("🧪 Tu bois une Potion de mana (+25 Mana). Mana : %d/%d\n", p.CurrentMana, p.MaxMana)
		p.RemoveFromInventory(itemName)

	case "Potion de poison", "potion de poison":
		fmt.Println("🧪 Tu prépares une Potion de poison à lancer sur le monstre !")
		p.RemoveFromInventory(itemName)

	case "Livre de Sort : Boule de Feu", "Livre de sort : Boule de feu":
		if p.SpellBook() {
			p.RemoveFromInventory(itemName)
		}

	default:
		fmt.Println("Cet objet ne peut pas être consommé.")
	}
}

// Retire l'objet de l'inventaire
func (p *Character) RemoveFromInventory(itemName string) {
	for i, item := range p.Inventory {
		if item == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			break
		}
	}
}
