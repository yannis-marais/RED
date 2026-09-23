package ProjetRED

import (
	"fmt"
	"os"
	"time"
)

func wait(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func jack1() {
	for {
		var saisie string
		fmt.Println("bonjours apprenti voyageur!")
		wait(500)
		fmt.Print("je suis jack ton fidèle amis donneur de quête...")
		wait(500)
		fmt.Println("\nchiante il faut se l'avouer")
		fmt.Println("mais que serai se monde sans ?")
		fmt.Println("c'est alors pour cela que j'ai besoin de ton aide!")
		fmt.Println("je voudrais que tu tue le troll des forêts?")
		fmt.Println("en échange d'une récompense bien sur")
		fmt.Println("si tu n'accepte pas il y aura des consequences desastreuses sur cette ville!")
		fmt.Println("alors si tu accepte ma quête appui sur o, sinon n")
		fmt.Scanln(&saisie)

		switch saisie {
		case "m", "M":
			menu()
			return
		case "o", "O":
			fmt.Println("merci tu vas sauver la ville et tout ses habitants")
			fmt.Println("je te donnerais t'as récompense une fois le troll battu")
			ville1()
			return
		case "n", "N":
			fmt.Println("ho non, trop tard, le troll arrive, il saccage la ville!")
			fmt.Println("c'est de t'as faute !")
			wait(500)
			fmt.Println("GAME OVER!")
			wait(500)
			fmt.Println("retry?")
			wait(2000)
			os.Exit(0)
		default:
			fmt.Println("erreur, veuillez entrer une lettre o, n ou m pour le menu")
		}
	}
}
