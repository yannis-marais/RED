package ProjetRED

import (
	"fmt"
	"math/rand"
)

func marchand() {
	type Item struct {
		Nom  string
		Prix int
	}
	items := []Item{
		{"Swordshield", 20},
		{"Elementalist Rings", 30},
		{"Bandit Helmet", 50},
		{"Bandit Armor", 50},
		{"Bandit Boots", 50},
		{"Bandit Spear", 50},
		{"Samourai Helmet", 60},
		{"Samourai Armor", 60},
		{"Samourai Boots", 60},
		{"Lord Armor", 80},
		{"Lord Boots", 80},
		{"Lord Battle Axe", 80},
		{"Mage Staff", 50},
		{"Mage Hood", 50},
		{"Mage Robe", 50},
		{"Mage Boots", 50},
		{"Katana", 50},
		{"Dagger", 40},
		{"Healing Potion", 60},
		{"Poison DOT Potion", 60},
	}
	n := 4
	if n > len(items) {
		n = len(items)
	}
	copie := make([]Item, len(items))
	copy(copie, items)
	rand.Shuffle(len(copie), func(i, j int) {
		copie[i], copie[j] = copie[j], copie[i]
	})
	tirage := copie[:n]

	total := 0
	for _, it := range tirage {
		fmt.Printf("%s : %d pièces\n", it.Nom, it.Prix)
		total += it.Prix
	}
	fmt.Println("Total :", total, "pièces")

	var saisie string
	_ = []func(){}
	fmt.Scanln(&saisie)
}
