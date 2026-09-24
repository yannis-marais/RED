package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
	"time"
)

func GameTick(p *personnage.Character) {
	p.UpdateEffects()
	p.UpdateCooldowns()
}

func StartCombat(player *personnage.Character, monster *enemies.MONSTER) {

	for {
		time.Sleep(time.Second)

		GameTick(player)

		// Si un des deux est mort, on arrête
		if player.PV <= 0 || monster.PV <= 0 {
			break
		}
	}
}
