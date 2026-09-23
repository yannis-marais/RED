package ProjetRED

import (
	"fmt"
)

func Ville3() {
	const d = "d"
	const j = "j"

	transportdansvilletrois := map[string]func(){
		j: Jack3,
		d: Desert,
	}

	var saisie string

	for {
		fmt.Println("vous êtes dans la ville 3")
		fmt.Println("appui sur m pour le menu")
		fmt.Println("appui sur j pour aller parler a jack (conseiller avant le desert)")
		fmt.Println("appui sur d pour aller dans le desert")

		fmt.Scanln(&saisie)

		if saisie == "m" {
			Menuvilleversville()
			return
		}
		if saisie == j || saisie == d {
			break
		}

		fmt.Println("erreur, veuillez entrer une lettre d, j ou m pour le menu")
	}

	transportdansvilletrois[saisie]()
}
