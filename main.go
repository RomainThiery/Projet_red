package main

import (
	"Projet-Red/character"
	"Projet-Red/menu"
	"Projet-Red/wallet"
)

func main() {
	myWallet := wallet.Wallet{GoldCoins: 100}

	player := character.Init()

	menu.StartMainMenu(&myWallet, &player)
}
