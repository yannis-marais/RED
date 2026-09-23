package main

import (
	Equipement "ProjetRED/Equipement"

	Menu "ProjetRED/Menu"
	personnage "ProjetRED/Personnage"
)

func main() {
	jean := personnage.CharacterCreation("jean", personnage.Classes["Ronin"])
	Equipement.AddConsumable(&jean, Equipement.HealingPotion)
	Equipement.AddItem(&jean, Equipement.Items["Swordshield"])
	Equipement.AddConsumable(&jean, Equipement.HealingPotion)
	Equipement.AddItem(&jean, Equipement.Items["Dagger"])
	Equipement.AddConsumable(&jean, Equipement.PoisonDOTPotion)
	jean.PV = 20
	Menu.DisplayInfo(jean)
	Menu.AccessInventory(jean)
	Menu.ManageInventory(&jean)
	Menu.MainMenu(&jean)

}
func turn() {
	// placeholder: besoin de définir une cible et un personnage pour lancer les combats
}