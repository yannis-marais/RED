package ProjetRED

import "fmt"

func ville2() {
	for {
		var saisie string
		fmt.Println("vous êtes dans la ville 2")
		fmt.Println("vous avez débloqué l'interaction avec le marchand")
		fmt.Println("appuie sur a pour acceder au marchand")
		fmt.Println("appui sur m pour le menu")
		fmt.Println("appui sur j pour aller parler a jack")
		fmt.Println("appui sur g pour aller dans la grotte")
		fmt.Scanln(&saisie)

		switch saisie {
		case "m", "M":
			menu()
			return
		case "j", "J":
			jack2()
			return
		case "g", "G":
			grotte()
			return
		case "a", "A":
			marchand()
			return
		default:
			fmt.Println("erreur, veuillez entrer g, j, a ou m")
		}
	}
}
