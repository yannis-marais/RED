package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var QuestJack int
var OpenMenu func(*personnage.Character)

func Ville1(p *personnage.Character) {
	for {
		fmt.Println("==== Forteresse du royaume ====")
		fmt.Println("1: Aller parler à Jack (il est attentif, comme s'il attendait ta venue).")
		fmt.Println("2: Partir explorer la ville")
		fmt.Println("3: pour ouvrir le menu")
		fmt.Println("0: pour quitter")

		choice, ok := readChoice("Votre choix : ")
		if !ok {
			fmt.Println("Erreur, veuillez entrer 1, 2, 3 ou 0 pour le menu")
			WaitForReturn()
			continue
		}

		switch choice {
		case 3:
			if OpenMenu != nil {
				OpenMenu(p)
				continue
			}
			fmt.Println("Le menu principal n'est pas disponible.")
		case 1:
			jack()
		case 2:
			observation(p)
			return
		case 0:
			return
		default:
			fmt.Println("Erreur, veuillez entrer 1, 2, 3 ou 0 pour le menu")
			WaitForReturn()
		}
	}
}
func jack() {
	fmt.Println("Vous avez parlé à Jack, il vous expliqua que deux princes étaient présents, un héritier et un déchu, les deux groupes cherchant à recruter des membre")
	QuestJack = 1
	WaitForReturn()
}

func readChoice(prompt string) (int, bool) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return 0, false
	}

	value, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0, false
	}

	return value, true
}
func WaitForReturn() {
	fmt.Println("Appuyez sur Entrée pour continuer...")
	_, _ = fmt.Scanln()
}

func observation(p *personnage.Character) {
	fmt.Printf("En faisant le tour de la forteresse, %s remarquera que deux maisons se démarquent l'une de l'autre. Les deux semblent appartenir à la royauté, sauf que la seconde est dénuée d'entretien.\n", p.Nom)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Cependant, à chaque portail, une file de personnes était présente, mais beaucoup plus élevée du côté de la maison mal entretenue ")
	time.Sleep(800 * time.Millisecond)
	fmt.Println("====== choix ======")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("1: Se diriger vers la belle maison ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("2: Se diriger vers l'autre maison")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("0: Retourner à la place centrale")

	choice, ok := readChoice("Votre choix : ")
	if !ok {
		fmt.Println("Choix invalide !")
		observation(p)
	}

	switch choice {
	case 1:
		MaisonPrince(p)
	case 2:
		MaisonPrinceDechue(p)
	case 0:
		Ville1(p)
	default:
		fmt.Println("Choix inconnu. Veuillez réessayer.")
	}
}

func MaisonPrince(p *personnage.Character) {
	if QuestJack == 1 {
		fmt.Println("Avec les informations que vous avez eues, vous vous enrôlez dans la maison du prince héritier")
	} else {
		fmt.Println("En vous dirigeant vers la maison, vous compreniez que c'est un recrutement ; ne sachant pas les conséquences qui en découlent, vous entrez donc au service de l'héritier")
	}
	WaitForReturn()
	VilleH(p)
}
func MaisonPrinceDechue(p *personnage.Character) {
	if QuestJack == 1 {
		fmt.Println("Avec les informations que vous avez eues, vous vous enrôlez dans la maison du prince déchu")
	} else {
		fmt.Println("En vous dirigeant vers la maison, vous compreniez que c'est un recrutement ; ne sachant pas les conséquences qui en découlent, vous entrez donc au service du déchu")
	}
	WaitForReturn()
	VilleD(p)
}
