package ProjetRED

import (
    Marchand "ProjetRED/Marchand"
    "fmt"
)

func Ville4() {
    const j = "j"
    const p = "p"
    const a = "a"

    transportdansvillequatre := map[string]func(){
        j: Jack4,
        p: Plaine,
        a: func() { Marchand.Marchand(Ville4) },
    }
    var saisie string

    for {
        fmt.Println("vous êtes dans la ville 4")
        fmt.Println("appui sur m pour le menu")
        fmt.Println("appui sur j pour aller parler a jack (conseiller avant la plaine)")
        fmt.Println("appui sur p pour aller dans la plaine")
        fmt.Println("ou bien a pour le marchand")

        fmt.Scanln(&saisie)

        if saisie == "m" {
            Menuvilleversville()
            return
        }
        if saisie == j || saisie == p {
            break
        }
        if saisie == a {
            break
        }

        fmt.Println("erreur, veuillez entrer une lettre p, j ou m pour le menu")
    }

    transportdansvillequatre[saisie]()
}