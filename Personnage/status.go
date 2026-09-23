package ProjetRED

import "fmt"

func (p *Character) UpdateEffects() {
	newEffects := make([]StatusEffect, 0, len(p.Effects))

	for _, effect := range p.Effects {
		p.PV -= effect.Damage
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		if p.PV < 0 {
			p.PV = 0
		}

		if effect.Damage > 0 {
			fmt.Println(p.Nom, "subit", effect.Damage, "dégâts de", effect.Name, "PV :", p.PV)
		} else {
			fmt.Println(p.Nom, "récupère", -effect.Damage, "PV grâce à", effect.Name, "PV :", p.PV)
		}

		effect.TimeLeft -= effect.Interval
		if effect.TimeLeft > 0 {
			newEffects = append(newEffects, effect)
		}
	}

	p.Effects = newEffects
}

func (p *Character) UpdateCooldowns() {
	for skillName, cooldown := range p.Cooldowns {
		if cooldown > 0 {
			p.Cooldowns[skillName] = cooldown - 1
		}
	}
}