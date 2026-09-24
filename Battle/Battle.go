package ProjetRED

import (
	"fmt"
	"math/rand"
	"strings"

	Equipement "ProjetRED/Equipement"
	Menu "ProjetRED/Menu"
	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func isDead(p *personnage.Character) bool {
	return p == nil || p.PV <= 0
}

func giveMonsterLoot(p *personnage.Character, monstre *enemies.MONSTER, log *strings.Builder) {
	if p == nil || monstre == nil {
		return
	}

	for _, materialName := range enemies.RollLoot(monstre) {
		if Equipement.AddMaterial(p, materialName) {
			if log != nil {
				fmt.Fprintf(log, "%s récupère 1 %s.\n", p.Nom, materialName)
			} else {
				fmt.Printf("%s récupère 1 %s.\n", p.Nom, materialName)
			}
		}
	}
}

func lifeBar(current, max, width int) string {
	if max <= 0 || width <= 0 {
		return "[--]"
	}
	filled := int(float64(current) / float64(max) * float64(width))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	return "[" + strings.Repeat("█", filled) + strings.Repeat("░", width-filled) + "]"
}

func renderCombatMenu(p personnage.Character, enemyName string, enemyHP, enemyMaxHP int) string {
	if p.PV < 0 {
		p.PV = 0
	}
	if enemyHP < 0 {
		enemyHP = 0
	}

	const width = 52
	line := strings.Repeat("─", width)
	playerBar := lifeBar(p.PV, p.PVMax, 18)
	enemyBar := lifeBar(enemyHP, enemyMaxHP, 18)

	var sb strings.Builder
	fmt.Fprintf(&sb, "\n╭%s╮\n", line)
	fmt.Fprintf(&sb, "│ %-48s │\n", "MENU DE COMBAT")
	fmt.Fprintf(&sb, "├%s┤\n", line)
	fmt.Fprintf(&sb, "│ %-48s │\n", fmt.Sprintf("%s : %d/%d %s", p.Nom, p.PV, p.PVMax, playerBar))
	fmt.Fprintf(&sb, "│ %-48s │\n", fmt.Sprintf("Ennemi : %s %d/%d %s", enemyName, enemyHP, enemyMaxHP, enemyBar))
	fmt.Fprintf(&sb, "├%s┤\n", line)
	fmt.Fprintf(&sb, "│ %-48s │\n", "1. Attaque basique")
	fmt.Fprintf(&sb, "│ %-48s │\n", "2. Attaque spéciale")
	fmt.Fprintf(&sb, "│ %-48s │\n", "3. Skill / Magie")
	fmt.Fprintf(&sb, "│ %-48s │\n", "4. Inventaire")
	fmt.Fprintf(&sb, "│ %-48s │\n", "5. Défendre")
	fmt.Fprintf(&sb, "│ %-48s │\n", "0. Fuir")
	fmt.Fprintf(&sb, "╰%s╯\n", line)
	return sb.String()
}

func characterTurn(p *personnage.Character, monstre *enemies.MONSTER) {
	if p == nil || monstre == nil {
		fmt.Println("Combat impossible : personnage ou monstre invalide.")
		return
	}

	for {
		fmt.Print(renderCombatMenu(*p, monstre.NOM, monstre.PV, monstre.PVMax))
		choice, ok := Menu.ReadChoice("Votre choix : ")
		if !ok {
			fmt.Println("Choix invalide.")
			continue
		}

		switch choice {
		case 1:
			degats := p.Strength
			if monstre.PV > 0 {
				monstre.PV -= degats
				if monstre.PV < 0 {
					monstre.PV = 0
				}
			}
			fmt.Println("\nVous utilisez Attaque basique.")
			fmt.Printf("%s inflige %d dégâts à %s\n", p.Nom, degats, monstre.NOM)
			fmt.Printf("%s : PV %d/%d\n", monstre.NOM, monstre.PV, monstre.PVMax)
			return

		case 2:
			degats := p.Reiki
			if monstre.PVR > 0 {
				monstre.PVR -= degats
				if monstre.PVR < 0 {
					monstre.PVR = 0
				}
			}
			fmt.Println("\nVous utilisez Attaque spéciale (Reiki).")
			fmt.Printf("%s inflige %d dégâts à %s\n", p.Nom, degats, monstre.NOM)
			fmt.Printf("%s : PV %d/%d\n", monstre.NOM, monstre.PVR, monstre.PVMAXR)
			return

		case 3:
			makeAWish(p, monstre)
			return

		case 4:
			Menu.ManageInventory(*p)
			Menu.WaitForReturn()
			continue

		case 5:
			fmt.Println("\nVous prenez une position défensive. Vous attendez le prochain coup.")
			return

		case 0:
			fmt.Println("\nVous abandonnez le combat et reculez.")
			monstre.PV = 0
			monstre.PVR = 0
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func makeAWish(p *personnage.Character, monstre *enemies.MONSTER) {
	fmt.Println("\nVous invoquez Make a Wish...")

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
		monstre.PVR = 0
	}
}

func runDemoAction(p *personnage.Character, monstre *enemies.MONSTER, choice int, log *strings.Builder) {
	switch choice {
	case 1:
		degats := p.Strength
		monstre.PV -= degats
		if monstre.PV < 0 {
			monstre.PV = 0
		}
		fmt.Fprintf(log, "%s attaque %s pour %d dégâts.\n", p.Nom, monstre.NOM, degats)
		fmt.Fprintf(log, "%s : PV %d/%d\n", monstre.NOM, monstre.PV, monstre.PVMax)
	case 2:
		degats := p.Reiki
		monstre.PVR -= degats
		if monstre.PVR < 0 {
			monstre.PVR = 0
		}
		fmt.Fprintf(log, "%s utilise un sort spirituel sur %s pour %d dégâts.\n", p.Nom, monstre.NOM, degats)
		fmt.Fprintf(log, "%s : PV %d/%d\n", monstre.NOM, monstre.PVR, monstre.PVMAXR)
	case 3:
		makeAWish(p, monstre)
		fmt.Fprintf(log, "%s lance Make a Wish.\n", p.Nom)
	case 5:
		fmt.Fprintf(log, "%s se défend et attend le prochain coup.\n", p.Nom)
	default:
		fmt.Fprintf(log, "%s hésite et perd son tour.\n", p.Nom)
	}
}

func RunDemoBattle(p *personnage.Character, monstre *enemies.MONSTER, actions []int) string {
	if p == nil || monstre == nil {
		return ""
	}

	var log strings.Builder
	log.WriteString("=== Combat de démonstration ===\n")

	for tour := 1; tour <= 12; tour++ {
		log.WriteString(fmt.Sprintf("--- Tour %d ---\n", tour))

		if p.Spd >= monstre.Spd {
			choice := 1
			if len(actions) > 0 {
				choice = actions[0]
				actions = actions[1:]
			}
			runDemoAction(p, monstre, choice, &log)
			if enemies.IsMonsterDead(monstre) && !isDead(p) {
				log.WriteString(fmt.Sprintf("%s est vaincu !\n", monstre.NOM))
				giveMonsterLoot(p, monstre, &log)
				return log.String()
			}

			enemies.MonsterAttackPattern(tour, monstre, p)
			log.WriteString(fmt.Sprintf("%s : PV %d/%d\n", p.Nom, p.PV, p.PVMax))
			if isDead(p) {
				log.WriteString(fmt.Sprintf("%s est vaincu !\n", p.Nom))
				return log.String()
			}
		} else {
			enemies.MonsterAttackPattern(tour, monstre, p)
			log.WriteString(fmt.Sprintf("%s : PV %d/%d\n", p.Nom, p.PV, p.PVMax))
			if isDead(p) {
				log.WriteString(fmt.Sprintf("%s est vaincu !\n", p.Nom))
				return log.String()
			}

			choice := 1
			if len(actions) > 0 {
				choice = actions[0]
				actions = actions[1:]
			}
			runDemoAction(p, monstre, choice, &log)
			if enemies.IsMonsterDead(monstre) && !isDead(p) {
				log.WriteString(fmt.Sprintf("%s est vaincu !\n", monstre.NOM))
				giveMonsterLoot(p, monstre, &log)
				return log.String()
			}
		}
	}

	if isDead(p) {
		log.WriteString(fmt.Sprintf("%s est vaincu !\n", p.Nom))
	} else if enemies.IsMonsterDead(monstre) && !isDead(p) {
		log.WriteString(fmt.Sprintf("%s est vaincu !\n", monstre.NOM))
		giveMonsterLoot(p, monstre, &log)
	} else {
		log.WriteString("Le combat se termine sans vainqueur clair.\n")
	}

	return log.String()
}

func trainingFight(p *personnage.Character, monstre *enemies.MONSTER) {
	tour := 1

	for {
		fmt.Println("\n=== TOUR", tour, "===")

		if p.Spd >= monstre.Spd {
			characterTurn(p, monstre)
			if enemies.IsMonsterDead(monstre) && !isDead(p) {
				fmt.Println(monstre.NOM, "est vaincu !")
				giveMonsterLoot(p, monstre, nil)
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
			if enemies.IsMonsterDead(monstre) && !isDead(p) {
				fmt.Println(monstre.NOM, "est vaincu !")
				giveMonsterLoot(p, monstre, nil)
				break
			}
		}

		tour++
	}
}
