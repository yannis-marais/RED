package ProjetRED

import (
	"fmt"
	personnage "ProjetRED/Personnage"
)

func UpdateEffects(p *personnage.Character) {
	newEffects := []personnage.StatusEffect{}

	for i := range p.Effects {
		e := &p.Effects[i]
		if e.TimeLeft > 0 {
			p.PV -= e.Damage
			if p.PV < 0 {
				p.PV = 0
			}

			fmt.Println(p.Nom, "subit", e.Damage, "dégâts de poison ! PV :", p.PV)

			e.TimeLeft -= e.Interval

			if e.TimeLeft > 0 {
				newEffects = append(newEffects, *e)
			}
		}
	}

	p.Effects = newEffects
}
