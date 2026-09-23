package ProjetRED

var Classes = map[string]Classe{
	// sabre / katana
	"Ronin": {
		Nom:      "Ronin",
		PVMax:    51,
		Strength: 27,
		Defense:  18,
		Reiki:    2,
		Spd:      20,
		Weapon:   "Bandit_Spear",
		Helmet:   "Bandit_Helmet",
		Armor:    "Bandit_Armor",
		Boots:    "Bandit_Boots",
	},
	// bouclier combat(donc epée + bouclier)
	"Cuirassé": {
		Nom:      "Cuirassé",
		PVMax:    58,
		Strength: 16,
		Defense:  22,
		Reiki:    6,
		Spd:      5,
		Weapon:   "Swordshield",
		Helmet:   "Samourai_Helmet",
		Armor:    "Samourai_Armor",
		Boots:    "Samourai_Boots",
	},
	// un seul et meme item (bague boucle bijoux)
	"mage spirituel": {
		Nom:      "mage spirituel",
		PVMax:    44,
		Strength: 8,
		Defense:  10,
		Reiki:    32,
		Spd:      13,
		Weapon:   "Elementalist_Rings",
		Helmet:   "Mage_Hood",
		Armor:    "Mage_Robe",
		Boots:    "Mage_Boots",
	},
}
