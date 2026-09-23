package ProjetRED

import (
	"fmt"
	"os"
	"time"
)

func Jack4() {
	var saisie string
	fmt.Println("bonjours apprenti voyageur!")
	time.Sleep(1 * time.Second)
	fmt.Print("je suis jack ton fidèle amis donneur de quête...")
	time.Sleep(1 * time.Second)
	fmt.Println("chiante il faut se l'avouer")
	time.Sleep(1 * time.Second)
	fmt.Println("mais que serai se monde sans ?")
	time.Sleep(1 * time.Second)
	fmt.Println("c'est alors pour cela que j'ai besoin de ton aide!")
	time.Sleep(1 * time.Second)
	fmt.Println("je voudrais que tu tue le boss Maxime?")
	time.Sleep(1 * time.Second)
	fmt.Println("en échange d'une récompense bien sur")
	time.Sleep(1 * time.Second)
	fmt.Println("si tu n'accepte pas il y aura des concéquances désastreuse sur cette ville!")
	time.Sleep(1 * time.Second)
	fmt.Println("alors si tu accepte ma quête appui sur o, sinon n, ou bien m pour le menu")
	for {
		saisie = ""
		fmt.Scanln(&saisie)

		if saisie == o {
			fmt.Println("merci tu vas sauver la ville et tout ses habitants")
			fmt.Println("je te donnerais t'as récompense une fois le boss battu")
			//Ville1()
			return
		}
		if saisie == n {
			fmt.Println("ho non, trop tard, le roi démon : MAXIME...")
			time.Sleep(3 * time.Second)
			fmt.Println("il...")
			time.Sleep(1 * time.Second)
			fmt.Println("il...")
			time.Sleep(1 * time.Second)
			fmt.Println("il...")
			time.Sleep(1 * time.Second)
			fmt.Println("il est là!! il saccage la ville! ")
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
