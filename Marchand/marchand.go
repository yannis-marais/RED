package ProjetRED

import (
	Equipement "ProjetRED/Equipement"
	personnage "ProjetRED/Personnage"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	inventoryUpgradePrice uint = 50
	inventoryUpgradeSize  uint = 5
)

func InventoryUpgradeCost(_ *personnage.Character) uint {
	return inventoryUpgradePrice
}

// UpgradeInventory augmente de cinq places la capacité de l'inventaire.
// Le paiement est effectué uniquement si le personnage possède assez d'or.
func UpgradeInventory(player *personnage.Character) bool {
	if player == nil {
		return false
	}

	// Les anciennes sauvegardes peuvent ne pas contenir la capacité.
	if player.Inventory.Capacity == 0 {
		player.Inventory.Capacity = personnage.DefaultInventoryCapacity
	}

	if player.Purse < inventoryUpgradePrice {
		return false
	}

	player.Purse -= inventoryUpgradePrice
	player.Inventory.Capacity += inventoryUpgradeSize
	return true
}

var achatCatalogue = []string{
	"Swordshield", "Elementalist_Rings",
	"Bandit_Helmet", "Bandit_Armor", "Bandit_Boots", "Bandit_Spear",
	"Samourai_Helmet", "Samourai_Armor", "Samourai_Boots",
	"Lord_Armor", "Lord_Boots", "Lord_Battle_Axe",
	"Mage_Staff", "Mage_Hood", "Mage_Robe", "Mage_Boots",
	"Katana", "Dagger",
	"Healing_Potion", "Poison_DOT_Potion", "Pain",
}

var prix = map[string]int{
	"Swordshield": 20, "Elementalist_Rings": 30,
	"Bandit_Helmet": 50, "Bandit_Armor": 50, "Bandit_Boots": 50, "Bandit_Spear": 50,
	"Samourai_Helmet": 60, "Samourai_Armor": 60, "Samourai_Boots": 60,
	"Lord_Armor": 80, "Lord_Boots": 80, "Lord_Battle_Axe": 80,
	"Mage_Staff": 50, "Mage_Hood": 50, "Mage_Robe": 50, "Mage_Boots": 50,
	"Katana": 50, "Dagger": 40,
	"Lord_Helmet":       80,
	"Archimage_Rings":   120,
	"Archimage_Hood":    120,
	"Archimage_Robe":    120,
	"Archimage_Boots":   120,
	"Huge_Cleaver":      100,
	"Odachi":            100,
	"Demonium_Staff":    100,
	"Straw_Hat":         10,
	"Leather_Patch":     10,
	"Boots":             10,
	"Fork":              5,
	"Healing_Potion":    20,
	"Poison_DOT_Potion": 30,
	"Pain":              15,
	"Fer":               5, "Bois": 5, "Cuir": 8, "Cristal": 25, "Diamant": 100,
}

var inventaire = map[string]int{}

// Marchand affiche le menu unique du marchand : achat d'objets tirés au
// hasard dans le catalogue, amélioration de l'inventaire et vente d'objets.
func Marchand(p *personnage.Character, retour func()) {
	if p == nil {
		return
	}

	n := 4
	if n > len(achatCatalogue) {
		n = len(achatCatalogue)
	}
	copie := make([]string, len(achatCatalogue))
	copy(copie, achatCatalogue)
	rand.Shuffle(len(copie), func(i, j int) {
		copie[i], copie[j] = copie[j], copie[i]
	})
	tirage := copie[:n]

	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("Tu as", p.Purse, "pièces")
		for i, nom := range tirage {
			fmt.Printf("%d - %s : %d pièces\n", i+1, nom, prix[nom])
		}
		fmt.Printf("u - Upgrade inventaire : +%d emplacements pour %d pièces\n", inventoryUpgradeSize, InventoryUpgradeCost(p))
		fmt.Println("v - Vendre un objet de ton inventaire")
		fmt.Println("q - Quitter le marchand")

		var saisie string
		fmt.Scanln(&saisie)

		switch saisie {
		case "q":
			if retour != nil {
				retour()
			}
			return
		case "v":
			vendre(p)
			continue
		case "u":
			cost := InventoryUpgradeCost(p)
			if UpgradeInventory(p) {
				fmt.Printf("Inventaire agrandi à %d emplacements pour %d pièces.\n", p.Inventory.Capacity, cost)
			} else {
				fmt.Println("Tu n'as pas assez de pièces.")
			}
			continue
		}

		choix, err := strconv.Atoi(saisie)
		if err != nil || choix < 1 || choix > n {
			fmt.Println("choix invalide")
			continue
		}

		nomChoisi := tirage[choix-1]
		prixAchat := prix[nomChoisi]
		if p.Purse < uint(prixAchat) {
			fmt.Println("tu n'as pas assez de pièces")
			continue
		}

		item, ok := Equipement.Items[nomChoisi]
		if !ok {
			fmt.Println("objet indisponible")
			continue
		}

		before := p.Inventory.Items[item.Name]
		Equipement.AddItem(p, item)
		if p.Inventory.Items[item.Name] <= before {
			fmt.Println("impossible d'ajouter l'objet à l'inventaire (inventaire plein ?)")
			continue
		}

		p.Purse -= uint(prixAchat)
		inventaire[nomChoisi]++
		fmt.Printf("tu as acheté %s pour %d pièces\n", nomChoisi, prixAchat)
	}
}

func vendre(p *personnage.Character) {
	if len(inventaire) == 0 {
		fmt.Println("tu n'as rien à vendre")
		return
	}

	fmt.Println("Objets vendables :")
	noms := make([]string, 0, len(inventaire))
	for nom, qte := range inventaire {
		fmt.Printf("%d - %s (x%d)\n", len(noms)+1, nom, qte)
		noms = append(noms, nom)
	}

	fmt.Println("tape le numéro de l'objet à vendre, ou 0 pour annuler")
	var saisie string
	fmt.Scanln(&saisie)

	choix, err := strconv.Atoi(saisie)
	if err != nil || choix < 0 || choix > len(noms) {
		fmt.Println("choix invalide")
		return
	}
	if choix == 0 {
		return
	}

	nom := noms[choix-1]
	prixDeVente := prix[nom] / 2

	p.Purse += uint(prixDeVente)
	inventaire[nom]--
	fmt.Printf("tu as vendu %s pour %d pièces\n", nom, prixDeVente)

	if inventaire[nom] <= 0 {
		delete(inventaire, nom)
	}
}
