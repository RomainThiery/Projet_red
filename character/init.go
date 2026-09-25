package character

import (
	"fmt"
	"time"
)

func Init() Character {
	var nom string

	// --- INTRO & LORE ---
	fmt.Println("==========================================================================")
	fmt.Println("📜  LES CHRONIQUES DU ROYAUME D'ELDORIA")
	fmt.Println("==========================================================================")
	fmt.Println("Depuis des siècles, la terre d'Eldoria vivait dans la paix et l'abondance.")
	fmt.Println("Mais l'ombre du Dragon Ancien, Ignis, s'est abattue sur la région.")
	fmt.Println("Brûlant les villages et terrorisant les habitants, le monstre s'est réapproprié")
	fmt.Println("les terres sauvages. La légende raconte que seul un aventurier au cœur pur")
	fmt.Println("et à la volonté d'acier pourra gravir les échelons et terrasser la bête.")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("Tu es cet aventurier. Ton périple commence au bas de l'échelle...")
	fmt.Println("==========================================================================")
	fmt.Println()

	// --- CHOIX DU NOM ---
	fmt.Print("⚔️  Aventurier, quel est ton nom ? : ")
	fmt.Scanln(&nom)

	// Sécurité si le joueur appuie directement sur Entrée
	if nom == "" {
		nom = "Héros Inconnu"
	}

	fmt.Println()
	fmt.Printf("🏰 Bienvenue à toi, %s ! Ton nom sera gravé dans l'histoire.\n", nom)
	fmt.Println("Prépare tes armes et ton entraînement. La traque d'Ignis commence maintenant !")
	fmt.Println("==========================================================================")

	time.Sleep(2 * time.Second) // Petite pause narrative avant le menu principal

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
