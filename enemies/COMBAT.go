package enemies

import (
	"fmt"

	personnage "ProjetRED/Personnage"
)

type MonsterPattern func(tour int, monstre *MONSTER, perso *personnage.Character)

func GoblinPattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength
	if tour%3 == 0 {
		dgt *= 2
	}

	perso.PV -= dgt
	fmt.Printf("%s frappe avec sa massue ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

func SkeletonPattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength + 2
	if tour%2 == 0 {
		dgt += 20
	}

	perso.PV -= dgt
	fmt.Printf("%s attaque avec son gros os ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

func TrollPattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength + 4
	if tour%2 == 0 {
		dgt += 6
	}

	perso.PV -= dgt
	fmt.Printf("%s fracasse ton gros crane avec sa massue ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

func VouivrePattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength + 2
	if tour%3 == 0 {
		dgt += 5
	}

	perso.PV -= dgt
	fmt.Printf("%s crache du venin ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

func LoupGarouPattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength + 2
	if tour%2 == 1 {
		dgt += 6
	}

	perso.PV -= dgt
	fmt.Printf("%s bondit et mord profondément  ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

func ZombiePattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength
	if tour%4 == 0 {
		dgt += 2
	}

	perso.PV -= dgt
	fmt.Printf("%s donne un coup lent mais tenace ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

func OrcPattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength + 5
	if tour%3 == 0 {
		dgt += 2
	}

	perso.PV -= dgt
	fmt.Printf("%s donne un coup de hache brutal ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

func DragonPattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	dgt := monstre.Strength + 8
	if tour%2 == 0 {
		dgt += 10
	}

	perso.PV -= dgt
	fmt.Printf("%s te carbonise la gueule ! %d dégâts à %s\n", monstre.NOM, dgt, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}

var monsterPatterns = map[string]MonsterPattern{
	"gobelin":                GoblinPattern,
	"Gobelin_d_entrainement": GoblinPattern,
	"squelette":              SkeletonPattern,
	"SKELETON":               SkeletonPattern,
	"troll":                  TrollPattern,
	"vouivre":                VouivrePattern,
	"loup-garou":             LoupGarouPattern,
	"zombie":                 ZombiePattern,
	"orc":                    OrcPattern,
	"dragon":                 DragonPattern,
	"Le B.":                  DragonPattern,
	"Le Ant":                 DragonPattern,
	"Le L":                   DragonPattern,
	"Le RELOU":               DragonPattern,
}

func MonsterAttackPattern(tour int, monstre *MONSTER, perso *personnage.Character) {
	if monstre == nil || perso == nil {
		return
	}

	if pattern, ok := monsterPatterns[monstre.NOM]; ok {
		pattern(tour, monstre, perso)
		return
	}

	defaultDamage := monstre.Strength
	if tour%3 == 0 {
		defaultDamage *= 2
	}

	perso.PV -= defaultDamage
	fmt.Printf("%s attaque ! %d dégâts à %s\n", monstre.NOM, defaultDamage, perso.Nom)
	fmt.Printf("%s : PV %d/%d\n", perso.Nom, perso.PV, perso.PVMax)
}
