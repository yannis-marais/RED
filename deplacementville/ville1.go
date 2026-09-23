package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"fmt"
	Menu "ProjetRED/Menu"
)

func Ville1(p *personnage.Character) {
	for {
		const j = "j"
		const f = "f"

		transportdansvilleun := map[string]func(){
			j: Jack1,
			f: Foret,
		}

		var saisie string

		for {
			fmt.Println("==== Forteresse du royaume ====")
			fmt.Println("j: pour aller parler a jack (conseiller avant la forêt)")
			fmt.Println("appui sur f pour aller dans la forêt")
			fmt.Println("m: pour ouvrire le menu")

			fmt.Scanln(&saisie)

			if saisie == "m" {
				Menu.MainMenu(p)
			}
			if saisie == j || saisie == f {
				break
			}

			fmt.Println("erreur, veuillez entrer une lettre f, j ou m pour le menu")
		}

		transportdansvilleun[saisie]()
	}
}
