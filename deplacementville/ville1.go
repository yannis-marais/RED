package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"fmt"
	"time"
)

var QuestJack int
var OpenMenu func(*personnage.Character)

func Ville1(p *personnage.Character) {

	var saisie string
	for {
		fmt.Println("==== Forteresse du royaume ====")
		fmt.Println("j: Aller parler à Jack (il est attentif, comme s'il attendait ta venue).")
		fmt.Println("e: Partir explorer la ville")
		fmt.Println("m: pour ouvrir le menu")
		fmt.Println("0: pour quitter")
		fmt.Scanln(&saisie)
		if saisie == "m" {
			if OpenMenu != nil {
				OpenMenu(p)
				return
			}
			fmt.Println("Le menu principal n'est pas disponible.")
		}
		if saisie == "j" {
			jack()
		}
		if saisie == "e" {
			observation(p)
			return
		}
		if saisie == "0" {
			return
		}
		fmt.Println("Erreur, veuillez entrer une lettre e, j ou m pour le menu")
		fmt.Println("Appuyez sur Entrée pour continuer...")
		_, _ = fmt.Scanln()
	}

}
func jack() {
	fmt.Println("Vous avez parlé à Jack, il vous expliqua que deux princes étaient présents, un héritier et un déchu, les deux groupes cherchant")
	QuestJack = 1
}

func observation(p *personnage.Character) {
	fmt.Printf("En faisant le tour de la forteresse, %s remarquera que deux maisons se démarquent l'une de l'autre. Les deux semblent appartenir à la royauté, sauf que la seconde est dénuée d'entretien.\n", p.Nom)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Cependant, à chaque portail, une file de personnes était présente, mais beaucoup plus élevée du côté de la maison mal entretenue ")
	time.Sleep(800 * time.Millisecond)
	fmt.Println("=== choix ===")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("1: Se diriger vers la belle maison ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("2: Se diriger vers l'autre maison")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("0: Retourner à la place centrale")

	fmt.Print("Votre choix : ")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Choix invalide !")
		return
	}
	switch choice {
	case 1:
		MaisonPrince(p)
	case 2:
		MaisonPrinceDechue(p)
	case 0:
		Ville1(p)
	}
}

func MaisonPrince(p *personnage.Character) {
	if QuestJack == 1 {
		fmt.Printf("Avec les informations que vous avez eues, vous vous enrôlez dans la maison du prince héritier")
	} else {
		fmt.Printf("En vous dirigeant vers la maison, vous compreniez que c'est un recrutement ; ne sachant pas les conséquences qui en découlent, vous entrez donc au service de l'héritier")
	}
	VilleH(p)
}
func MaisonPrinceDechue(p *personnage.Character) {
	if QuestJack == 1 {
		fmt.Printf("Avec les informations que vous avez eues, vous vous enrôlez dans la maison du prince déchu")
	} else {
		fmt.Printf("En vous dirigeant vers la maison, vous compreniez que c'est un recrutement ; ne sachant pas les conséquences qui en découlent, vous entrez donc au service du déchu")
	}
	VilleD(p)
}
