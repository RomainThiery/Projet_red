package fight

import (
	"Projet-Red/character"
	"fmt"
)

func Dead(c *character.Character) {
	if c.CurrentHP <= 0 {
		fmt.Println("\n☠️ Défaite... Tu as été vaincu par le monstre !")
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("🩹 Tu te relèves péniblement avec %d/%d PV.\n", c.CurrentHP, c.MaxHP)
	}
}
