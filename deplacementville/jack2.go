package ProjetRED

import (
	"fmt"
	"os"
	"time"
)

func Jack2() {
	var saisie string
	fmt.Println("bonjours apprenti voyageur!")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("je suis jack ton fidèle amis donneur de quête...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("chiante il faut se l'avouer")
	fmt.Println("mais que serai se monde sans ?")
	fmt.Println("c'est alors pour cela que j'ai besoin de ton aide!")
	fmt.Println("je voudrais que tu tue l'Onihitokuchi?")
	fmt.Println("en échange d'une récompense bien sur")
	fmt.Println("si tu n'accepte pas il y aura des concéquances désastreuse sur cette ville!")
	fmt.Println("alors si tu accepte ma quête appui sur o, sinon n, ou bien m pour le menu")
	for {
		saisie = ""
		fmt.Scanln(&saisie)

		if saisie == "m" {
			Menuvilleversville()
			return
		}
		if saisie == o {
			fmt.Println("merci tu vas sauver la ville et tout ses habitants")
			fmt.Println("je te donnerais t'as récompense une fois l'ogre  battu")
			Ville1()
			return
		}
		if saisie == n {
			fmt.Println("ho non, trop tard, l'ogre arrive, il saccage la ville!")
			fmt.Println("c'est de t'as faute !")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("GAME OVER! ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("retry?")
			time.Sleep(2 * time.Second)
			os.Exit(0)
		}
		if saisie != o && saisie != n && saisie != m {
			fmt.Println("ça n'a pas marcher il faut inscrire o, n ou m pour le menu")
			fmt.Println("veuillez saisir un choix valide")
		}
	}
}
