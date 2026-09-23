package ProjetRED

import personnage "ProjetRED/Personnage"

func RecalculateStats(p *personnage.Character) {
	for _, slot := range []string{p.Weapon, p.Armor, p.Boots, p.Helmet} {
		if item, ok := Items[slot]; ok {
			p.PVMax = p.Classe.PVMax + item.BonusPV
			p.Strength = p.Classe.Strength + item.BonusAtk
			p.Defense = p.Classe.Defense + item.BonusDef
			p.Reiki = p.Classe.Reiki + item.BonusReiki
			p.Spd = p.Classe.Spd + item.BonusSpeed
		}
	}

	if p.PV > p.PVMax {
		p.PV = p.PVMax
	}
}
