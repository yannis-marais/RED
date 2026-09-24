package ProjetRED

import personnage "ProjetRED/Personnage"

func GameTick(p *personnage.Character) {
	if p == nil {
		return
	}
	p.UpdateEffects()
	p.UpdateCooldowns()
}
