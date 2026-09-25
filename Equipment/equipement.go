package equipment

import "fmt"

type Equipment struct {
	Head  string
	Torso string
	Legs  string
	Feet  string
}

func removeSliceItem(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func EquipItem(e *Equipment, inventory *[]string, maxHP *int, itemName string) {
	itemIndex := -1
	for i, item := range *inventory {
		if item == itemName {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		fmt.Printf("Tu ne possèdes pas %s dans ton inventaire.\n", itemName)
		return
	}

	switch itemName {
	case "Helmet":
		if e.Head != "" {
			*inventory = append(*inventory, e.Head)
			*maxHP -= 20
		}
		e.Head = itemName
		*maxHP += 20
		fmt.Println("Casque équipé ! (+20 PV max)")

	case "Chestplate":
		if e.Torso != "" {
			*inventory = append(*inventory, e.Torso)
			*maxHP -= 40
		}
		e.Torso = itemName
		*maxHP += 40
		fmt.Println("Plastron équipé ! (+40 PV max)")

	case "Leggings":
		if e.Legs != "" {
			*inventory = append(*inventory, e.Legs)
			*maxHP -= 25
		}
		e.Legs = itemName
		*maxHP += 25
		fmt.Println("Pantalon équipé ! (+25 PV max)")

	case "Boots":
		if e.Feet != "" {
			*inventory = append(*inventory, e.Feet)
			*maxHP -= 15
		}
		e.Feet = itemName
		*maxHP += 15
		fmt.Println("Bottes équipées ! (+15 PV max)")

	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
		return
	}

	*inventory = removeSliceItem(*inventory, itemIndex)
}
