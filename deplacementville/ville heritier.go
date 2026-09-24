package ProjetRED

import (
	enemies "ProjetRED/enemies"
	personnage "ProjetRED/Personnage"
	"ProjetRED/world"
	"fmt"
	"time"
)

// HeritierProgress suit l'avancée du joueur dans la maison du prince héritier.
// 0: vient d'arriver, 1: a parlé au capitaine (peut partir en patrouille)
var HeritierProgress int

func VilleH(p *personnage.Character) {
	Faction = "Heritier"

	for {
		fmt.Println("\n==== Maison du Prince Héritier ====")
		fmt.Printf("%s se trouve dans la cour bien entretenue de la maison royale.\n", p.Nom)
		fmt.Println("1: Parler au capitaine Kaito, chef de la garde du prince")
		fmt.Println("2: Accompagner la patrouille aux abords du camp")
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
			capitaineKaito(p)
		case 2:
			if patrouille(p) {
				// La quête de la maison du prince héritier est terminée :
				// on enchaîne directement sur la capitale.
				Ville4(p, "Heritier")
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

func capitaineKaito(p *personnage.Character) {
	if HeritierProgress == 0 {
		fmt.Println("Le capitaine Kaito vous accueille d'un signe de tête respectueux.")
		fmt.Println("« Le prince héritier est un homme juste. Nous nous battons pour que l'ordre soit maintenu après le couronnement. »")
		fmt.Println("« Si tu veux vraiment nous rejoindre, prouve-le sur le terrain : accompagne la prochaine patrouille. »")
		HeritierProgress = 1
	} else {
		fmt.Println("Kaito vous salue à nouveau. « Prêt à servir le futur roi ? »")
	}
	WaitForReturn()
}

// patrouille fait affronter un monstre de ville2 au joueur.
// Elle renvoie true si la quête est validée (victoire), ce qui déclenche
// automatiquement l'enchaînement vers Ville4.
func patrouille(p *personnage.Character) bool {
	fmt.Printf("%s se présente devant les soldats en formation, prêts pour la ronde du matin.\n", p.Nom)
	time.Sleep(500 * time.Millisecond)

	if HeritierProgress < 1 {
		fmt.Println("Un garde vous arrête. « On ne prend que des recrues recommandées. Parle d'abord au capitaine. »")
		WaitForReturn()
		return false
	}
	HeritierProgress = 2
	fmt.Println("La patrouille avance en silence jusqu'à la lisière des bois qui bordent le camp...")
	time.Sleep(500 * time.Millisecond)

	monstre, ok := enemies.SpawnMonster(world.Aleatoire("ville2"))
	if !ok || monstre == nil {
		fmt.Println("La ronde se termine sans incident : les bois sont calmes ce matin-là.")
		WaitForReturn()
		return false
	}

	fmt.Printf("Un %s jaillit soudain des fourrés, crocs en avant !\n", monstre.NOM)

	if StartCombat == nil {
		fmt.Println("Le système de combat n'est pas encore connecté (la variable StartCombat n'a pas été assignée).")
		return false
	}

	if StartCombat(p, monstre) {
		fmt.Println("La patrouille rentre victorieuse. Le vétéran vous frappe l'épaule : « Tu as ta place parmi nous. En route pour la capitale. »")
		WaitForReturn()
		return true
	}

	fmt.Println("La patrouille bat en retraite dans le désordre. Il faudra retenter ta chance pour gagner leur confiance.")
	WaitForReturn()
	return false
}