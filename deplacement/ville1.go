package ProjetRED

import "fmt"

func ville1() {
	for {
		var saisie string
		fmt.Println("vous êtes dans la ville 1")
		fmt.Println("appui sur m pour le menu")
		fmt.Println("appui sur j pour aller parler a jack")
		fmt.Println("appui sur f pour aller dans la forêt")
		fmt.Scanln(&saisie)

		switch saisie {
		case "m", "M":
			menu()
			return
		case "j", "J":
			jack1()
			return
		case "f", "F":
			foret()
			return
		default:
			fmt.Println("erreur, veuillez entrer f, j ou m pour le menu")
		}
	}
}
