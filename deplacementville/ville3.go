package ProjetRED

import (
	Menu "PrjoetRED/Menu"
	Marchand "ProjetRED/Marchand"
	"fmt"
)

func Ville3() {
	for {
		if !QueteVille2Terminee {
			fmt.Println("Tu dois d'abord terminer la quête de la grotte (ville 2) avant d'accéder à la ville 3.")
			Menu.MainMenu()
		}

		const d = "d"
		const j = "j"
		const a = "a"

		transportdansvilletrois := map[string]func(){
			j: Jack3,
			d: Desert,
			a: func() { Marchand.Marchand(Ville3) },
		}

		var saisie string

		for {
			fmt.Println("vous êtes dans la ville 3")
			fmt.Println("appui sur m pour le menu")
			fmt.Println("appui sur j pour aller parler a jack (conseiller avant le desert)")
			fmt.Println("appui sur d pour aller dans le desert")
			fmt.Println("ou bien a pour le marchand")

			fmt.Scanln(&saisie)

			if saisie == "m" {
				Menu.MainMenu()
			}
			if saisie == j || saisie == d || saisie == a {
				break
			}

			fmt.Println("erreur, veuillez entrer une lettre d, j ou m pour le menu")
		}

		transportdansvilletrois[saisie]()
	}
}
