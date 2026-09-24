package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func readInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		value, err := strconv.Atoi(text)
		if err == nil {
			return value
		}
		fmt.Println("Saisie invalide.")
	}
}

func renderBattle(player personnage.Character, enemy *enemies.MONSTER) {
	fmt.Println("\n=== COMBAT ===")
	fmt.Printf("Joueur : %s | PV %d/%d | Force %d | Reiki %d | Vitesse %d\n",
		player.Nom, player.PV, player.PVMax, player.Strength, player.Reiki, player.Spd)
	fmt.Printf("Ennemi : %s | PV %d/%d | PVR %d/%d | Force %d | Vitesse %d\n",
		enemy.NOM, enemy.PV, enemy.PVMax, enemy.PVR, enemy.PVMAXR, enemy.Strength, enemy.Spd)
	fmt.Println("1. Attaque physique")
	fmt.Println("2. Attaque spirituelle")
	fmt.Println("3. Défendre")
	fmt.Println("0. Fuir")
}

func runBattleDemo() {
	player := personnage.CharacterCreation("Jean", personnage.Classes["Ronin"])
	player.PV = player.PVMax
	player.Spd = 25

	enemy, ok := enemies.NewMonsterByName("squelette")
	if !ok {
		fmt.Println("Impossible de générer le monstre demandé : squelette")
		return
	}
	enemy.PV = enemy.PVMax
	enemy.PVR = enemy.PVMAXR
	fmt.Printf("Tu rencontres un %s !\n\n", enemy.NOM)

	for player.PV > 0 && enemy.PV > 0 && enemy.PVR > 0 {
		renderBattle(player, enemy)
		choice := readInt("Choix : ")

		switch choice {
		case 1:
			multiplier := 1.0
			if rand.Intn(100) < 25 {
				multiplier = 1.8
				fmt.Println("Coup critique !")
			}
			dmg := int(float64(player.Strength) * multiplier)
			if dmg <= 0 {
				dmg = 1
			}
			enemy.PV -= dmg
			if enemy.PV < 0 {
				enemy.PV = 0
			}
			fmt.Printf("Tu frappes pour %d dégâts.\n", dmg)
		case 2:
			multiplier := 1.0
			if rand.Intn(100) < 30 {
				multiplier = 1.7
				fmt.Println("Attaque spirituelle critique !")
			}
			dmg := int(float64(player.Reiki) * multiplier)
			if dmg <= 0 {
				dmg = 1
			}
			enemy.PVR -= dmg
			if enemy.PVR < 0 {
				enemy.PVR = 0
			}
			fmt.Printf("Ton énergie spirituelle inflige %d dégâts.\n", dmg)
		case 3:
			fmt.Println("Tu prends une position défensive. Le monstre attaque moins fort ce tour.")
		case 0:
			fmt.Println("Tu fuis le combat.")
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if enemy.PV <= 0 || enemy.PVR <= 0 {
			fmt.Printf("Tu as vaincu le %s !\n", enemy.NOM)
			return
		}

		if choice == 3 {
			fmt.Println("Le monstre te frappe avec moins de force.")
			player.PV -= max(0, enemy.Strength/2)
		} else {
			enemyDmg := enemy.Strength + rand.Intn(8)
			player.PV -= enemyDmg
			if player.PV < 0 {
				player.PV = 0
			}
			fmt.Printf("Le %s te frappe pour %d dégâts.\n", enemy.NOM, enemyDmg)
		}

		if player.PV <= 0 {
			fmt.Println("Tu as perdu le combat.")
			return
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	runBattleDemo()
}
