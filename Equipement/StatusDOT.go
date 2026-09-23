package ProjetRED

import (
	"fmt"

	personnage "ProjetRED/Personnage"
)

func ApplyPoisonDOT(p *personnage.Character) {
	effect := personnage.StatusEffect{
		Name:     "Poison",
		Damage:   10,
		Duration: 3,
		Interval: 1,
		TimeLeft: 3,
	}
	p.Effects = append(p.Effects, effect)
	fmt.Println(p.Nom, "est empoisonné !")
}
func ApplyBreadRegen(p *personnage.Character) {
	effect := personnage.StatusEffect{
		Name:     "Regen du Pain",
		Damage:   -5, // -5 = soin de 5 PV
		Duration: 5,
		Interval: 1,
		TimeLeft: 5,
	}

	p.Effects = append(p.Effects, effect)
	fmt.Println(p.Nom, "mange du pain et commence à se régénérer !")
}
