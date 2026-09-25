package character

import "Projet-Red/equipment"

type Character struct {
	Name        string
	Classe      string
	Niveau      int
	MaxHP       int
	CurrentHP   int
	CurrentMana int
	MaxMana     int
	Equipment   equipment.Equipment
	Inventory   []string
	Skill       []string
}
