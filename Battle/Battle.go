package ProjetRED

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"

	Equipement "ProjetRED/Equipement"
	Menu "ProjetRED/Menu"
	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func isDead(p *personnage.Character) bool {
	return p == nil || p.PV <= 0
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

func renderCombatMenu(p personnage.Character, enemyName string, enemyHP, enemyMaxHP, enemyRHP, enemyRMaxHP int) string {
	if p.PV < 0 {
		p.PV = 0
	}
	if enemyHP < 0 {
		enemyHP = 0
	}
	if enemyRHP < 0 {
		enemyRHP = 0
	}

	const width = 52
	line := strings.Repeat("─", width)
	playerBar := lifeBar(p.PV, p.PVMax, 18)
	enemyBar := lifeBar(enemyHP, enemyMaxHP, 18)
	enemyRBar := lifeBar(enemyRHP, enemyRMaxHP, 18)

	var sb strings.Builder
	fmt.Fprintf(&sb, "\n╭%s╮\n", line)
	fmt.Fprintf(&sb, "│ %-48s │\n", "MENU DE COMBAT")
	fmt.Fprintf(&sb, "├%s┤\n", line)
	fmt.Fprintf(&sb, "│ %-48s │\n", fmt.Sprintf("%s : %d/%d %s", p.Nom, p.PV, p.PVMax, playerBar))
	fmt.Fprintf(&sb, "│ %-48s │\n", fmt.Sprintf("Ennemi PV : %s %d/%d %s", enemyName, enemyHP, enemyMaxHP, enemyBar))
	fmt.Fprintf(&sb, "│ %-48s │\n", fmt.Sprintf("Ennemi PVR : %s %d/%d %s", enemyName, enemyRHP, enemyRMaxHP, enemyRBar))
	fmt.Fprintf(&sb, "├%s┤\n", line)
	fmt.Fprintf(&sb, "│ %-48s │\n", "1. Attaque physique")
	fmt.Fprintf(&sb, "│ %-48s │\n", "2. Attaque spirituelle")
	fmt.Fprintf(&sb, "│ %-48s │\n", "3. Skills")
	fmt.Fprintf(&sb, "│ %-48s │\n", "4. Make a Wish")
	fmt.Fprintf(&sb, "│ %-48s │\n", "5. Inventaire")
	fmt.Fprintf(&sb, "│ %-48s │\n", "6. Défendre")
	fmt.Fprintf(&sb, "│ %-48s │\n", "0. Fuir")
	fmt.Fprintf(&sb, "╰%s╯\n", line)
	return sb.String()
}

func applySkill(p *personnage.Character, monstre *enemies.MONSTER, skillName string) bool {
	if p == nil || monstre == nil {
		return false
	}

	if cooldown, exists := p.Cooldowns[skillName]; exists && cooldown > 0 {
		fmt.Printf("Le skill %s est en cooldown (%d tours restants).\n", skillName, cooldown)
		return false
	}

	skillDef, okSkill := Equipement.SkillList[skillName]
	if !okSkill {
		fmt.Println("Ce skill n'existe pas dans la liste.")
		return false
	}

	dmg := Equipement.CalculateSkillDamage(*p, skillDef)
	if skillDef.Type == "Magic" {
		if monstre.PVR > 0 {
			monstre.PVR -= dmg
			if monstre.PVR < 0 {
				monstre.PVR = 0
			}
		}
		fmt.Printf("\n%s lance %s et inflige %d dégâts spirituels à %s\n", p.Nom, skillName, dmg, monstre.NOM)
		fmt.Printf("%s : PVR %d/%d\n", monstre.NOM, monstre.PVR, monstre.PVMAXR)
	} else {
		if monstre.PV > 0 {
			monstre.PV -= dmg
			if monstre.PV < 0 {
				monstre.PV = 0
			}
		}
		fmt.Printf("\n%s utilise %s et inflige %d dégâts physiques à %s\n", p.Nom, skillName, dmg, monstre.NOM)
		fmt.Printf("%s : PV %d/%d\n", monstre.NOM, monstre.PV, monstre.PVMax)
	}

	if skillDef.Cooldown > 0 {
		p.Cooldowns[skillName] = skillDef.Cooldown
	}
	return true
}

func usePlayerSkill(p *personnage.Character, monstre *enemies.MONSTER) bool {
	if p == nil || monstre == nil {
		return false
	}

	if len(p.Skills) == 0 {
		fmt.Println("Tu n'as aucun skill débloqué. Trouve un SkillBook pendant l'aventure.")
		return false
	}

	names := make([]string, 0, len(p.Skills))
	for name := range p.Skills {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("=== Skills débloqués ===")
	for i, name := range names {
		skill := p.Skills[name]
		cooldown := p.Cooldowns[name]
		fmt.Printf("%d. %s (%s) - CD:%d\n", i+1, name, skill.Type, cooldown)
	}

	choice, ok := Menu.ReadChoice("Choisissez un skill : ")
	if !ok || choice < 1 || choice > len(names) {
		fmt.Println("Choix de skill invalide.")
		return false
	}

	return applySkill(p, monstre, names[choice-1])
}

func characterTurn(p *personnage.Character, monstre *enemies.MONSTER) {
	if p == nil || monstre == nil {
		fmt.Println("Combat impossible : personnage ou monstre invalide.")
		return
	}

	for {
		fmt.Print(renderCombatMenu(*p, monstre.NOM, monstre.PV, monstre.PVMax, monstre.PVR, monstre.PVMAXR))
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
			if usePlayerSkill(p, monstre) {
				return
			}
			continue

		case 4:
			makeAWish(p, monstre)
			return

		case 5:
			Menu.ManageInventory(*p)
			Menu.WaitForReturn()
			continue

		case 6:
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
