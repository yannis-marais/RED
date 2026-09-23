package ProjetRED

import (
	Menu "ProjetRED/Menu"
	personnage "ProjetRED/Personnage"
	"fmt"
	"time"
)

func Ville1(p *personnage.Character) {

	var saisie string
	for {
		fmt.Println("==== Forteresse du royaume ====")
		fmt.Println("j: Aller parler à Jack (il est attentif, comme s’il attendait ta venue).")
		fmt.Println("e: Partire explorer la ville")
		fmt.Println("m: pour ouvrire le menu")
		fmt.Scanln(&saisie)
		if saisie == "m" {
			Menu.MainMenu(p)
		}
		if saisie == "j" {
			break
		}
		if saisie == "e" {
			observation(p)
		}
		fmt.Println("erreur, veuillez entrer une lettre f, j ou m pour le menu")
	}

}
func jack(p *personnage.Character) {

}

func observation(p *personnage.Character) {
	fmt.Printf("En faisant le tour de la forteresse, %s remarquera que deux maisons se démarquent l'une de l'autre. Les deux semblent appartenir à la royauté, sauf que la seconde est dénuée d'entretien.\n", p.Nom)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Cependant a chaque portail une fil de personne etait presente mais beaucoup elever du coté de la maison mal entretenu ")
	time.Sleep(800 * time.Millisecond)
	fmt.Println("=== choix ===")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("1: Se diriger vers la belle maison ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("2: Se dirigier vers l'autrre maison")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("0: Retourné a la place centrale")

	choice, reponse := Menu.ReadChoice("Votre choix : ")
	if !reponse {
		fmt.Println("Choix invalide !")
	}
	switch choice {
	case 1:
		//MaisonPrince
	case 2:
		//MaisonPrinceDechue
	case 0:
		return
	}

}

// choice, reponse := ReadChoice("Votre choix : ")
// if !reponse {
// 	fmt.Println("Choix invalide !")
// 	continue
// }
// switch choice {
// case 1:
// 	City.Ville1(p)
// case 2:

// %s designe du caractere quand
