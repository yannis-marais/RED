package ProjetRED

import personnage "ProjetRED/Personnage"

func EquipItem(p *personnage.Character, slot string, item Item) {
	switch slot {
	case "weapon":
		p.Weapon = item.Name
	case "armor":
		p.Armor = item.Name
	case "boots":
		p.Boots = item.Name
	case "helmet":
		p.Helmet = item.Name
	}

	RecalculateStats(p)
}
