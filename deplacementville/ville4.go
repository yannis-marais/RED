package ProjetRED

import (
	enemies "ProjetRED/enemies"
	personnage "ProjetRED/Personnage"
	"ProjetRED/world"
	"fmt"
)

// Faction retient le camp choisi par le joueur en ville1 ("Heritier" ou "Dechu").
var Faction string

// StartCombat est injecté par le package principal (voir menu.go, SetCurrentPlayer)
// pour éviter un import circulaire entre ce package et celui qui contient Battle.go.
var StartCombat func(p *personnage.Character, monstre *enemies.MONSTER) bool

func Ville4(p *personnage.Character, faction string) {
	if faction != "" {
		Faction = faction
	}

	for {
		fmt.Println("\n==== La Capitale ====")
		fmt.Printf("%s arrive devant les portes du palais royal. Les deux routes, celle de l'héritier et celle du déchu, se rejoignent ici.\n", p.Nom)
		fmt.Println("1: Avancer vers la salle du trône")
		fmt.Println("2: Écouter les rumeurs qui circulent dans la capitale")
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
			salleDuTrone(p)
			return
		case 2:
			rumeurs(p)
		case 0:
			return
		default:
			fmt.Println("Erreur, veuillez entrer un choix valide")
			WaitForReturn()
		}
	}
}

func rumeurs(p *personnage.Character) {
	switch Faction {
	case "Heritier":
		fmt.Println("Les habitants murmurent que le conseil attend impatiemment un roi capable de rétablir l'ordre.")
	case "Dechu":
		fmt.Println("Dans les ruelles, on chuchote qu'une partie du peuple espère enfin un souverain qui les ait vraiment écoutés.")
	default:
		fmt.Println("Les rumeurs vont bon train, mais rien de précis n'en ressort.")
	}
	WaitForReturn()
}

func salleDuTrone(p *personnage.Character) {
	switch Faction {
	case "Heritier":
		fmt.Println("La garde du trône bloque le passage : pour que le couronnement soit légitime, il faut prouver sa valeur les armes à la main.")
	case "Dechu":
		fmt.Println("Les fidèles de l'ancien ordre se dressent devant vous : le trône ne se prend pas sans combattre pour lui.")
	default:
		fmt.Println("Une garde inconnue vous barre la route vers le trône.")
	}
	WaitForReturn()

	boss, ok := enemies.SpawnMonster(world.Aleatoire("Ville 4"))
	if !ok || boss == nil {
		fmt.Println("Le champion du trône ne s'est pas présenté... la voie est étrangement libre.")
		coronation(p)
		return
	}

	if StartCombat == nil {
		fmt.Println("Le système de combat n'est pas encore connecté (la variable StartCombat n'a pas été assignée).")
		return
	}

	victoire := StartCombat(p, boss)
	if !victoire {
		fmt.Println("Le combat ne tourne pas en votre faveur : le trône devra attendre.")
		return
	}

	coronation(p)
}

func coronation(p *personnage.Character) {
	switch Faction {
	case "Heritier":
		fmt.Printf("Dans une salle du trône enfin silencieuse, %s regarde le prince héritier s'agenouiller devant le grand prêtre.\n", p.Nom)
		fmt.Println("La couronne se pose sur son front : il est désormais roi, légitime et reconnu par tout le royaume.")
	case "Dechu":
		fmt.Printf("%s regarde le prince déchu gravir lentement les marches du trône, sous les acclamations de ceux qui n'y croyaient plus.\n", p.Nom)
		fmt.Println("Il s'assoit sur le trône : par la force et la loyauté de ses partisans, il devient le nouveau souverain.")
	default:
		fmt.Println("Le trône reste vide un instant de plus, le destin du royaume encore incertain.")
	}
	WaitForReturn()
}