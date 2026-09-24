package ProjetRED

import personnage "ProjetRED/Personnage"

func RecalculateStats(p *personnage.Character) {
	if p == nil {
		return
	}

	p.PVMax = p.Classe.PVMax
	p.Strength = p.Classe.Strength
	p.Defense = p.Classe.Defense
	p.Reiki = p.Classe.Reiki
	p.Spd = p.Classe.Spd

	for _, slot := range []string{p.Weapon, p.Armor, p.Boots, p.Helmet} {
		if item, ok := Items[slot]; ok {
			p.PVMax += item.BonusPV
			p.Strength += item.BonusAtk
			p.Defense += item.BonusDef
			p.Reiki += item.BonusReiki
			p.Spd += item.BonusSpeed
		}
	}

	if p.PV > p.PVMax {
		p.PV = p.PVMax
	}
}
