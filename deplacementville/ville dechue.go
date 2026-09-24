package ProjetRED

import (
	enemies "ProjetRED/enemies"
	personnage "ProjetRED/Personnage"
	"ProjetRED/world"
	"fmt"
	"time"
)

// DechuProgress suit l'avancée du joueur dans la maison du prince déchu.
// 0: vient d'arriver, 1: a parlé à Renard (peut partir chasser)
var DechuProgress int

func VilleD(p *personnage.Character) {
	Faction = "Dechu"
	if DechuProgress == 2 {
		for {
			fmt.Println("\n==== Maison du Prince Déchu ====")
			fmt.Printf("%s se trouve dans la cour désordonnée mais grouillante de vie de la maison déchue.\n", p.Nom)
			fmt.Println("1: Parler à Renard, meneur officieux des partisans")
			fmt.Println("2: Partir chasser la bête qui rôde près du campement")
			fmt.Println("3: Se Diriger vers la capital")
			fmt.Println("4: pour ouvrir le menu")
			fmt.Println("0: pour quitter")

			choice, ok := readChoice("Votre choix : ")
			if !ok {
				fmt.Println("Erreur, veuillez entrer un choix valide")
				WaitForReturn()
				continue
			}

			switch choice {
			case 4:
				if OpenMenu != nil {
					OpenMenu(p)
					continue
				}
				fmt.Println("Le menu principal n'est pas disponible.")
			case 1:
				renard(p)
			case 2:
				if chasse(p) {
					// La quête de la maison du prince déchu est terminée :
					// on enchaîne directement sur la capitale.
					Ville4(p, "Dechu")
					return
				}
			case 3:
				Ville4(p,"Dechu")
			case 0:
				return
			default:
				fmt.Println("Erreur, veuillez entrer un choix valide")
				WaitForReturn()
			}
		}		

	}else{
		for {
			fmt.Println("\n==== Maison du Prince Déchu ====")
			fmt.Printf("%s se trouve dans la cour désordonnée mais grouillante de vie de la maison déchue.\n", p.Nom)
			fmt.Println("1: Parler à Renard, meneur officieux des partisans")
			fmt.Println("2: Partir chasser la bête qui rôde près du campement")
			fmt.Println("3: pour ouvrir le menu")
			fmt.Println("0: pour quitter")

			choice, ok := readChoice("Votre choix : ")
			if !ok {
				fmt.Println("Erreur, veuillez entrer un choix valide")
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
				renard(p)
			case 2:
				if chasse(p) {
					// La quête de la maison du prince déchu est terminée :
					// on enchaîne directement sur la capitale.
					Ville4(p, "Dechu")
					return
				}
			case 0:
				return
			default:
				fmt.Println("Erreur, veuillez entrer un choix valide")
				WaitForReturn()
			}
		}
	}
}

func renard(p *personnage.Character) {
	if DechuProgress == 0 {
		fmt.Println("Renard vous jauge un instant avant d'éclater de rire. « Un nouveau ! On accueille tout le monde ici, tant qu'on tient debout. »")
		fmt.Println("« Le prince déchu n'a ni titre ni fortune, seulement des gens prêts à se battre pour qu'on les écoute enfin. »")
		fmt.Println("« Si tu veux qu'on te fasse confiance, viens donc chasser avec nous. Y'a une bête qui rôde près du campement. »")
		DechuProgress = 1
	} else {
		fmt.Println("Renard vous fait un clin d'œil. « Toujours partant pour renverser l'ordre établi ? »")
	}
	WaitForReturn()
}

// chasse fait affronter un monstre de ville2 au joueur.
// Elle renvoie true si la quête est validée (victoire), ce qui déclenche
// automatiquement l'enchaînement vers Ville4.
func chasse(p *personnage.Character) bool {
	fmt.Printf("%s rejoint un groupe de recrues déjà prêtes, torches et bâtons en main.\n", p.Nom)
	time.Sleep(500 * time.Millisecond)

	if DechuProgress < 1 {
		fmt.Println("Les recrues vous regardent avec méfiance. « On ne part pas avec n'importe qui. Va d'abord voir Renard. »")
		WaitForReturn()
		return false
	}
	DechuProgress = 2
	fmt.Println("Le groupe s'enfonce dans les broussailles aux abords du campement, sur les traces de la bête...")
	time.Sleep(500 * time.Millisecond)

	monstre, ok := enemies.SpawnMonster(world.Aleatoire("ville2"))
	if !ok || monstre == nil {
		fmt.Println("La chasse ne donne rien : la créature s'est déjà enfuie plus loin.")
		WaitForReturn()
		return false
	}

	fmt.Printf("Un %s surgit soudain, vous barrant la route !\n", monstre.NOM)

	if StartCombat == nil {
		fmt.Println("Le système de combat n'est pas encore connecté (la variable StartCombat n'a pas été assignée).")
		return false
	}

	if StartCombat(p, monstre) {
		fmt.Println("Les recrues vous acclament en rentrant au feu de camp : « Bienvenue chez les oubliés ! En route pour la capitale. »")
		WaitForReturn()
		return true
	}

	fmt.Println("Le groupe recule en désordre, la bête toujours en liberté. Il faudra retenter l'aventure.")
	WaitForReturn()
	return false
}