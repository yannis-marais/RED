package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"fmt"
)

func Forgeron(p *personnage.Character) {
	if p == nil {
		fmt.Println("Le forgeron ne peut pas traiter un personnage vide.")
		return
	}
	fmt.Println("Le forgeron est temporairement indisponible.")
}
