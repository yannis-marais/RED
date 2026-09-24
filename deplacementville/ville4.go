package ProjetRED

import (
	Marchand "ProjetRED/Marchand"
	personnage "ProjetRED/Personnage"

	//Menu "ProjetRED/Menu"
	"fmt"
)

func Ville4(p *personnage.Character) {
	for {
		if !QueteVille3Terminee {
			fmt.Println("Tu dois d'abord terminer la quête du désert (ville 3) avant d'accéder à la ville 4.")
			//Menu.MainMenu()
			return
		}

		const j = "j"
		const plaine = "p"
		const a = "a"

		transportdansvillequatre := map[string]func(){
			j:      Jack4,
			plaine: Plaine,
			a:      func() { Marchand.Marchand(p) },
		}
		var saisie string

		for {
			fmt.Println("vous êtes dans la ville 4")
			fmt.Println("appui sur m pour le menu")
			fmt.Println("appui sur j pour aller parler a jack (conseiller avant la plaine)")
			fmt.Println("appui sur p pour aller dans la plaine")
			fmt.Println("ou bien a pour le marchand")

			fmt.Scanln(&saisie)

			if saisie == "m" {
				//Menu.MainMenu()
			}
			if saisie == j || saisie == plaine {
				break
			}
			if saisie == a {
				break
			}

			fmt.Println("erreur, veuillez entrer une lettre p, j ou m pour le menu")
		}

		transportdansvillequatre[saisie]()
	}
}
