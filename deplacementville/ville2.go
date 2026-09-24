package ProjetRED

import (
	Menu "PrjoetRED/Menu"
	Marchand "ProjetRED/Marchand"
	"fmt"
)

func Ville2() {
	for {
		if !QueteVille1Terminee {
			fmt.Println("Tu dois d'abord terminer la quête de la forêt (ville 1) avant d'accéder à la ville 2.")
			Menu.MainMenu()
		}

		const j = "j"
		const g = "g"
		const a = "a"

		transportdansvilledeux := map[string]func(){
			j: Jack2,
			g: Grotte,
			a: func() { Marchand.Marchand(Ville2) },
		}

		var saisie string

		for {
			fmt.Println("vous êtes dans la ville 2")
			fmt.Println("vous avez débloquer l'interaction avec le marchand")
			fmt.Println("appuie sur a pour acceder au marchand")
			fmt.Println("appui sur m pour le menu")
			fmt.Println("appui sur j pour aller parler a jack (conseiller avant la grotte)")
			fmt.Println("appui sur g pour aller dans la grotte")
			fmt.Println("appui sur a pour le marchand")
			fmt.Scanln(&saisie)

			if saisie == "m" {
				Menu.MainMenu()
			}
			if saisie == j || saisie == g || saisie == a {
				break
			}

			fmt.Println("erreur, veuillez entrer une lettre g, j, a ou m")
		}

		transportdansvilledeux[saisie]()
	}
}
