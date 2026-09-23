package ProjetRED

import (
	"fmt"
)

func Menuvilleversville() {
	const o = "o"
	const n = "n"
	var saisie string
	for {
		fmt.Println("êtes vous sûr de vouloir aller vers une ville?")
		fmt.Println("appui sur m pour menu principal")
		fmt.Println("appui sur o pour oui sinon n")

		fmt.Scanln(&saisie)

		if saisie == "m" {
			Menuvilleversville()
			return
		}
		if saisie == o {
			Deplacement()
			return
		}
		if saisie == n {
			break
		}
		fmt.Println("erreur, veuillez entrer une lettre o, n ou m pour le menu principal")
	}
}
