package ProjetRED

import (
	Equipement "ProjetRED/Equipement"
	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func isDead(p *personnage.Character) bool {
	if p == nil || p.PV > 0 {
		return false
	}

	p.PV = p.PVMax / 2
	fmt.Printf("%s tombe au combat... mais se relève avec %d/%d PV !\n", p.Nom, p.PV, p.PVMax)
	return true
}

func GiveMonsterLoot(p *personnage.Character, monstre *enemies.MONSTER) []string {
	if p == nil || monstre == nil || !enemies.IsMonsterDead(monstre) {
		return nil
	}

	dropped := make([]string, 0)
	for _, lootName := range enemies.RollLoot(monstre) {
		if Equipement.AddLoot(p, lootName) {
			dropped = append(dropped, lootName)
		}
	}
	return dropped
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

func padCombatLine(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}
	return text + strings.Repeat(" ", width-len(text))
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

	lines := []string{
		"MENU DE COMBAT",
		fmt.Sprintf("%s : %d/%d %s", p.Nom, p.PV, p.PVMax, playerBar),
		fmt.Sprintf("Ennemi PV : %s %d/%d %s", enemyName, enemyHP, enemyMaxHP, enemyBar),
		fmt.Sprintf("Ennemi PVR : %s %d/%d %s", enemyName, enemyRHP, enemyRMaxHP, enemyRBar),
		"1. Attaque physique",
		"2. Attaque spirituelle",
		"3. Skills",
		"4. Make a Wish",
		"5. Inventaire",
		"6. Défendre",
		"0. Fuir",
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "\n╭%s╮\n", line)
	for _, l := range lines[:1] {
		fmt.Fprintf(&sb, "│ %-s │\n", padCombatLine(l, 48))
	}
	fmt.Fprintf(&sb, "├%s┤\n", line)
	for _, l := range lines[1:] {
		fmt.Fprintf(&sb, "│ %-50s │\n", padCombatLine(l, 48))
	}
	fmt.Fprintf(&sb, "╰%s╯\n", line)
	return sb.String()
}

func performQTE() (float64, string) {
	const (
		targetMin = 40
		targetMax = 60
		step      = 20 * time.Millisecond
	)

	fmt.Println("\nQTE : appuie sur Entrée quand le X est au centre !")
	fmt.Println("       0         25         50         75        100")
	fmt.Println("       |----------|----------|----------|----------|")

	start := time.Now()
	pos := 0
	direction := 1
	reader := bufio.NewReader(os.Stdin)
	inputDone := make(chan struct{})

	go func() {
		_, _ = reader.ReadString('\n')
		close(inputDone)
	}()

	for time.Since(start) < 2*time.Second {
		pos += direction
		if pos <= 0 || pos >= 100 {
			direction *= -1
			pos += direction
		}

		bar := make([]rune, 100)
		for i := range bar {
			bar[i] = '-'
		}
		bar[pos] = 'X'
		for i := targetMin; i <= targetMax; i++ {
			if i >= 0 && i < len(bar) {
				bar[i] = '='
			}
		}
		if pos >= targetMin && pos <= targetMax {
			fmt.Printf("\r[%s]  CIBLE  < %d >", string(bar), pos)
		} else {
			fmt.Printf("\r[%s]  < %d >", string(bar), pos)
		}

		select {
		case <-inputDone:
			if pos >= targetMin && pos <= targetMax {
				fmt.Println("\n* CRIT *")
				return 2.0, "CRIT"
			}
			if pos >= targetMin-15 && pos <= targetMax+15 {
				fmt.Println("\n* BON *")
				return 1.5, "BON"
			}
			fmt.Println("\n* RATÉ *")
			return 0.8, "RATÉ"
		default:
		}

		time.Sleep(step)
	}

	fmt.Println("\n* RATÉ *")
	return 0.8, "RATÉ"
}

func applyAttackDamage(p *personnage.Character, monstre *enemies.MONSTER, baseDamage int, target string, multiplier float64) {
	if p == nil || monstre == nil {
		return
	}

	dmg := int(float64(baseDamage) * multiplier)
	if dmg <= 0 {
		dmg = 0
	}

	switch target {
	case "PV":
		if monstre.PV > 0 {
			monstre.PV -= dmg
			if monstre.PV < 0 {
				monstre.PV = 0
			}
		}
		fmt.Printf("%s inflige %d dégâts physiques à %s\n", p.Nom, dmg, monstre.NOM)
		fmt.Printf("%s : PV %d/%d\n", monstre.NOM, monstre.PV, monstre.PVMax)
	case "PVR":
		if monstre.PVR > 0 {
			monstre.PVR -= dmg
			if monstre.PVR < 0 {
				monstre.PVR = 0
			}
		}
		fmt.Printf("%s inflige %d dégâts spirituels à %s\n", p.Nom, dmg, monstre.NOM)
		fmt.Printf("%s : PVR %d/%d\n", monstre.NOM, monstre.PVR, monstre.PVMAXR)
	}
}

func applySkillWithMultiplier(p *personnage.Character, monstre *enemies.MONSTER, skillName string, multiplier float64) bool {
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

	baseDamage := Equipement.CalculateSkillDamage(*p, skillDef)
	if skillDef.Type == "Magic" {
		applyAttackDamage(p, monstre, baseDamage, "PVR", multiplier)
	} else {
		applyAttackDamage(p, monstre, baseDamage, "PV", multiplier)
	}

	if skillDef.Cooldown > 0 {
		p.Cooldowns[skillName] = skillDef.Cooldown
	}
	return true
}

func applySkill(p *personnage.Character, monstre *enemies.MONSTER, skillName string) bool {
	multiplier, _ := performQTE()
	return applySkillWithMultiplier(p, monstre, skillName, multiplier)
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

	choice, ok := ReadChoice("Choisissez un skill : ")
	if !ok || choice < 1 || choice > len(names) {
		fmt.Println("Choix de skill invalide.")
		return false
	}

	return applySkill(p, monstre, names[choice-1])
}

func RenderCombatMenu(p personnage.Character, enemyName string, enemyHP, enemyMaxHP, enemyRHP, enemyRMaxHP int) string {
	return renderCombatMenu(p, enemyName, enemyHP, enemyMaxHP, enemyRHP, enemyRMaxHP)
}

func PerformQTE() (float64, string) {
	return performQTE()
}

func ApplyAttackDamage(p *personnage.Character, monstre *enemies.MONSTER, baseDamage int, target string, multiplier float64) {
	applyAttackDamage(p, monstre, baseDamage, target, multiplier)
}

func UsePlayerSkill(p *personnage.Character, monstre *enemies.MONSTER) bool {
	return usePlayerSkill(p, monstre)
}

func MakeAWish(p *personnage.Character, monstre *enemies.MONSTER) {
	makeAWish(p, monstre)
}

func StartCombat(player *personnage.Character, monster *enemies.MONSTER) bool {
	if player == nil || monster == nil {
		fmt.Println("Combat impossible : personnage ou monstre invalide.")
		return false
	}
	if player.PV <= 0 || monster.PV <= 0 {
		fmt.Println("Combat impossible : un combattant est déjà hors jeu.")
		return false
	}

	for player.PV > 0 && !enemies.IsMonsterDead(monster) {

		// Si le monstre est plus rapide, il attaque avant que tu n'agisses
		if monster.Spd > player.Spd {
			monsterDamage := monster.Strength
			if monsterDamage < 0 {
				monsterDamage = 0
			}
			player.PV -= monsterDamage
			if player.PV < 0 {
				player.PV = 0
			}
			fmt.Printf("%s (plus rapide) te frappe pour %d dégâts.\n", monster.NOM, monsterDamage)
			fmt.Printf("%s : PV %d/%d\n", player.Nom, player.PV, player.PVMax)

			isDead(player)
		}

		fmt.Print(renderCombatMenu(*player, monster.NOM, monster.PV, monster.PVMax, monster.PVR, monster.PVMAXR))

		choice, ok := ReadChoice("Votre choix : ")
		if !ok {
			fmt.Println("Choix invalide.")
			continue
		}

		switch choice {
		case 1:
			multiplier, _ := performQTE()
			ApplyAttackDamage(player, monster, player.Strength, "PV", multiplier)
		case 2:
			multiplier, _ := performQTE()
			ApplyAttackDamage(player, monster, player.Reiki, "PVR", multiplier)
		case 3:
			if !UsePlayerSkill(player, monster) {
				continue
			}
		case 4:
			MakeAWish(player, monster)
		case 5:
			ManageInventory(player)
			WaitForReturn()
			continue
		case 6:
			fmt.Println("Tu prends une position défensive et attends le prochain coup.")
		case 0:
			fmt.Println("Tu fuis le combat.")
			return false
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if enemies.IsMonsterDead(monster) {
			fmt.Println("Victoire ! Le monstre est vaincu.")

			gold := uint(rand.Intn(51) + 50) // 50 à 100 inclus
			player.Purse += gold
			fmt.Printf("Vous récupérez %d pièces d'or.\n", gold)

			for _, materialName := range GiveMonsterLoot(player, monster) {
				fmt.Println("Drop récupéré :", materialName)
			}
			return true
		}

		// Le monstre attaque après toi seulement s'il n'a pas déjà joué ce tour
		if monster.Spd <= player.Spd && choice != 5 && choice != 6 && choice != 0 {
			monsterDamage := monster.Strength
			if monsterDamage < 0 {
				monsterDamage = 0
			}
			player.PV -= monsterDamage
			if player.PV < 0 {
				player.PV = 0
			}
			fmt.Printf("%s te frappe pour %d dégâts.\n", monster.NOM, monsterDamage)
			fmt.Printf("%s : PV %d/%d\n", player.Nom, player.PV, player.PVMax)
		}

		isDead(player)
	}

	if player.PV <= 0 {
		fmt.Println("Tu es mort au combat.")
		return false
	}

	fmt.Println("Combat terminé.")
	return true
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
