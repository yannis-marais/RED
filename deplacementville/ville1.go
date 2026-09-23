package ProjetRED

import (
	"fmt"
)

func Ville1() {
	const j = "j"
	const f = "f"

	transportdansvilleun := map[string]func(){
		j: Jack1,
		f: Foret,
	}

	var saisie string

	for {
		fmt.Println("vous êtes dans la ville 1")
		fmt.Println("appui sur m pour le menu")
		fmt.Println("appui sur j pour aller parler a jack (conseiller avant la forêt)")
		fmt.Println("appui sur f pour aller dans la forêt")

		fmt.Scanln(&saisie)

		if saisie == "m" {
			Menuvilleversville()
			return
		}
		if saisie == j || saisie == f {
			break
		}

		fmt.Println("erreur, veuillez entrer une lettre f, j ou m pour le menu")
	}

	transportdansvilleun[saisie]()
}
