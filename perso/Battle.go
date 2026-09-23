package ProjetRED

import (
	"fmt"
	"math/rand"

	Menu "ProjetRED/Menu"
	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func isDead(p *personnage.Character) bool {
	return p == nil || p.PV <= 0
}

func characterTurn(p *personnage.Character, monstre *enemies.MONSTER) {
	fmt.Println("Menu")
	fmt.Println("1. Attaque basique (Force)")
	fmt.Println("2. Attaque spéciale (Reiki)")
	fmt.Println("3. Make a Wish")
	fmt.Println("4. Inventaire")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		degats := p.Strength
		monstre.PV -= degats

		fmt.Println("Vous utilisez Attaque basique")
		fmt.Printf("%s inflige %d dégâts à %s\n", p.Nom, degats, monstre.NOM)
		fmt.Printf("%s : PV %d/%d\n", monstre.NOM, monstre.PV, monstre.PVMax)

	case 2:
		degats := p.Reiki
		monstre.PVR -= degats

		fmt.Println("Vous utilisez Attaque spéciale (Reiki)")
		fmt.Printf("%s inflige %d dégâts à %s\n", p.Nom, degats, monstre.NOM)
		fmt.Printf("%s : PV %d/%d\n", monstre.NOM, monstre.PVR, monstre.PVMAXR)

	case 3:
		makeAWish(p, monstre)

	case 4:
		Menu.AccessInventory(*p)
		// à remplacer par une vraie fonction d'utilisation d'objet

	default:
		fmt.Println("Choix invalide")
		characterTurn(p, monstre)
	}

}

func makeAWish(p *personnage.Character, monstre *enemies.MONSTER) {
	fmt.Println("Vous invoquez Make a Wish...")

	tirage := rand.Intn(100)

	switch {
	case tirage < 35:
		degats := rand.Intn(20) + 5
		p.PV -= degats
		fmt.Printf("Le sort se retourne contre vous ! Vous subissez %d dégâts\n", degats)
		fmt.Printf("%s : PV %d/%d\n", p.Nom, p.PV, p.PVMax)

	case tirage < 60:
		degats := rand.Intn(15) + 5
		monstre.PV -= degats
		fmt.Printf("Décharge instable ! %d dégâts infligés à %s\n", degats, monstre.NOM)
		fmt.Printf("%s : PV %d/%d\n", monstre.NOM, monstre.PV, monstre.PVMax)

	case tirage < 80:
		soin := 30
		p.PV += soin
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		fmt.Printf("Vous êtes soigné de %d PV\n", soin)
		fmt.Printf("%s : PV %d/%d\n", p.Nom, p.PV, p.PVMax)

	case tirage < 95:
		effetStatus := personnage.StatusEffect{
			Name:     "Poison",
			Damage:   5,
			Duration: 3,
			Interval: 1,
			TimeLeft: 3,
		}
		p.Effects = append(p.Effects, effetStatus)
		fmt.Println("Un effet de statut vous affecte :", effetStatus.Name)

	default:
		fmt.Println("One Shot ! L'ennemi est anéanti d'un coup")
		monstre.PV = 0
	}
}

func trainingFight(p *personnage.Character, monstre *enemies.MONSTER) {
	tour := 1

	for {
		fmt.Println("--- Tour", tour, "---")

		if p.Spd >= monstre.Spd {
			characterTurn(p, monstre)
			if enemies.IsMonsterDead(monstre) {
				fmt.Println(monstre.NOM, "est vaincu !")
				break
			}

			enemies.MonsterAttackPattern(tour, monstre, p)
			if isDead(p) {
				fmt.Println(p.Nom, "est vaincu !")
				break
			}

		} else {
			enemies.MonsterAttackPattern(tour, monstre, p)
			if isDead(p) {
				fmt.Println(p.Nom, "est vaincu !")
				break
			}

			characterTurn(p, monstre)
			if enemies.IsMonsterDead(monstre) {
				fmt.Println(monstre.NOM, "est vaincu !")
				break
			}
		}

		tour++
	}
}
