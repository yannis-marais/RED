package ProjetRED

import (
	"fmt"
	"strconv"
)

var vv1 = true
var vv2 = false
var vv3 = false
var vv4 = false

func Deplacement() {
	villesValides := []bool{vv1, vv2, vv3, vv4}
	villetransport := []func(){Ville1, Ville2, Ville3, Ville4}

	for {
		fmt.Println("aller vers : ")
		fmt.Println("appui sur 1 pour la ville 1")
		fmt.Println("appui sur 2 pour la ville 2")
		fmt.Println("appui sur 3 pour la ville 3")
		fmt.Println("appui sur 4 pour la ville 4")
		fmt.Println("appui sur m pour le menu")

		var saisie string
		fmt.Scanln(&saisie)

		if saisie == "m" {
			Menuvilleversville()
			return
		}

		choixville, err := strconv.Atoi(saisie)
		if err != nil || choixville < 1 || choixville > 4 {
			fmt.Println("erreur, veuillez entrer un chiffre entre 1 et 4, ou m pour le menu")
			continue
		}

		if !villesValides[choixville-1] {
			fmt.Println("vous ne pouvez pas encore aller dans cette ville")
			continue
		}

		villetransport[choixville-1]()
		return
	}
}
